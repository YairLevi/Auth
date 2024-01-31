package service

import (
	"auth-service/database"
	"auth-service/database/types"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

const ApplicationHeader = "X-App-ID"
const ConsoleOrigin = "http://wails.localhost"
const ConsoleDevOrigin = "http://localhost:5173"

func DynamicCORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		origin := c.Request().Header.Get("Origin")
		appID := c.Request().Header.Get(ApplicationHeader)
		var allowed types.Origin
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
