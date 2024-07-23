package usecase

import "github.com/jonattasmoraes/mimas/internal/product/domain"

type DeleteCategoryUseCase struct {
	CategoryRepository domain.CategoryRepository
}

func NewDeleteCategoryUseCase(repo domain.CategoryRepository) *DeleteCategoryUseCase {
	return &DeleteCategoryUseCase{
		CategoryRepository: repo,
	}
}

func (u *DeleteCategoryUseCase) Execute(id string) error {
	// Check if category exists
	_, err := u.CategoryRepository.FindCategoryById(id)
	if err != nil {
		return err
	}

	// Delete category
	err = u.CategoryRepository.DeleteCategory(id)
	if err != nil {
		return err
	}

	return nil
}
