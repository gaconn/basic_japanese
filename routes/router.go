package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/quan12xz/basic_japanese/controller"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	kataGroup := r.Group("/word/katakana")
	{
		kataGroup.GET("/", controller.GetAllKatakana)
		kataGroup.GET("/{id}", controller.GetByIDKatakana)
		kataGroup.POST("/", controller.AddKatakana)
		kataGroup.DELETE("/", controller.DeleteKatakana)
		kataGroup.PUT("/", controller.UpdateKatakana)
	}

	return r
}
