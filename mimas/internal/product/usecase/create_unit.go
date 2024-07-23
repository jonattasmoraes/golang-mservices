package usecase

import (
	"github.com/jonattasmoraes/mimas/internal/product/domain"
)

type UnitRequestDTO struct {
	Name string `json:"name"`
}

type UnitResponseDTO struct {
	Name string `json:"name"`
}

type CreateUnitUseCase struct {
	UnitRepository domain.UnitRepository
}

func NewCreateUnitUseCase(repo domain.UnitRepository) *CreateUnitUseCase {
	return &CreateUnitUseCase{
		UnitRepository: repo,
	}
}

func (u *CreateUnitUseCase) Execute(input *UnitRequestDTO) (*UnitResponseDTO, error) {
	// Create unit
	unity, err := domain.NewUnit(input.Name)
	if err != nil {
		return nil, err
	}

	err = u.UnitRepository.CreateUnit(unity)
	if err != nil {
		return nil, err
	}

	// Return unit
	response := &UnitResponseDTO{
		Name: unity.Name,
	}

	return response, nil
}
