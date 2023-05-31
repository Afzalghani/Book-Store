package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {

	d, err := gorm.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/book_store?charset=utf8mb4&parseTime=True&loc=Local")
	//d, err = gorm.Open("mysql", ":password@/dbname?charset=utf8&parseTime=True&loc=Local")

	//d, err := gorm.Open("user:admin@tcp(127.0.0.1:3306)/book_store?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Println("This has error")
		panic("failed to connect database")
	}

	db = d
}
func GetDb() *gorm.DB {
	return db
}
