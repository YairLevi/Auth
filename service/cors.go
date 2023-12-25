package service

import (
	"auth-service/database"
	"auth-service/database/types"
)

func AllowOriginsFunc(origin string) bool {
	var allowedOrigin types.AllowedOrigins
	return database.DB.Where("origin = ?", origin).First(&allowedOrigin).Error == nil
}
