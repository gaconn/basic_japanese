package utils

import "github.com/gin-gonic/gin"

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SendResponse(r *gin.Context, code int, msg string, data interface{}) {
	r.JSON(code, Response{Code: code, Msg: msg, Data: data})
}
