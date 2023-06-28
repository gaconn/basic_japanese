package utils

import "github.com/gin-gonic/gin"

type Response struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Cache bool        `json:"cache"`
}

func SendResponse(r *gin.Context, code int, msg string, data interface{}, cache ...bool) {
	var isCache = false
	if len(cache) > 0 {
		isCache = cache[0]
	}
	r.JSON(code, Response{Code: code, Msg: msg, Data: data, Cache: isCache})
}

func (r *Response) SendResponse(c *gin.Context) {
	c.JSON(r.Code, r)
}
