package router

import (
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/zthiagovalle/gopportunities/docs"
	"github.com/zthiagovalle/gopportunities/handler"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize Handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET("/opening", handler.ShowOpeningHanlder)
		v1.POST("/opening", handler.CreateOpeningHanlder)
		v1.DELETE("/opening", handler.DeleteOpeningHanlder)
		v1.PUT("/opening", handler.UpdateOpeningHanlder)
		v1.GET("/openings", handler.ListOpeningsHanlder)
	}

	// Initialize swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
