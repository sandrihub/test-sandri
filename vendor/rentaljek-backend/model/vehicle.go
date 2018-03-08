package model

import "github.com/jinzhu/gorm"

type Vehicle struct {
	gorm.Model
	VehicleCode	string 	`json:"vehicle_code",gorm:"size:10;primary_key;unique"`
	CategoryCode	string 	`json:"category_code",gorm:"size:10;index:idx_vehicles_category_code"`
	LicensePlate	string 	`json:"license_plate",gorm:"size:25"`
	Manufacturer	string 	`json:"manufacturer",gorm:"size:50"`
	Name		string 	`json:"name",gorm:"size:80"`
	Type		string 	`json:"type",gorm:"size:20"`
	Fuel		string 	`json:"fuel",gorm:"size:20"`
	Year		string 	`json:"year",gorm:"size:4"`
	Capacity	int8	`json:"capacity"`
	InitialPriceDaily	int32	`json:"initial_price_daily"`
	InitialPriceWeekly	int32	`json:"initial_price_weekly"`
	InitialPriceMonthly	int32   `json:"initial_price_monthly"`
	DiscountDaily		int32	`json:"discount_daily"`
	DiscountWeekly		int32	`json:"discount_weekly"`
	DiscountMonthly		int32	`json:"discount_monthly"`
	ActualPriceDaily	int32	`json:"actual_price_daily"`
	ActualPriceWeekly	int32	`json:"actual_price_weekly"`
	ActualPriceMonthly	int32	`json:"actual_price_monthly"`
	CancelationFee		float32	`json:"cancelation_fee"`
	Rating		int8	`json:"rating"`
	Review		int8	`json:"review"`
	Photos		[]Photo `json:"photos",gorm:"ForeignKey:VehicleCode;AssociationForeignKey:VehicleCode"`
	PaymentTypes	[]PaymentType `json:"payment_types",gorm:"ForeignKey:VehicleCode;AssociationForeignKey:VehicleCode"`
}
