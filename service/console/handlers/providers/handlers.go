package providers

import (
	"auth/service/database"
	"auth/service/database/types"
	"errors"
	"fmt"
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

	providerName := ctx.Param("provider")

	var provider types.OAuthProvider
	err := db.Unscoped().Where("name = ?", providerName).First(&provider).Error
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
	providerName := ctx.Param("provider")
	var provider types.OAuthProvider
	// ignore DELETED_AT.
	err := db.Unscoped().Where("name = ?", providerName).First(&provider).Error

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
			Config: types.OAuthConfig{Config: oauth2Config},
			Name:   providerName,
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
	var provider types.OAuthProvider
	// ignore DELETED_AT.
	err := db.Where("provider = ?", providerName).First(&provider).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	if err := db.Delete(&provider).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.NoContent(http.StatusNoContent)
}

func GetProvidersStateHandler(ctx echo.Context) error {
	var providers []types.OAuthProvider
	if err := db.Unscoped().Find(&providers).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	type ProviderState struct {
		Enabled      bool   `json:"enabled"`
		ClientID     string `json:"clientID"`
		ClientSecret string `json:"clientSecret"`
	}

	states := make(map[string]ProviderState, 1)

	for _, provider := range providers {
		states[provider.Name] = ProviderState{
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
