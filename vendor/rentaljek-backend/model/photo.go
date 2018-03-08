package model

import "github.com/jinzhu/gorm"

type Photo struct {
	gorm.Model
	VehicleCode	string 	`json:"vehicle_code,"gorm:"size:10"`
	Caption		string	`json:"caption"`
	Source		string	`json:"source"`
}
