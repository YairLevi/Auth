package console

import (
	"auth/service/console/handlers"
	"auth/service/console/handlers/providers"
	"auth/service/database"
	"github.com/labstack/echo/v4"
)

var db = database.DB

func setupSecurityEndpoints(router *echo.Group) {
	router.GET("/", handlers.GetSecuritySettingsHandler)
	router.PUT("/lockout/threshold", handlers.SetLockoutThresholdHandler)
	router.PUT("/lockout/duration", handlers.SetLockoutDurationHandler)
	router.PUT("/origins", handlers.AddOriginHandler)
	router.PUT("/session", handlers.SetSessionKeyHandler)
	router.DELETE("/origins/:originId", handlers.RemoveOriginHandler)
	router.POST("/emails", handlers.AddEmailFilterHandler)
	router.DELETE("/emails/:emailId", handlers.RemoveEmailFilterHandler)
}

func setupUsersEndpoints(router *echo.Group) {
	router.GET("/", handlers.ListUsersHandler)
	router.POST("/", handlers.CreateUserHandler)
	router.DELETE("/:userId", handlers.DeleteUserHandler)
	router.PATCH("/:userId", handlers.UpdateUserHandler)
}

func setupOAuthEndpoints(router *echo.Group) {
	router.GET("/providers", providers.GetProvidersStateHandler)
	router.PUT("/:provider", providers.UpdateProviderCredentialsHandler)
	router.POST("/:provider", providers.EnableProviderHandler)
	router.DELETE("/:provider", providers.DisableProviderHandler)
}

func SetupEndpoints(server *echo.Echo) {
	router := server.Group("/console")
	setupUsersEndpoints(router.Group("/users"))
	setupOAuthEndpoints(router.Group("/oauth"))
	setupSecurityEndpoints(router.Group("/security"))
}
