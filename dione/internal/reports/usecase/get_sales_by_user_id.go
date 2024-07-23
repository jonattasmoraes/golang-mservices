package usecase

import (
	"github.com/jonattasmoraes/dione/internal/reports/domain"
)

type GetSalesByUserIDUsecase struct {
	repo domain.SaleReportRepository
}

func NewGetSalesByUserIDUsecase(repo domain.SaleReportRepository) *GetSalesByUserIDUsecase {
	return &GetSalesByUserIDUsecase{
		repo: repo,
	}
}

func (u *GetSalesByUserIDUsecase) Execute(userID string) ([]*domain.SaleReport, error) {
	report, err := u.repo.GetSalesByUserId(userID)
	if err != nil {
		return nil, err
	}

	return report, nil
}
