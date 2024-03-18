package main

import (
	"auth/service/api"
	"auth/service/console"
	"auth/service/database"
	"auth/service/database/types"
	authMiddleware "auth/service/middleware"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
)

func Init() {
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

	database.DB = db
}

func main() {
	Init()
	server := echo.New()
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials:                         true,
		UnsafeWildcardOriginWithAllowCredentials: true,
	}))
	server.Use(authMiddleware.DynamicCORS)

	console.SetupEndpoints(server)
	api.SetupEndpoints(server)

	// HealthChecker Endpoint
	server.GET("/test", func(ctx echo.Context) error {
		log.Println("Someone entered the service.")
		return ctx.JSON(http.StatusOK, "")
	})

	log.Fatal(server.Start(":9999"))
}
