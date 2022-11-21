package v1

import (
	"base_service/docs"
	"github.com/gin-gonic/gin"
	"github.com/gogovan-korea/ggx-kr-service-utils/tracing"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func MapRoutes(router *gin.Engine, userHandler *UserHandler) {
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Title = "S14E API"
	docs.SwaggerInfo.Description = "S14E API"
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.Use(tracing.CreateNewRelicGinMiddleware())
	v1 := router.Group("/api/v1")
	{
		v1.POST("/users", userHandler.GetUser)
		v1.POST("/test", userHandler.Test)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
