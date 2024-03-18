package roles

import (
	"auth/service/database/types"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRoles(t *testing.T) {
	roles := []types.Role{
		{Name: "role1"},
		{Name: "role2"},
	}
	for _, role := range roles {
		db.Create(&role)
	}

	defer func() {
		db.Where("Name = ?", "role1").Delete(&types.Role{})
		db.Where("Name = ?", "role2").Delete(&types.Role{})
	}()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := testServer.NewContext(req, rec)

	err := GetRoles(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	b := []map[string]string{}
	json.Unmarshal(rec.Body.Bytes(), &b)
	assert.Equal(t, 2, len(b))
	assert.Equal(t, "role1", b[0]["name"])
	assert.Equal(t, "role2", b[1]["name"])
}

func TestAddRole(t *testing.T) {
	defer func() {
		db.Where("Name = ?", "role").Delete(&types.Role{})
	}()

	role := types.Role{
		Name: "role",
	}
	jsonBody, _ := json.Marshal(role)

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := testServer.NewContext(req, rec)
	err := AddRole(c)
	assert.NoError(t, err)

	role = types.Role{}
	db.First(&role)

	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Equal(t, "role", role.Name)
}

func TestDeleteRole(t *testing.T) {
	role := types.Role{
		Name: "role",
	}
	db.Create(&role)
	id := role.ID
	defer func() {
		db.Unscoped().Where("id = ?", id).Delete(&types.Role{})
	}()

	jsonBody, _ := json.Marshal(role)
	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/%s", role.ID), bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := testServer.NewContext(req, rec)
	err := DeleteRole(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, rec.Code)

	role = types.Role{}
	err = db.Where("id = ?", id).First(&role).Error
	assert.Equal(t, err, gorm.ErrRecordNotFound)
}

func TestGetUserRoles(t *testing.T) {
	// TODO
}

func TestAddUserRoles(t *testing.T) {
	// TODO
}

func TestDeleteUserRoles(t *testing.T) {
	// TODO
}
