package database

import (
	types2 "auth/service/database/types"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open(sqlite.Open("database.database"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(fmt.Sprint("Failed to open database connection:", err))
	}

	err = db.AutoMigrate(
		&types2.User{},
		&types2.Origin{},
		&types2.App{},
		&types2.OAuthProvider{},
	)
	if err != nil {
		panic(fmt.Sprint("GORM failed to migrate types to proper SQL tables."))
	}

	DB = db
}
