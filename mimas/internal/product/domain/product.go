package domain

import (
	"errors"
	"time"

	"github.com/oklog/ulid/v2"
)

var (
	ErrAllParamsRequired = errors.New("all params are required, please try again")
	ErrNameIsRequired    = errors.New("param: 'Name' is required, please try again")
	ErrPriceIsRequired   = errors.New("param: 'Price' is required, please try again")
	ErrUnityIsRequired   = errors.New("param: 'UnitID' is required, please try again")
	ErrGroupIsRequired   = errors.New("param: 'CategoryID' is required, please try again")
)

type Product struct {
	ID         string
	Name       string
	Price      int
	UnitID     int
	CategoryID int
	Stock      Stock
	Category   Category
	Unit       Unit
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

func NewProduct(name string, price, unitID, categorId int) (*Product, error) {
	product := &Product{
		ID:         ulid.Make().String(),
		Name:       name,
		Price:      price,
		UnitID:     unitID,
		CategoryID: categorId,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := product.Validate(); err != nil {
		return nil, err
	}

	initialStock := 0
	stock := &Stock{
		Quantity: initialStock,
	}

	product.Stock = *stock

	return product, nil
}

func (product *Product) Validate() error {

	if product.Name == "" && product.Price == 0 && product.UnitID == 0 && product.CategoryID == 0 {
		return ErrAllParamsRequired
	}

	if product.Name == "" {
		return ErrNameIsRequired
	}

	if product.Price < 0 {
		return ErrPriceIsRequired
	}

	if product.UnitID == 0 {
		return ErrUnityIsRequired
	}

	if product.CategoryID == 0 {
		return ErrGroupIsRequired
	}

	return nil
}
