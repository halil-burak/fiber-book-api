package handler

import (
	"github.com/gofiber/fiber"
	"github.com/halil-burak/fiber-rest-api/database"
	"github.com/halil-burak/fiber-rest-api/model"
)

func AddUser(c *fiber.Ctx) {

	db := database.DBConn
	userC := new(model.UserCreate)

	if err := c.BodyParser(userC); err != nil {
		c.Status(503).Send(err)
		return
	}

	var newUser model.User
	newUser.Name = userC.Name

	db.Create(&newUser)

	for _, lang := range userC.Languages {
		language, err := CreateLanguageIfNotExists(lang)
		if err != nil {
			c.Status(501).Send(err)
			db.Rollback()
			return
		}
		err = AddLanguageUser(&newUser, language)
		if err != nil {
			c.Status(501).Send(err)
			db.Rollback()
			return
		}
	}
	c.JSON(newUser)
}

func AddLanguageUser(user *model.User, lang *model.Language) error {
	db := database.DBConn
	response := db.Model(user).Association("Languages").Append(lang)
	db.Save(user)
	return response.Error
}

func CreateLanguageIfNotExists(langID string) (*model.Language, error) {
	db := database.DBConn
	var language model.Language
	response := db.FirstOrCreate(&language, model.Language{Name: langID})
	if response.Error != nil {
		return nil, response.Error
	}
	return &language, nil
}

func AddLanguage(c *fiber.Ctx) {
	db := database.DBConn
	language := new(model.LanguageCreate)

	if err := c.BodyParser(language); err != nil {
		c.Status(503).Send(err)
		return
	}

	var newLang model.Language
	newLang.Name = language.Name
	db.Create(&newLang)
	c.JSON(newLang)
}

func GetAllUsers(c *fiber.Ctx) {
	db := database.DBConn
	var users []model.User
	db.Preload("Languages").Find(&users)
	userlist := make([]model.UserGet, len(users))

	for i, u := range users {
		var userg model.UserGet
		userg.ID = u.ID
		userg.Name = u.Name
		ls := make([]string, len(u.Languages))
		for i, l := range u.Languages {
			ls[i] = l.Name
		}
		userg.Languages = ls
		userlist[i] = userg
	}
	c.JSON(userlist)
}

func GetOneUser(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var user model.User
	var userg model.UserGet
	db.Preload("Languages").First(&user, id)
	switch {
	case db.Error != nil:
		c.Send(503)
		return
	case user.ID == 0:
		c.Status(404).Send("Not found")
		return
	}

	langs := make([]string, len(user.Languages))
	userg.ID = user.ID
	userg.Name = user.Name
	for i, l := range user.Languages {
		langs[i] = l.Name
	}
	userg.Languages = langs
	c.JSON(userg)
}
