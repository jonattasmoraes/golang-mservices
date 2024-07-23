package usecase

import (
	"errors"

	"github.com/jonattasmoraes/mimas/internal/product/domain"
)

var ErrInvalidQuantity = errors.New("invalid quantity, enter a number greater than 0")

type AddStockRequestDTO struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type AddStockResponseDTO struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type AddStockUsecase struct {
	repo     domain.StockRepository
	prodRepo domain.ProductRepository
}

func NewAddStockUsecase(repo domain.StockRepository, prodRepo domain.ProductRepository) *AddStockUsecase {
	return &AddStockUsecase{
		repo:     repo,
		prodRepo: prodRepo,
	}
}

func (u *AddStockUsecase) Execute(stock *AddStockRequestDTO) (*AddStockResponseDTO, error) {
	// Check if product exists
	product, err := u.prodRepo.FindProductById(stock.ProductID)
	if err != nil {
		return nil, err
	}

	// Check if quantity is valid
	if stock.Quantity < 1 {
		return nil, ErrInvalidQuantity
	}

	// Create domain stock entity
	domainStock := &domain.Stock{
		ProductID: product.ID,
		Quantity:  stock.Quantity,
	}

	// Update stock
	qtd, err := u.repo.IncreaseStock(domainStock)
	if err != nil {
		return nil, err
	}

	// Return stock response
	response := &AddStockResponseDTO{
		ProductID: stock.ProductID,
		Quantity:  qtd.Quantity,
	}

	return response, nil
}
