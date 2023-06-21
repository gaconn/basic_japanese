package main

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/quan12xz/basic_japanese/migrate/migrateSetting"
	migrateModel "github.com/quan12xz/basic_japanese/migrate/models"
	"github.com/quan12xz/basic_japanese/models"
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

func main() {
	migrateSetting.Setup()
	migrateModel.DBSetup()
	migrateModel.DB.AutoMigrate(&models.Alphabet{})
	migrateModel.DB.Migrator().DropTable(&models.Sentence{}, &models.Word{}, &models.Lesson{})
	migrateModel.DB.AutoMigrate(&models.Sentence{}, &models.Word{}, &models.Lesson{})
	fmt.Println("Hope all good :))")
}
