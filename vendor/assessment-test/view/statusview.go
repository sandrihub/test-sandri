package view

type StatusView struct {
	Status		string	`json:"status"`
	Result		interface{}	`json:"result"`
	Error		interface{}	`json:"error"`
}
