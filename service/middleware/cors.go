package middleware

import (
	"auth/service/database"
	"auth/service/database/types"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

const (
	ConsoleOrigin    = "http://wails.localhost"
	ConsoleDevOrigin = "http://localhost:5173"
)

func DynamicCORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		origin := c.Request().Header.Get("Origin")
		var allowed types.Origin
		err := database.DB.Where(&types.Origin{URL: origin}).First(&allowed).Error
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
