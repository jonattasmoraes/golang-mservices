package domain

import (
	"errors"
	"time"

	"github.com/oklog/ulid/v2"
)

var (
	ErrRequiredPaymentType = errors.New("param: 'PaymentType' is required, please try again")
	ErrInvalidPaymentType  = errors.New("param: 'PaymentType' must be 'cash', 'credit' or 'debit', please try again")
	ErrProductsRequired    = errors.New("at least one product is required, please try again")
)

type Sale struct {
	SaleID      string
	UserID      string
	Products    []SaleProduct
	Total       int
	PaymentType string
	Date        time.Time
}

type SaleProduct struct {
	ProductID string
	Name      string
	Unit      string
	Category  string
	Quantity  int
	Price     int
}

func NewSale(userID string, products []SaleProduct, paymentType string) (*Sale, error) {
	BRTime, _ := time.LoadLocation("America/Sao_Paulo")
	sale := &Sale{
		SaleID:      ulid.Make().String(),
		UserID:      userID,
		Products:    products,
		PaymentType: paymentType,
		Date:        time.Now().In(BRTime),
	}

	if err := sale.Validate(); err != nil {
		return nil, err
	}

	total := 0
	for _, product := range products {
		total += product.Price * product.Quantity
	}
	sale.Total = total

	return sale, nil
}

func (s *Sale) Validate() error {
	if s.PaymentType == "" {
		return ErrRequiredPaymentType
	}

	if s.PaymentType != "cash" && s.PaymentType != "credit" && s.PaymentType != "debit" {
		return ErrInvalidPaymentType
	}

	if len(s.Products) == 0 {
		return ErrProductsRequired
	}

	return nil
}
