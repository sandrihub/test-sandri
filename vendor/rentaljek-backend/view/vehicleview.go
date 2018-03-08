package view


type VehicleView struct {
	Base
	CategoryCode    	string		`json:"category_code"`
	CategoryId		int             `json:"category_id"`
	CategoryName		string          `json:"category_name"`
	LicensePlate		string		`json:"license_plate"`
	VehicleCode		string		`json:"vehicle_code"`
	Manufacturer		string		`json:"manufacturer"`
	Name			string		`json:"name"`
	Type			string		`json:"type"`
	Fuel			string		`json:"fuel"`
	Year			string		`json:"year"`
	Capacity		int 		`json:"capacity"`
	Price			struct{ Daily	Prices `json:"daily"`
					     Weekly	Prices `json:"weekly"`
					     Monthly	Prices `json:"monthly"`
					     CancelationFee	float32			`json:"cancelation_fee"`
					     PaymentTypes	[]PaymentTypeView  	`json:"payment_type"`
			    } `json:"price"`

	Photos			[]PhotoView         `json:"photos"`
	Rating			int		`json:"rating"`
	Review			int		`json:"review"`
}
