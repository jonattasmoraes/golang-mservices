package usecase

import (
	"github.com/jonattasmoraes/mimas/internal/product/domain"
)

type CategoryRequestDTO struct {
	Name string `json:"name"`
}

type CategoryResponseDTO struct {
	Name string `json:"name"`
}

type CreateCategoryUseCase struct {
	CategoryRepository domain.CategoryRepository
}

func NewCreateCategoryUseCase(repo domain.CategoryRepository) *CreateCategoryUseCase {
	return &CreateCategoryUseCase{
		CategoryRepository: repo,
	}
}

func (u *CreateCategoryUseCase) Execute(input *CategoryRequestDTO) (*CategoryResponseDTO, error) {
	// Create category
	category, err := domain.NewCategory(input.Name)
	if err != nil {
		return nil, err
	}

	err = u.CategoryRepository.CreateCategory(category)
	if err != nil {
		return nil, err
	}

	// Return category
	response := &CategoryResponseDTO{
		Name: category.Name,
	}

	return response, nil
}
