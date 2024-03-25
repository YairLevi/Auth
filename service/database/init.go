package database

import (
	"auth/service/database/types"
	"gorm.io/gorm"
)

var DB *gorm.DB

func PrepareDB(db *gorm.DB) {
	d, _ := db.DB()

	// for sqlite, enable foreign_keys support.
	var enabled bool
	d.QueryRow("PRAGMA foreign_keys = ON;").Scan(&enabled)
	err := db.AutoMigrate(
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
		panic(err)
	}
}
