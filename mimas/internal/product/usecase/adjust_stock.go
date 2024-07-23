package usecase

import (
	"errors"

	"github.com/jonattasmoraes/mimas/internal/product/domain"
)

var (
	ErrInvalidOperationType = errors.New("operation must be 'increase' or 'decrease', please try again")
)

type AdjustStockRequestDTO struct {
	ProductID string `json:"product_id"`
	Operation string `json:"operation"`
	Quantity  int    `json:"quantity"`
}

type AdjustStockResponseDTO struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type AdjustStockUsecase struct {
	stockRepo domain.StockRepository
	prodRepo  domain.ProductRepository
}

func NewAdjustStockUsecase(stockRepo domain.StockRepository, prodRepo domain.ProductRepository) *AdjustStockUsecase {
	return &AdjustStockUsecase{
		stockRepo: stockRepo,
		prodRepo:  prodRepo,
	}
}

func (u *AdjustStockUsecase) Execute(stock *AdjustStockRequestDTO) (*AdjustStockResponseDTO, error) {
	// Check if product exists
	_, err := u.prodRepo.FindProductById(stock.ProductID)
	if err != nil {
		return nil, err
	}

	reqStock := &domain.Stock{
		ProductID: stock.ProductID,
		Quantity:  stock.Quantity,
	}

	// Check Operation type
	if stock.Operation != "increase" && stock.Operation != "decrease" {
		return nil, ErrInvalidOperationType
	}

	var adjustedStock *domain.Stock
	// Adjust Stock
	if stock.Operation == "increase" {
		adjustedStock, err = u.stockRepo.IncreaseStock(reqStock)
		if err != nil {
			return nil, err
		}
	}

	if stock.Operation == "decrease" {
		adjustedStock, err = u.stockRepo.DecreaseStock(reqStock)
		if err != nil {
			return nil, err
		}
	}

	// Return Product ID and Quantity
	response := &AdjustStockResponseDTO{
		ProductID: stock.ProductID,
		Quantity:  adjustedStock.Quantity,
	}

	return response, nil
}
