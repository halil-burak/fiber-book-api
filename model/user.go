package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name      string
	Languages []Language `gorm:"many2many:user_languages;"`
}

type UserCreate struct {
	Name      string   `json:"name"`
	Languages []string `json:"languages"`
}

type UserGet struct {
	ID        uint     `json:"id"`
	Name      string   `json:"name"`
	Languages []string `json:"languages"`
}
