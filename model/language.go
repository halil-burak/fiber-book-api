package model

import "github.com/jinzhu/gorm"

type Language struct {
	gorm.Model
	Name string
}

type LanguageCreate struct {
	Name string `json:"name"`
}
