package main

import (
	"auth-service/console"
	"auth-service/service"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func main() {
	server := echo.New()
	server.Use(service.DynamicCORS)

	console.SetupEndpoints(server)
	service.SetupEndpoints(server)

	// HealthChecker Endpoint
	server.GET("/test", func(ctx echo.Context) error {
		log.Println("Someone entered the service.")
		return ctx.JSON(http.StatusOK, "")
	})

	for _, route := range server.Routes() {
		fmt.Println("base:\t", route.Path)
	}

	log.Fatal(server.Start(":9999"))
}
