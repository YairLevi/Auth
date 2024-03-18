package roles

import (
	"auth/service/database"
	"auth/service/database/types"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	db = database.DB
)

func GetRoles(ctx echo.Context) error {
	var roles []types.Role
	if err := db.Preload("UserRoles").Find(&roles).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, &roles)
}

func AddRole(ctx echo.Context) error {
	role := types.Role{}
	if err := ctx.Bind(&role); err != nil {
		return ctx.JSON(http.StatusBadRequest, "error in payload")
	}

	if err := db.Create(&role).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, "error adding role")
	}

	return ctx.JSON(http.StatusCreated, &role)
}

func DeleteRole(ctx echo.Context) error {
	roleName := ctx.Param("role")
	if err := db.Unscoped().Where("name = ?", roleName).Delete(&types.Role{}).Error; err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusInternalServerError, "error deleting")
	}

	return ctx.NoContent(http.StatusNoContent)
}

func GetUserRoles(ctx echo.Context) error {
	userID := ctx.Param("userId")

	var userRoles []types.UserRole
	if err := db.Where(&types.UserRole{UserID: userID}).Find(&userRoles).Error; err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, "error getting user role IDs")
	}

	roleIDs := make([]string, 0)
	for _, userRole := range userRoles {
		roleIDs = append(roleIDs, userRole.RoleID)
	}

	roles := make([]types.Role, 0)
	if err := db.Where("id IN (?)", roleIDs).Find(&roles).Error; err != nil {
		return ctx.JSON(http.StatusBadRequest, "error getting user roles")
	}

	roleNames := make([]string, 0)
	for _, role := range roles {
		roleNames = append(roleNames, role.Name)
	}

	return ctx.JSON(http.StatusOK, roleNames)
}

func AssignRoleToUser(ctx echo.Context) error {
	userID := ctx.Param("userId")

	dto := struct {
		Role string `json:"role"`
	}{}
	if err := ctx.Bind(&dto); err != nil {
		return ctx.JSON(http.StatusBadRequest, "error in payload")
	}

	role := types.Role{
		Name: dto.Role,
	}
	if err := db.Where(&role).Find(&role).Error; err != nil {
		return ctx.JSON(http.StatusBadRequest, fmt.Sprint("error no such role", dto.Role, "for app"))
	}

	userRole := types.UserRole{
		RoleID: role.ID,
		UserID: userID,
	}
	if err := db.Create(&userRole).Error; err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusInternalServerError, "error assign role to user")
	}

	return ctx.NoContent(http.StatusCreated)
}

func RevokeRoleFromUser(ctx echo.Context) error {
	userID := ctx.Param("userId")

	dto := struct {
		Role string `json:"role"`
	}{}
	if err := ctx.Bind(&dto); err != nil {
		return ctx.JSON(http.StatusBadRequest, "error in payload")
	}

	role := types.Role{
		Name: dto.Role,
	}
	if err := db.Find(&role).Error; err != nil {
		return ctx.JSON(http.StatusBadRequest, fmt.Sprint("error no such role", dto.Role, "for app"))
	}

	userRole := types.UserRole{
		RoleID: role.ID,
		UserID: userID,
	}

	if err := db.Unscoped().Where(&userRole).Delete(&userRole).Error; err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, fmt.Sprint("failed to revoke role", role.Name, "from user."))
	}
	return ctx.NoContent(http.StatusNoContent)
}
