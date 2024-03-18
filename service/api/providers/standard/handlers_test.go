package standard

import (
	"auth/service/database/types"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterHandler(t *testing.T) {
	// Create a request body
	requestBody := map[string]string{
		"email":    "test@example.com",
		"username": "username123",
		"password": "password123",
	}
	jsonRequestBody, _ := json.Marshal(requestBody)

	// Create a request
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonRequestBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := testServer.NewContext(req, rec)

	// Call the handler
	err := RegisterHandler(c)

	// Check if there was an error
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	// Check if the user was created in the database
	var user types.User
	db.Where("email = ?", "test@example.com").First(&user)
	assert.Equal(t, "username123", user.Username)
}
