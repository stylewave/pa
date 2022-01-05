package model

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB

	username string = "root"
	password string = "root"
	dbName   string = "spiders"
)

func init() {
	var err error
//	DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbName))
	//DB, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/spiders?charset=utf8")
	DB,err := gorm.Open("mysql","root:root@tcp(localhost:3306)/spiders?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf(" gorm.Open.err: %v", err)
	}

	DB.SingularTable(true)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "sp_" + defaultTableName
	}
}
