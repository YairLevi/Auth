package main

import (
	"auth-service/console"
	"auth-service/database"
	"auth-service/database/types"
	"auth-service/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
)

const ApplicationHeader = "X-App-ID"
const ConsoleOrigin = "http://wails.localhost"
const ConsoleDevOrigin = "http://localhost:5173"

func dynamicCORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		origin := c.Request().Header.Get("Origin")
		appID := c.Request().Header.Get(ApplicationHeader)
		var allowed types.AllowedOrigins
		err := database.DB.Where("origin = ? AND app_id = ?", origin, appID).First(&allowed).Error
		isOriginAllowed := err == nil

		// If the origin is allowed, set the CORS headers
		if isOriginAllowed || origin == ConsoleOrigin || origin == ConsoleDevOrigin {
			return middleware.CORSWithConfig(middleware.CORSConfig{
				AllowOrigins:     []string{origin},
				AllowCredentials: true,
			})(next)(c)
		}

		// If the origin is not allowed, handle it as needed
		return c.JSON(http.StatusForbidden, map[string]string{"error": "Origin not allowed"})
	}
}

func main() {
	server := echo.New()

	//server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//	AllowCredentials: true,
	//	AllowOrigins: []string{
	//		"http://localhost:5173",
	//		"http://localhost:5174",
	//		"http://localhost:3000",
	//		"http://wails.localhost",
	//	},
	//}))

	server.Use(dynamicCORS)

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
