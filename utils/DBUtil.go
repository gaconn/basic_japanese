package utils

import (
	"fmt"
	"github.com/quan12xz/basic_japanese/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var DB *gorm.DB
var retryNumber = 3

func DBSetup() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Port,
		setting.DatabaseSetting.Name,
		setting.DatabaseSetting.Charset,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	for err != nil {
		log.Fatal(err)
		if retryNumber > 1 {
			log.Fatal("Trying to reconnect")
			retryNumber--
			time.Sleep(3)
			DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
			continue
		}
		panic("Unable to connect database")
	}
}
