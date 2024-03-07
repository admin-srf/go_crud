package handler

import (
	"github.com/admin-srf/go_crud/config"
	"github.com/admin-srf/go_crud/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitHandler() {
	db = config.GetDb()
}

func GetUser(c *gin.Context) (userId int, err error) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		return 0, services.MissingValue("userID")
	}

	userIDUint, ok := userIDInterface.(uint)
	if !ok {
		return 0, services.CustomError("Cannot pass userId")
	}

	return int(userIDUint), nil
}
