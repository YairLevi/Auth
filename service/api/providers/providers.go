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

func GetConfigByAppID(appID string, provider string) (*types.OAuthConfig, error) {
	var oauthConfig types.OAuthProvider
	if err := db.Where("app_id = ? AND provider = ?", appID, provider).First(&oauthConfig).Error; err != nil {
		return nil, err
	}
	return &oauthConfig.Config, nil
}
