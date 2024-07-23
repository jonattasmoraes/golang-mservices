package domain

import "time"

type SaleReport struct {
	UserID      string
	SaleID      string
	PaymentType string
	SaleDate    time.Time `json:"date"`
	Total       int
	Products    []ProductReport
}

type ProductReport struct {
	ProductID string
	Name      string
	Unit      string
	Category  string
	Quantity  int
	Price     int
}
