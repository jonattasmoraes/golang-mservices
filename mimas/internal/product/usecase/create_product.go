package usecase

import (
	"errors"
	"strconv"
	"time"

	"github.com/jonattasmoraes/mimas/internal/product/domain"
)

var (
	ErrCategoryNotFound = errors.New("category not found, please enter a valid category and try again")
	ErrUnitNotFound     = errors.New("unit not found, please enter a valid unit and try again")
)

type CreateProductUsecase struct {
	repo     domain.ProductRepository
	catRepo  domain.CategoryRepository
	unitRepo domain.UnitRepository
}

func NewCreateProductUsecase(repo domain.ProductRepository, catRepo domain.CategoryRepository, unitRepo domain.UnitRepository) *CreateProductUsecase {
	return &CreateProductUsecase{
		repo:     repo,
		catRepo:  catRepo,
		unitRepo: unitRepo,
	}
}

type ProductRequestDTO struct {
	Name       string    `json:"name"`
	Price      int       `json:"price"`
	UnitID     int       `json:"unit_id"`
	CategoryID int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ProductResponseDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	Unit      string `json:"unit"`
	Category  string `json:"category"`
	Stock     int    `json:"stock"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (u *CreateProductUsecase) Execute(input *ProductRequestDTO) (*ProductResponseDTO, error) {
	// Convert IDs from int to string
	unitIDStr := strconv.Itoa(input.UnitID)
	categoryIDStr := strconv.Itoa(input.CategoryID)

	// Check if unit exists
	unit, err := u.unitRepo.FindUnitById(unitIDStr)
	if err != nil {
		return nil, ErrUnitNotFound
	}

	// Check if category exists
	category, err := u.catRepo.FindCategoryById(categoryIDStr)
	if err != nil {
		return nil, ErrCategoryNotFound
	}

	// Create product
	product, err := domain.NewProduct(
		input.Name,
		input.Price,
		unit.ID,
		category.ID,
	)
	if err != nil {
		return nil, err
	}

	// Create product in repository
	err = u.repo.CreateProduct(product)
	if err != nil {
		return nil, err
	}

	// Return product response
	response := &ProductResponseDTO{
		ID:        product.ID,
		Name:      product.Name,
		Price:     product.Price,
		Unit:      unit.Name,
		Category:  category.Name,
		Stock:     product.Stock.Quantity,
		CreatedAt: product.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: product.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return response, nil
}
