package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/quan12xz/basic_japanese/controller"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", controller.GetKatakana)

	return r
}
