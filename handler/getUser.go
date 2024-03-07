package handler

import (
	"net/http"

	"github.com/admin-srf/go_crud/schemas"
	"github.com/admin-srf/go_crud/services"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary User
// @Description Detail of user on platform
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} schemas.User
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /users/:userId [get]
func GetUserHandler(c *gin.Context) {

	userId := c.Param("userId")
	if userId == "" {
		sendError(c, http.StatusBadRequest, services.ErrParamIsRequired("userId", "int").Error())
		return
	}
	var user schemas.User
	if err := db.First(&user, "id = ?", userId).Error; err != nil {
		validationException(c, err)
		return
	}

	sendSuccess(c, "detail-user", user)

}

// @BasePath /api/v1
// @Summary Users
// @Description List all users on platform
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {array} schemas.User
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /me/ [get]
func GetCurrentUsersHandler(c *gin.Context) {
	userId, err := GetUser(c)
	if err != nil {
		validationException(c, err)
		return
	}
	user := schemas.User{}
	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		validationException(c, err)
		return
	}

	sendSuccess(c, "get-current-user", user)

}
