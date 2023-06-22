package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type App struct {
	PrefixUrl        string
	LogSavePath      string
	LogSaveName      string
	LogFileExtension string
	TimeFormat       string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type     string
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Charset  string
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

var RedisSetting = &Redis{}
var config *ini.File

func Setup() {
	var err error
	config, err = ini.Load("config/app.ini")

	if err != nil {
		log.Fatal(err)
	}
	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("redis", RedisSetting)
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
}

func mapTo(name string, v interface{}) {
	err := config.Section(name).MapTo(v)

	if err != nil {
		log.Fatal(err)
	}
}
