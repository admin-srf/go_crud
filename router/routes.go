package router

import (
	"github.com/gin-gonic/gin"

	"github.com/admin-srf/go_crud/handler"

	docs "github.com/admin-srf/go_crud/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializerRoutes(router *gin.Engine) {
	handler.InitHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)

	{
		v1.POST("/auth", handler.AuthHandler)

		v1.POST("/signup", handler.SignUpHandler)
	}

	private := router.Group(basePath)
	private.Use(handler.VerifyJwt)
	{
		private.GET("/users/:userId", handler.GetUserHandler)
		private.GET("/users/", handler.ListUsersHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
