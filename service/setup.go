package service

import (
	"auth/service/providers/github"
	"auth/service/providers/google"
	"auth/service/providers/standard"
	"github.com/labstack/echo/v4"
)

func setupServiceApiEndpoints(router *echo.Group) {
	router.POST("/register", standard.RegisterHandler)
	router.POST("/login", standard.EmailPasswordLoginHandler)
	router.GET("/login", standard.CookieLoginHandler)
	router.POST("/logout", standard.LogoutHandler)
}

func setupGoogleAuthRoutes(router *echo.Group) {
	router.GET("/:appId/login", google.LoginHandler)
	router.GET("/auth/callback", google.CallbackHandler)
}

func setupGithubAuthRoutes(router *echo.Group) {
	router.GET("/:appId/login", github.LoginHandler)
	router.GET("/auth/callback", github.CallbackHandler)
}

func SetupEndpoints(server *echo.Echo) {
	setupServiceApiEndpoints(server.Group("/api"))
	setupGoogleAuthRoutes(server.Group("/api/google"))
	setupGithubAuthRoutes(server.Group("/api/github"))
}
