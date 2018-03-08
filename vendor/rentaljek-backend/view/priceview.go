package view

type Prices struct {
	InitialPrice	int32 `json:"initial_price"`
	Discount	int32 `json:"discount"`
	ActualPrice	int32 `json:"actual_price"`
}
