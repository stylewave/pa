package main

import (
	"fmt"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

func Init()  {
	DB,err := gorm.Open("mysql","root:root@tcp(localhost:3306)/spiders?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf(" gorm.Open.err: %v", err)
	}

	DB.SingularTable(true)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "sp_" + defaultTableName
	}
}

func main()  {

	DB,err := gorm.Open("mysql","root:root@tcp(localhost:3306)/spiders?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf(" gorm.Open.err: %v", err)
	}
	fmt.Println(err)

	fmt.Println( model.DB)

	DB.SingularTable(true)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "sp_" + defaultTableName
	}
	fmt.Println("22222222222")
}
