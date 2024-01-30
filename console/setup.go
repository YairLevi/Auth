package console

import (
	"auth-service/console/handlers"
	"auth-service/console/handlers/providers"
	"github.com/labstack/echo/v4"
)

func setupAppsEndpoints(router *echo.Group) {
	router.GET("/", handlers.ListAppsHandler)
	router.POST("/", handlers.CreateAppHandler)
	router.DELETE("/:appId", handlers.DeleteAppHandler)
}

func setupOriginsEndpoints(router *echo.Group) {
	router.GET("/", handlers.GetOriginsHandler)
	router.POST("/", handlers.AddOriginHandler)
}

func setupUsersEndpoints(router *echo.Group) {
	router.GET("/", handlers.ListUsersHandler)
	router.POST("/", handlers.CreateUserHandler)
	router.DELETE("/:userId", handlers.DeleteUserHandler)
}

func setupOAuthEndpoints(router *echo.Group) {
	router.GET("/providers", providers.GetProvidersStateHandler)
	router.PUT("/:provider/enable", providers.EnableProviderHandler)
	router.PUT("/:provider/update", providers.UpdateProviderCredentialsHandler)
	router.DELETE("/:provider/disable", providers.DisableProviderHandler)
}

func SetupEndpoints(server *echo.Echo) {
	setupAppsEndpoints(server.Group("/apps"))
	setupUsersEndpoints(server.Group("/apps/:appId/users"))
	setupOAuthEndpoints(server.Group("/apps/:appId/oauth"))
	setupOriginsEndpoints(server.Group("/apps/:appId/origins"))
}
