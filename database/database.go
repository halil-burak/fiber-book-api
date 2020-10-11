package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	// DBConn is the global database connection instance
	DBConn *gorm.DB
)
