package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/oauth2"
)

type OAuthConfig struct {
	*oauth2.Config
}

type OAuthProvider struct {
	Model
	Config OAuthConfig
	Name   string
}

func (config *OAuthConfig) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal OAuthConfig value:", value))
	}

	return json.Unmarshal(bytes, config)
}

func (config OAuthConfig) Value() (driver.Value, error) {
	return json.Marshal(config)
}
