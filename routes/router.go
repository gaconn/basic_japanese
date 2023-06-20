package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/quan12xz/basic_japanese/controller"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	kataGroup := r.Group("/word")
	{
		kataGroup.GET("/type/:type", controller.GetAllByType)
		kataGroup.GET("/:id", controller.GetByID)
		kataGroup.PUT("/", controller.Update)
	}

	return r
}
