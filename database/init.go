package database

import (
	"auth-service/database/types"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open(sqlite.Open("database.database"), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprint("Failed to open database connection:", err))
	}

	err = db.AutoMigrate(
		&types.User{},
		&types.App{},
		&types.OAuthProvider{},
		&types.AllowedOrigins{},
	)
	if err != nil {
		panic(fmt.Sprint("GORM failed to migrate types to proper SQL tables."))
	}

	DB = db
}
