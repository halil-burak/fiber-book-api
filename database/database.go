package database

import (
	"fmt"

	"github.com/halil-burak/fiber-rest-api/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	// DBConn is the global database connection instance
	DBConn *gorm.DB
)

//ConnectDB connects to the db
func ConnectDB() {
	var err error
	DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Could not open connection to the db")
	}

	fmt.Println("Connected to the db!")
	fmt.Println("Connection Opened to Database")
	DBConn.AutoMigrate(&model.Book{})
	DBConn.AutoMigrate(&model.Author{})
	DBConn.AutoMigrate(&model.Category{})
	fmt.Println("Database Migrated")
}
