package database

import (
	"auth/service/database/types"
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
	d, _ := db.DB()

	// for sqlite, enable foreign_keys support.
	var enabled bool
	d.QueryRow("PRAGMA foreign_keys = ON;").Scan(&enabled)
	err = db.AutoMigrate(
		&types.User{},
		&types.Origin{},
		&types.SecurityConfig{},
		&types.OAuthProvider{},
		&types.Role{},
		&types.UserRole{},
		&types.EmailFilter{},
		&types.Lockout{},
	)
	if err != nil {
		panic(fmt.Sprint("GORM failed to migrate types to proper SQL tables."))
	}

	DB = db
}
