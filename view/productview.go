package view

import "time"

type ProductView struct {
	Id        	uint 		`json:"id"`
	Name		string		`json:"name"`
	Price		int32		`json:"price"`
	Imageurl	string		`json:"imageurl"`
	CreatedAt 	*time.Time     	`json:"created_at"`
	UpdatedAt 	*time.Time 	`json:"updated_at"`
}
