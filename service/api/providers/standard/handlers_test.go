package standard

import (
	"auth/service/database/types"
	"auth/service/tools/password"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

/**
I read that the tests are run concurrently.
But for some reason, it seems like they are run
sequentially...
Still, each test setups and tear downs the test user.
This is the design decision I made with my current information.
*/

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

	db.Where("id = ?", user.ID).Delete(&types.User{})
}

func TestEmailPasswordLoginHandler_Correct(t *testing.T) {
	user := types.User{
		Email:        "test@example.com",
		PasswordHash: password.Encrypt("password123"),
	}
	err := db.Create(&user).Error
	assert.NoError(t, err, "Failed to create test user for login test")

	// Create a request body
	requestBody := map[string]string{
		"email":    "test@example.com",
		"password": "password123",
	}
	jsonRequestBody, _ := json.Marshal(requestBody)

	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonRequestBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := testServer.NewContext(req, rec)

	// Call the handler
	err = EmailPasswordLoginHandler(c)

	// Check if there was an error
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	db.Where("id = ?", user.ID).Delete(&types.User{})
}

func TestEmailPasswordLoginHandler_Wrong(t *testing.T) {
	user := types.User{
		Email:        "test@example.com",
		PasswordHash: password.Encrypt("password123"),
	}
	err := db.Create(&user).Error
	assert.NoError(t, err, "Failed to create test user for login test")

	// Create a request body
	requestBody := map[string]string{
		"email":    "test2@example.com",
		"password": "password1234",
	}
	jsonRequestBody, _ := json.Marshal(requestBody)

	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonRequestBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := testServer.NewContext(req, rec)

	// Call the handler
	err = EmailPasswordLoginHandler(c)

	// Check if there was an error
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.NoError(t, err)

	db.Where("id = ?", user.ID).Delete(&types.User{})
}

func TestCookieLoginHandler(t *testing.T) {
	user := types.User{
		Email:        "test@example.com",
		PasswordHash: password.Encrypt("password123"),
	}
	err := db.Create(&user).Error
	userID := user.ID
	assert.NoError(t, err, "Failed to create test user")

	// Create a request body
	requestBody := map[string]string{
		"email":    "test@example.com",
		"password": "password123",
	}
	jsonRequestBody, _ := json.Marshal(requestBody)
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonRequestBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := testServer.NewContext(req, rec)

	// Call the handler
	err = EmailPasswordLoginHandler(c)
	assert.NoError(t, err, "Failed to login")
	jwt := rec.Result().Cookies()[0]

	req = httptest.NewRequest(http.MethodGet, "/login", bytes.NewBuffer(jsonRequestBody))
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(jwt)
	rec = httptest.NewRecorder()
	c = testServer.NewContext(req, rec)
	err = CookieLoginHandler(c)
	assert.NoError(t, err, "Failed to get user from cookie")

	user = types.User{}
	json.Unmarshal(rec.Body.Bytes(), &user)

	assert.Equal(t, userID, user.ID)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NoError(t, err)

	db.Where("id = ?", user.ID).Delete(&types.User{})
}
