package migrateModel

import (
	"fmt"
	"github.com/quan12xz/basic_japanese/migrate/migrateSetting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func DBSetup() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		migrateSetting.DatabaseSetting.User,
		migrateSetting.DatabaseSetting.Password,
		migrateSetting.DatabaseSetting.Host,
		migrateSetting.DatabaseSetting.Port,
		migrateSetting.DatabaseSetting.Name,
		migrateSetting.DatabaseSetting.Charset,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
}
