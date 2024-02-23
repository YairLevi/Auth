package api

import (
	"auth/service/api/providers/github"
	"auth/service/api/providers/google"
	"auth/service/api/providers/standard"
	"auth/service/api/roles"
	"github.com/labstack/echo/v4"
)

func setupBasicAuthEndpoints(router *echo.Group) {
	router.POST("/register", standard.RegisterHandler)
	router.POST("/login", standard.EmailPasswordLoginHandler)
	router.GET("/login", standard.CookieLoginHandler)
	router.POST("/logout", standard.LogoutHandler)
}

func setupGoogleAuthRoutes(router *echo.Group) {
	router.GET("/login", google.LoginHandler)
	router.GET("/auth/callback", google.CallbackHandler)
}

func setupGithubAuthRoutes(router *echo.Group) {
	router.GET("/login", github.LoginHandler)
	router.GET("/auth/callback", github.CallbackHandler)
}

func setupRolesRoutes(router *echo.Group) {
	router.GET("/", roles.GetRoles)
	router.POST("/", roles.AddRole)
	router.DELETE("/:role", roles.DeleteRole)
}

func setupUsersRoutes(router *echo.Group) {
	router.GET("/:userId/roles", roles.GetUserRoles)
	router.POST("/:userId/roles", roles.AssignRoleToUser)
	router.DELETE("/:userId/roles/:role", roles.RevokeRoleFromUser)
}

func SetupEndpoints(server *echo.Echo) {
	apiV1 := server.Group("/api")
	setupBasicAuthEndpoints(apiV1.Group("/"))
	setupGoogleAuthRoutes(apiV1.Group("/google"))
	setupGithubAuthRoutes(apiV1.Group("/github"))
	setupRolesRoutes(apiV1.Group("/roles"))
	setupUsersRoutes(apiV1.Group("/users"))
}
