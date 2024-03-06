package handler

import (
	"github.com/admin-srf/go_crud/config"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitHandler() {
	db = config.GetDb()
}
