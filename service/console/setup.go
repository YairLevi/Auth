package console

import (
	handlers2 "auth/service/console/handlers"
	"auth/service/console/handlers/providers"
	"auth/service/database"
	"github.com/labstack/echo/v4"
)

var db = database.DB

func setupAppsEndpoints(router *echo.Group) {
	router.GET("/", handlers2.ListAppsHandler)
	router.POST("/", handlers2.CreateAppHandler)
	router.GET("/:appId", handlers2.GetAppHandler)
	router.DELETE("/:appId", handlers2.DeleteAppHandler)
}

func setupSecurityEndpoints(router *echo.Group) {
	router.GET("/", handlers2.GetSecuritySettingsHandler)
	router.POST("/lockout/threshold", handlers2.SetLockoutThresholdHandler)
	router.POST("/lockout/duration", handlers2.SetLockoutDurationHandler)
	router.POST("/origins", handlers2.AddOriginHandler)
	router.POST("/session", handlers2.SetSessionKeyHandler)
}

func setupUsersEndpoints(router *echo.Group) {
	router.GET("/", handlers2.ListUsersHandler)
	router.POST("/", handlers2.CreateUserHandler)
	router.DELETE("/:userId", handlers2.DeleteUserHandler)
}

func setupOAuthEndpoints(router *echo.Group) {
	router.GET("/providers", providers.GetProvidersStateHandler)
	router.PUT("/:provider/enable", providers.EnableProviderHandler)
	router.PUT("/:provider/update", providers.UpdateProviderCredentialsHandler)
	router.DELETE("/:provider/disable", providers.DisableProviderHandler)
}

func SetupSingleAppEndpoints(router *echo.Group) {
	setupUsersEndpoints(router.Group("/users"))
	setupOAuthEndpoints(router.Group("/oauth"))
	setupSecurityEndpoints(router.Group("/security"))
}

func SetupEndpoints(server *echo.Echo) {
	setupAppsEndpoints(server.Group("/apps"))
	SetupSingleAppEndpoints(server.Group("/apps/:appId"))
}
