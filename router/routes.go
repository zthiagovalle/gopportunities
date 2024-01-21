package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zthiagovalle/gopportunities/handler"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOpeningHanlder)
		v1.POST("/opening", handler.CreateOpeningHanlder)
		v1.DELETE("/opening", handler.DeleteOpeningHanlder)
		v1.PUT("/opening", handler.UpdateOpeningHanlder)
		v1.GET("/openings", handler.ListOpeningsHanlder)
	}
}
