package utils

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/quan12xz/basic_japanese/pkg/setting"
	"log"
	"time"
)

type DBUtil struct {
	DB *gorm.DB
}

var DBUtilInstance = &DBUtil{}

func DBSetup() {
	var err error
	DBUtilInstance.DB, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Name,
		setting.DatabaseSetting.Charset,
	))

	if err != nil {
		log.Fatal(err)
	}

	DBUtilInstance.DB.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampCreateCallback)
}

func updateTimeStampCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()

		if createTimeField, ok := scope.FieldByName("CreateOn"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if updateTimeField, ok := scope.FieldByName("UpdateOn"); ok {
			if updateTimeField.IsBlank {
				updateTimeField.Set(nowTime)
			}
		}
	}
}
