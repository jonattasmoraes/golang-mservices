package usecase

import "github.com/jonattasmoraes/mimas/internal/product/domain"

type ListCategoriesUsecase struct {
	repo domain.CategoryRepository
}

func NewListCategoriesUsecase(repo domain.CategoryRepository) *ListCategoriesUsecase {
	return &ListCategoriesUsecase{
		repo: repo,
	}
}

func (u *ListCategoriesUsecase) Execute() ([]CategoryResponseDTO, error) {
	// List categories
	categories, err := u.repo.ListCategories()
	if err != nil {
		return nil, err
	}

	// Return categories
	response := make([]CategoryResponseDTO, len(categories))
	for i, category := range categories {
		response[i] = CategoryResponseDTO{
			Name: category.Name,
		}
	}

	return response, nil
}
