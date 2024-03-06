package handler

import (
	"github.com/admin-srf/go_crud/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Users
// @Description List all users on platform
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {array} schemas.User
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /users/ [get]
func ListUsersHandler(c *gin.Context) {

	users := []schemas.User{}
	if err := db.Find(&users).Error; err != nil {
		validationException(c, err)
		return
	}

	sendSuccess(c, "list-user", users)

}
