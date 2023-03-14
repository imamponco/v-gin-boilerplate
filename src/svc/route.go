package svc

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoute(router *gin.Engine, controllers *Controllers) *gin.Engine {
	// Get health API
	router.GET("/", controllers.GetHealth.GetHealthAPI)

	// Open API 3.0 Swagger
	router.GET("/documentation/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
