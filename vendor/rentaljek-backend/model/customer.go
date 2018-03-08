package model

import (
	"github.com/jinzhu/gorm"
)

type Customer struct {
	gorm.Model
	Username	string  	`json:"username"`
	Password	string 		`json:"password"`
	Email		string		`json:"email"`
	Fullname	string		`json:"fullname"`
	//Photos		[]PhotoCustomer `json:"photos"`
	Rating		float32		`json:"rating"`
	Phone		string		`json:"phone"`
	
}

//type PhotoCustomer struct
//{
//	Caption		string	`json:"caption"`
//	Src		string	`json:"src"`
//}


