package model

import (
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	OrderCode	string        	`json:"order_code"`
	CustomerId	int		`json:"customer_id"`
	VehicleId	int		`json:"vehicle_id"`
	VehicleCode	string		`json:"vehicle_code",gorm:"size:10"`
	StartDate	string		`json:"start_date"`
	Periode		string		`json:"periode"`
	PeriodeUnit	string		`json:"periode_unit"`
	Driver 		bool		`json:"driver"`
	PickupPoint	string          `json:"pickup_point"`
	PickupAddress	string		`json:"pickup_address"`
	PickupTime	string		`json:"pickup_time"`
	PickupNote	string		`json:"pickup_note"`
	PriceType	string		`json:"price_type"`
	InitialPrice	int32		`json:"initial_price"`
	ActualPrice	int32		`json:"actual_price"`
	Discount	int32           `json:"discount"`
	Message 	string		`json:"message"`
	Rate 		int		`json:"rate"`
}
