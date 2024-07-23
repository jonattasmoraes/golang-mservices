package usecase

import (
	"errors"

	"github.com/jonattasmoraes/mimas/internal/product/domain"
)

var ErrProductNotFound = errors.New("product not found")

type GetProductByIdUsecase struct {
	repo domain.ProductRepository
}

func NewGetProductByIdUsecase(repo domain.ProductRepository) *GetProductByIdUsecase {
	return &GetProductByIdUsecase{
		repo: repo,
	}
}

func (u *GetProductByIdUsecase) Execute(id string) (*ProductResponseDTO, error) {
	// Check if product exists
	product, err := u.repo.FindProductById(id)
	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, ErrProductNotFound
	}

	// Return product
	response := &ProductResponseDTO{
		ID:        product.ID,
		Name:      product.Name,
		Price:     product.Price,
		Unit:      product.Unit.Name,
		Category:  product.Category.Name,
		Stock:     product.Stock.Quantity,
		CreatedAt: product.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: product.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return response, nil
}
