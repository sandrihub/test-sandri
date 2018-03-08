package config

import (
	"github.com/jinzhu/gorm"
	"log"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/assessment?charset=utf8&parseTime=True&loc=Local")
	log.Println("db")
	if err!=nil {
		log.Println(err)
		panic(err)
	}


	db.LogMode(true)
	log.Println(db)
	return db
}

