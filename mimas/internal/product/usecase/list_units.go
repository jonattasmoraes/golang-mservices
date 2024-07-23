package usecase

import "github.com/jonattasmoraes/mimas/internal/product/domain"

type ListUnitsUsecase struct {
	repo domain.UnitRepository
}

func NewListUnitsUsecase(repo domain.UnitRepository) *ListUnitsUsecase {
	return &ListUnitsUsecase{
		repo: repo,
	}
}

func (u *ListUnitsUsecase) Execute() ([]UnitResponseDTO, error) {
	unities, err := u.repo.ListUnits()
	if err != nil {
		return nil, err
	}

	response := make([]UnitResponseDTO, len(unities))
	for i, unit := range unities {
		response[i] = UnitResponseDTO{
			Name: unit.Name,
		}
	}

	return response, nil
}
