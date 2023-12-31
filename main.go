package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quan12xz/basic_japanese/pkg/setting"
	"github.com/quan12xz/basic_japanese/routes"
)

func init() {
	setting.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	handler := routes.InitRouter()
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	maxHeaderByte := 1 << 20
	server := &http.Server{
		Addr:           endPoint,
		Handler:        handler,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderByte,
	}
	log.Printf("[info] start http listening on port %s", endPoint)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
