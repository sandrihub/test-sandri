package model

import "github.com/jinzhu/gorm"

type PaymentType struct {
	gorm.Model
	VehicleCode	string	`json:"vehicle_code,"gorm:"size:10"`
	Name		string	`json:"name"`
	Note		string	`json:"note"`
}
