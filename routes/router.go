package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/quan12xz/basic_japanese/controller"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	alphabetGroup := r.Group("/word")
	{
		alphabetGroup.GET("/type/:type", controller.GetAllByType)
		alphabetGroup.GET("/:id", controller.GetByID)
		alphabetGroup.PUT("/", controller.Update)
	}

	lessonGroup := r.Group("/lesson")
	{
		lessonGroup.GET("/", controller.GetLessons)
		lessonGroup.GET("/:id", controller.GetLesson)
		lessonGroup.POST("/", controller.AddLesson)
		lessonGroup.PUT("/", controller.UpdateLesson)
		lessonGroup.DELETE("/:id", controller.DeleteLesson)
	}
	r.GET("/checkcache", controller.CheckCache)
	return r
}
