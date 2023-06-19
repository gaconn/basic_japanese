package utils

import (
	"github.com/quan12xz/basic_japanese/pkg/setting"
	"log"
	"os"
	"runtime"
)

func WriteErrorLog(message error) {
	f, err := os.OpenFile(setting.AppSetting.LogSavePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
	_, file, line, _ := runtime.Caller(1)
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString("File: " + file + "----Line: " + string(line) + "---message: " + message.Error() + "\n")
	f.Close()
}
