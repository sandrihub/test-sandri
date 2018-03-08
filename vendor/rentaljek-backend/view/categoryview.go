package view

type CategoryView struct {
	Base
	VehicleCode	string	`json:"vehicle_code"`
	Name		string  `json:"name"`
	Note		string  `json:"note"`
}
