package providers

import (
	"auth/service/database"
	"auth/service/database/types"
)

const (
	Google = "google"
	Github = "github"
)

var db = database.DB

func GetOAuthConfig(providerName string) (*types.OAuthConfig, error) {
	var oauthConfig types.OAuthProvider
	if err := db.Unscoped().Where("name = ?", providerName).First(&oauthConfig).Error; err != nil {
		return nil, err
	}
	return &oauthConfig.Config, nil
}
