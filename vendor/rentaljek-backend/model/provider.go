package model

import (
	"github.com/jinzhu/gorm"
)

type Provider struct {
	gorm.Model
	ProviderCode	string        	`json:"provider_code"`
	Name		string  	`json:"name"`
	Address		string 		`json:"address"`
	Email		string		`json:"email"`
	Owner		string		`json:"owner"`
	Photos		[]Photo 	`json:"photos"`
	Phone		string		`json:"phone"`
	Latitude	string          `json:"latitude"`
	Longitude	string          `json:"longitude"`
}



