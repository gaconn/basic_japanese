package controller

import "github.com/gin-gonic/gin"

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func GetKatakana(r *gin.Context) {
	r.JSON(200, Response{200, "Successfully", nil})
}
