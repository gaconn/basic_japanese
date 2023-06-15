package utils

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/quan12xz/basic_japanese/pkg/setting"
	"log"
)

var db *gorm.DB

func DBConnect() {
	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Name,
		setting.DatabaseSetting.Charset,
	))

	if err != nil {
		log.Fatal(err)
	}

}
