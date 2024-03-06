package handler

import (
	"github.com/admin-srf/go_crud/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary SignUp
// @Description Create new user with password hashed
// @Tags auth
// @Accept json
// @Produce json
// @Param request body CreateUserRequest true "Request body"
// @Success 200 {object} UserResponse
// @Failure 400 {object} ErrorResponse
// @Router /signup [post]
func SignUpHandler(c *gin.Context) {

	request := CreateUserRequest{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		validationException(c, err)
		return
	}

	hashedPassword, err := GeneratePassword(request.Password)

	if err != nil {
		validationException(c, err)
		return
	}

	user := schemas.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: hashedPassword,
	}

	if err := db.Create(&user).Error; err != nil {
		validationException(c, err)
		return
	}

	sendSuccess(c, "post", user)

}
