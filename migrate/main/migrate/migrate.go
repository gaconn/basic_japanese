package main

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/quan12xz/basic_japanese/migrate/migrateSetting"
	migrateModel "github.com/quan12xz/basic_japanese/migrate/models"
	"github.com/quan12xz/basic_japanese/models"
	"log"
)

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
var config *ini.File

func Setup() {
	var err error
	config, err = ini.Load("../../../config/app.ini")

	if err != nil {
		log.Fatal(err)
	}
	mapTo("database", DatabaseSetting)
}
func mapTo(name string, v interface{}) {
	err := config.Section(name).MapTo(v)

	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	migrateSetting.Setup()
	migrateModel.DBSetup()
	migrateModel.DB.AutoMigrate(&models.Alphabet{})
	fmt.Println("Hope all good :))")
}
