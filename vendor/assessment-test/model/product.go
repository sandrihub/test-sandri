package model

import "time"

type Product struct {
	//gorm.Model
	Id        	uint `json:"id",gorm:"primary_key"`
	Name		string 	`json:"name",gorm:"size:80"`
	Price		int32	`json:"price"`
	Imageurl	string	`json:"imageurl"`
	CreatedAt 	*time.Time     `json:"created_at"`
	UpdatedAt 	*time.Time 	`json:"updated_at"`
}
