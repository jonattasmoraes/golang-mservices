package usecase

import (
	"errors"

	"github.com/jonattasmoraes/mimas/internal/product/domain"
)

var (
	ErrInvalidPageNumber = errors.New("invalid page number, enter a number greater than 0")
	ErrProductsNotFound  = errors.New("products not found")
)

type ListProductsUsecase struct {
	repo domain.ProductRepository
}

func NewListProductsUsecase(repo domain.ProductRepository) *ListProductsUsecase {
	return &ListProductsUsecase{
		repo: repo,
	}
}

func (u *ListProductsUsecase) Execute(page int) ([]ProductResponseDTO, error) {
	// Check if page number is valid
	if page < 1 {
		return nil, ErrInvalidPageNumber
	}

	// List products
	products, err := u.repo.ListProducts(page)
	if err != nil {
		return nil, err
	}

	// Check if products exist
	if len(products) == 0 {
		return nil, ErrProductsNotFound
	}

	// Return products
	response := make([]ProductResponseDTO, len(products))
	for i, product := range products {
		response[i] = ProductResponseDTO{
			ID:        product.ID,
			Name:      product.Name,
			Price:     product.Price,
			Unit:      product.Unit.Name,
			Category:  product.Category.Name,
			Stock:     product.Stock.Quantity,
			CreatedAt: product.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: product.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	return response, nil
}
