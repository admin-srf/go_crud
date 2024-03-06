package handler

import (
	"net/http"

	"github.com/admin-srf/go_crud/schemas"
	"github.com/admin-srf/go_crud/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserResponse struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

// @BasePath /api/v1
// @Summary Authenticate
// @Description Authenticate email and password for token generate
// @Tags auth
// @Accept json
// @Produce json
// @Param request body AuthenticateUserRequest true "Request body"
// @Success 200 {object} UserResponse
// @Failure 400 {object} ErrorResponse
// @Router /auth [post]
func AuthHandler(c *gin.Context) {

	request := AuthenticateUserRequest{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		validationException(c, err)
		return
	}

	user := &schemas.User{}

	if err := db.Where("email = ?", request.Email).First(user).Error; err != nil {
		validationException(c, err)
		return
	}

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword {
		sendError(c, http.StatusBadRequest, "Invalid login credentials. Please try again")
		return
	}

	tokenString, err := services.NewAccessToken(*user)
	if err != nil {
		validationException(c, err)
		return
	}

	sendSuccess(c, "authentication", gin.H{
		"accessToken": tokenString,
		"user":        user,
	})

}
