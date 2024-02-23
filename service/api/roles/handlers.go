package roles

import (
	"auth/service/database"
	"auth/service/database/types"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetRoles(ctx echo.Context) error {
	var roles []types.Role
	if err := database.DB.Preload("UserRoles").Find(&roles).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, &roles)
}

func AddRole(ctx echo.Context) error {
	role := types.Role{}
	if err := ctx.Bind(&role); err != nil {
		return ctx.JSON(http.StatusBadRequest, "error in payload")
	}

	if err := database.DB.Create(&role).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, "error adding role")
	}

	return ctx.JSON(http.StatusOK, &role)
}

func DeleteRole(ctx echo.Context) error {
	roleName := ctx.Param("role")
	role := types.Role{
		Name: roleName,
	}
	if err := database.DB.Unscoped().Where(&role).Delete(&role).Error; err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusInternalServerError, "error deleting")
	}

	return ctx.NoContent(http.StatusNoContent)
}

func GetUserRoles(ctx echo.Context) error {
	userID := ctx.Param("userId")

	var userRoles []types.UserRole
	if err := database.DB.Where(&types.UserRole{UserID: userID}).Find(&userRoles).Error; err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, "error getting user role IDs")
	}

	roleIDs := make([]string, 0)
	for _, userRole := range userRoles {
		roleIDs = append(roleIDs, userRole.RoleID)
	}

	roles := make([]types.Role, 0)
	if err := database.DB.Where("id IN (?)", roleIDs).Find(&roles).Error; err != nil {
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
	if err := database.DB.Where(&role).Find(&role).Error; err != nil {
		return ctx.JSON(http.StatusBadRequest, fmt.Sprint("error no such role", dto.Role, "for app"))
	}

	userRole := types.UserRole{
		RoleID: role.ID,
		UserID: userID,
	}
	if err := database.DB.Create(&userRole).Error; err != nil {
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
	if err := database.DB.Find(&role).Error; err != nil {
		return ctx.JSON(http.StatusBadRequest, fmt.Sprint("error no such role", dto.Role, "for app"))
	}

	userRole := types.UserRole{
		RoleID: role.ID,
		UserID: userID,
	}

	if err := database.DB.Unscoped().Where(&userRole).Delete(&userRole).Error; err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, fmt.Sprint("failed to revoke role", role.Name, "from user."))
	}
	return ctx.NoContent(http.StatusNoContent)
}
