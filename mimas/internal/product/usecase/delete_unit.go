package usecase

import (
	"errors"

	"github.com/jonattasmoraes/mimas/internal/product/domain"
)

var (
	ErrUnitIsUsed = errors.New("unit is used, can't be deleted")
)

type DeleteUnitUsecase struct {
	unitRepo domain.UnitRepository
}

func NewDeleteUnitUsecase(unitRepo domain.UnitRepository, productRepo domain.ProductRepository) *DeleteUnitUsecase {
	return &DeleteUnitUsecase{
		unitRepo: unitRepo,
	}
}

func (u *DeleteUnitUsecase) Execute(id string) error {
	// Verify if unit exists
	_, err := u.unitRepo.FindUnitById(id)
	if err != nil {
		return err
	}

	// Delete the unit
	err = u.unitRepo.DeleteUnit(id)
	if err != nil {
		return err
	}

	return nil
}
