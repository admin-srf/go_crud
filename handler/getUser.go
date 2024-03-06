package handler

import (
	"net/http"

	"github.com/admin-srf/go_crud/schemas"
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
		sendError(c, http.StatusBadRequest, errParamIsRequired("userId", "int").Error())
		return
	}
	var user schemas.User
	if err := db.First(&user, "id = ?", userId).Error; err != nil {
		validationException(c, err)
		return
	}

	sendSuccess(c, "detail-user", user)

}
