package providers

import (
	"auth/service/database"
	"auth/service/database/types"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
	"gorm.io/gorm"
	"net/http"
)

var db = database.DB

type ProviderInfo struct {
	RedirectURL string
	Scopes      []string
	Endpoint    oauth2.Endpoint
}

func getProviderInfo(providerName string) ProviderInfo {
	switch providerName {
	case Google:
		return ProviderInfo{
			RedirectURL: fmt.Sprint("http://localhost:9999/api/", providerName, "/auth/callback"),
			Scopes:      []string{"profile", "email"},
			Endpoint:    google.Endpoint,
		}
	case Github:
		return ProviderInfo{
			RedirectURL: fmt.Sprint("http://localhost:9999/api/", providerName, "/auth/callback"),
			Scopes:      []string{"user:email"},
			Endpoint:    github.Endpoint,
		}
	default:
		return ProviderInfo{}
	}
}

func UpdateProviderCredentialsHandler(ctx echo.Context) error {
	dto := struct {
		ClientID     string `json:"clientId"`
		ClientSecret string `json:"clientSecret"`
	}{}
	if err := ctx.Bind(&dto); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	appID := ctx.Param("appId")
	providerName := ctx.Param("provider")
	if appID == "" || providerName == "" {
		return ctx.JSON(http.StatusBadRequest, "empty app ID or provider name")
	}

	var provider types.OAuthProvider
	err := db.Unscoped().Where("app_id = ? AND provider = ?", appID, providerName).First(&provider).Error
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	provider.Config.ClientID = dto.ClientID
	provider.Config.ClientSecret = dto.ClientSecret
	if err := db.Save(&provider).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.NoContent(http.StatusOK)
}

func EnableProviderHandler(ctx echo.Context) error {
	appID := ctx.Param("appId")
	providerName := ctx.Param("provider")
	if appID == "" || providerName == "" {
		return ctx.JSON(http.StatusBadRequest, "empty app ID or provider name")
	}

	var provider types.OAuthProvider
	// ignore DELETED_AT.
	err := db.Unscoped().Where("app_id = ? AND provider = ?", appID, providerName).First(&provider).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		info := getProviderInfo(providerName)
		oauth2Config := &oauth2.Config{
			ClientID:     "",
			ClientSecret: "",
			RedirectURL:  info.RedirectURL,
			Scopes:       info.Scopes,
			Endpoint:     info.Endpoint,
		}

		config := types.OAuthProvider{
			AppID:    appID,
			Provider: providerName,
			Config:   types.OAuthConfig{Config: oauth2Config},
		}

		if err := db.Create(&config).Error; err != nil {
			return ctx.JSON(http.StatusInternalServerError, err)
		}

	} else if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)

	} else /* There is a provider */ {
		if err := db.Unscoped().Model(&provider).Update("deleted_at", nil).Error; err != nil {
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}

	return ctx.NoContent(http.StatusOK)
}

func DisableProviderHandler(ctx echo.Context) error {
	providerName := ctx.Param("provider")
	appID := ctx.Param("appId")
	if appID == "" || providerName == "" {
		return ctx.JSON(http.StatusBadRequest, "empty app ID or provider name")
	}

	var provider types.OAuthProvider
	// ignore DELETED_AT.
	err := db.Where("app_id = ? AND provider = ?", appID, providerName).First(&provider).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	if err := db.Delete(&provider).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.NoContent(http.StatusNoContent)
}

func GetProvidersStateHandler(ctx echo.Context) error {
	appID := ctx.Param("appId")
	if appID == "" {
		return ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid app ID"})
	}

	var providers []types.OAuthProvider
	if err := db.Unscoped().Where("app_id = ?", appID).Find(&providers).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	type ProviderState struct {
		Enabled      bool   `json:"enabled"`
		ClientID     string `json:"clientID"`
		ClientSecret string `json:"clientSecret"`
	}

	states := make(map[string]ProviderState, 1)

	for _, provider := range providers {
		states[provider.Provider] = ProviderState{
			Enabled:      !provider.DeletedAt.Valid,
			ClientID:     provider.Config.ClientID,
			ClientSecret: provider.Config.ClientSecret,
		}
	}

	for _, providerName := range ProviderList {
		if _, ok := states[providerName]; !ok {
			states[providerName] = ProviderState{
				Enabled:      false,
				ClientID:     "",
				ClientSecret: "",
			}
		}
	}

	return ctx.JSON(http.StatusOK, &states)
}
