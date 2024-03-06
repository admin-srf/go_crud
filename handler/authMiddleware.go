package handler

import (
	"net/http"

	"github.com/admin-srf/go_crud/services"
	"github.com/gin-gonic/gin"
)

func VerifyJwt(c *gin.Context) {
	jwtToken, err := services.ExtractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse{Message: "Must provide Bearer token authorization"})
		return
	}

	claims, err := services.ParseToken(jwtToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{Message: "Bad jwt token"})
		return
	}

	if claims.UserID == 0 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorResponse{
			Message: "unable to parse claims",
		})
		return
	}
	c.Next()

}
