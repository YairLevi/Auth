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
	for i := range roles {
		db.Create(&roles[i])
	}
	defer func() {
		for i := range roles {
			db.Unscoped().Where("id = ?", roles[i].ID).Delete(&types.Role{})
		}
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
	originalRole := types.Role{
		Name: "role",
	}
	t.Cleanup(func() {
		db.Unscoped().Where("name = ?", originalRole.Name).Delete(&types.Role{})
	})
	jsonBody, _ := json.Marshal(originalRole)

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := testServer.NewContext(req, rec)
	err := AddRole(c)
	assert.NoError(t, err)

	role := types.Role{}
	json.Unmarshal(rec.Body.Bytes(), &role)
	db.Where("id = ?", role.ID).First(&role)

	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Equal(t, originalRole.Name, role.Name)
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

	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	values := req.URL.Query()
	values.Add("roleId", role.ID)
	req.URL.RawQuery = values.Encode()
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
	user := types.User{}
	db.First(&user)
	role := types.Role{
		Name: "role10",
	}
	assert.NoError(t, db.Create(&role).Error)
	assert.NoError(t, db.Create(&types.UserRole{UserID: user.ID, RoleID: role.ID}).Error)
	defer func() {
		db.Where("name = ?", role.Name).Delete(&types.Role{})
		db.Where("username = ?", user.Username).Delete(&types.User{})
	}()

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/%s", user.ID), nil)
	rec := httptest.NewRecorder()
	c := testServer.NewContext(req, rec)
	err := GetUserRoles(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, `["role10"]`+"\n", rec.Body.String())
}

func TestAddUserRoles(t *testing.T) {
	// TODO
}

func TestDeleteUserRoles(t *testing.T) {
	// TODO
}
