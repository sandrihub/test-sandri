package view

type DataProvider struct {
	ProviderCode		string  	`json:"provider_code"`
	Name    		string 		`json:"name"`
	Address    		string 		`json:"address"`
	Phone			string 		`json:"phone"`
	Email			string 		`json:"email"`
	Owner 			string 		`json:"owner"`
	Latitude		string  	`json:"latitude"`
	Longitude		string  	`json:"longitude"`
	Photos			[]PhotoView     `json:"photos"`
}

type ProviderView struct {
	Id 			int     `json:"id"`
	ProviderCode		string  `json:"provider_code"`
	ProviderName    	string 	`json:"provider_name"`
	ProviderAddress    	string 	`json:"provider_address"`
	ProviderPhone		string 	`json:"provider_phone"`
	ProviderEmail		string  `json:"provider_email"`
	ProviderOwner 		string 	`json:"proivder_owner"`
	Latitude		string  `json:"latitude"`
	Longitude		string  `json:"longitude"`
}

