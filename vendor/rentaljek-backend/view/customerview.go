package view

import "time"

type Customer struct {
	Id 		int 		`json:"id"`
	Username	string  	`json:"username"`
	CreatedAt	time.Time	`json:"created_at"`
	Email		string		`json:"email"`
	Fullname	string		`json:"fullname"`
	Photos		[]PhotoCustomer `json:"photos"`
	Rating		float32		`json:"rating"`
}

type PhotoCustomer struct
{
	Caption		string	`json:"caption"`
	Src		string	`json:"src"`
}