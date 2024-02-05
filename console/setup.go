package console

import (
	"auth/console/handlers"
	"auth/console/handlers/providers"
	"auth/database"
	"auth/database/types"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

var db = database.DB

func setupAppsEndpoints(router *echo.Group) {
	router.GET("/", handlers.ListAppsHandler)
	router.POST("/", handlers.CreateAppHandler)
	router.GET("/:appId", handlers.GetAppHandler)
	router.DELETE("/:appId", handlers.DeleteAppHandler)
}

func setupSecurityEndpoints(router *echo.Group) {
	router.GET("/", handlers.GetSecuritySettingsHandler)
	router.POST("/lockout/threshold", handlers.SetLockoutThresholdHandler)
	router.POST("/lockout/duration", handlers.SetLockoutDurationHandler)
	router.POST("/origins", handlers.AddOriginHandler)
	router.POST("/session", handlers.SetSessionKeyHandler)
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

func SetupSingleAppEndpoints(router *echo.Group) {
	router.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			appID := ctx.Param("appId")
			if err := db.Where("id = ?", appID).First(&types.App{}).Error; err != nil {
				fmt.Println(err)
				return ctx.JSON(http.StatusBadRequest, "invalid app ID")
			}
			return next(ctx)
		}
	})

	setupUsersEndpoints(router.Group("/users"))
	setupOAuthEndpoints(router.Group("/oauth"))
	setupSecurityEndpoints(router.Group("/security"))
}

func SetupEndpoints(server *echo.Echo) {
	setupAppsEndpoints(server.Group("/apps"))
	SetupSingleAppEndpoints(server.Group("/apps/:appId"))
}
