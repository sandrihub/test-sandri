package model

import (
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func SetDatabase(db *gorm.DB) {
	DB = db

	db.AutoMigrate(&Vehicle{},&Photo{},&Category{},&PaymentType{},&Order{},&Customer{},&Provider{})
}

