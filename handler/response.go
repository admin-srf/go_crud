package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendError(c *gin.Context, code int, message string) {
	c.Header("Content-type", "application/json")
	c.JSON(code, ErrorResponse{
		Message: message,
	})

}

func validationException(c *gin.Context, err error) {
	c.Header("Content-type", "application/json")
	c.JSON(http.StatusBadRequest, gin.H{"detail": err.Error()})

}

func sendSuccess(c *gin.Context, op string, data interface{}) {
	c.Header("Content-type", "application/json")
	c.JSON(http.StatusOK, data)
}

// func sendSuccessMessage(c *gin.Context, op string, message string) {
// 	c.Header("Content-type", "application/json")
// 	c.JSON(http.StatusOK, gin.H{"detail": message})
// }

type ErrorResponse struct {
	Message string `json:"message"`
}
