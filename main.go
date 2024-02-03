package main

import (
	"auth-service/console"
	"auth-service/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
)

func main() {
	server := echo.New()
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials:                         true,
		UnsafeWildcardOriginWithAllowCredentials: true,
	}))
	server.Use(service.DynamicCORS)

	console.SetupEndpoints(server)
	service.SetupEndpoints(server)

	// HealthChecker Endpoint
	server.GET("/test", func(ctx echo.Context) error {
		log.Println("Someone entered the service.")
		return ctx.JSON(http.StatusOK, "")
	})

	log.Fatal(server.Start(":9999"))
}
