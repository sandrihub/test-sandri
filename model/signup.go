package model

import "time"

type Signup struct {
	Id        	uint 		`json:"id",gorm:"primary_key"`
	Name		string 		`json:"name"`
	Email		string		`json:"email"`
	Password	string		`json:"password"`
	CreatedAt 	*time.Time     	`json:"created_at"`
	UpdatedAt 	*time.Time 	`json:"updated_at"`
}
