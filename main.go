package main

import (
	"auth-service/console"
	"auth-service/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
)

func main() {
	server := echo.New()

	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins: []string{
			"http://localhost:5173",
			"http://localhost:5174",
			"http://localhost:3000",
			"http://wails.localhost",
		},
	}))

	console.SetupEndpoints(server)
	service.SetupEndpoints(server)

	// HealthChecker Endpoint
	server.GET("/test", func(ctx echo.Context) error {
		log.Println("Someone entered the service.")
		return ctx.JSON(http.StatusOK, gin.H{"YES": "WORKED"})
	})

	for _, route := range server.Routes() {
		fmt.Println("base:\t", route.Path)
	}

	log.Fatal(server.Start(":9999"))
}
