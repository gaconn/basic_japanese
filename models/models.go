package models

import (
	"fmt"
	"log"
	"time"

	"github.com/quan12xz/basic_japanese/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var NumberRecordLimit = 30
var retryNumber = 4

type mysqlConnect struct {
	DB *gorm.DB
}

var dbInstance *mysqlConnect

func GetInstance() *mysqlConnect {
	if dbInstance == nil {
		dbInstance = &mysqlConnect{
			DB: dbSetup(),
		}
	}
	return dbInstance
}

func dbSetup() *gorm.DB {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Port,
		setting.DatabaseSetting.Name,
		setting.DatabaseSetting.Charset,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	for err != nil {
		log.Fatal(err)
		if retryNumber > 1 {
			log.Fatal("Trying to reconnect")
			retryNumber--
			time.Sleep(3)
			db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
			continue
		}
		panic("Unable to connect database")
	}

	return db
}
