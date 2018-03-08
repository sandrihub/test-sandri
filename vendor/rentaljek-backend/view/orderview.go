package view

type DataOrder struct {
	StartOrder    	string		`json:"start_order"`
	FinishOrder    	string		`json:"finish_order"`
	OrderId		int             `json:"order_id"`
	Rate 		int 		`json:"rate"`
	Message 	string          `json:"message"`
}

type OrderView struct {
	CutomerId    		int		`json:"cutomer_id"`
	HistoryOrder		struct{ Count 	int        	`json:"count"`
					Data   	[]DataOrder	`json:"data"`
				      } `json:"history_order"`
}
