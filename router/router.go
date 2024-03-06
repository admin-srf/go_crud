package router

import "github.com/gin-gonic/gin"

func Initialize() {

	router := gin.Default()
	initializerRoutes(router)
	router.Run()
}
