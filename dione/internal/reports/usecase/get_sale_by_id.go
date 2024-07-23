package usecase

import (
	"github.com/jonattasmoraes/dione/internal/reports/domain"
)

type GetSalesById struct {
	repo domain.SaleReportRepository
}

func NewGetSalesById(repo domain.SaleReportRepository) *GetSalesById {
	return &GetSalesById{
		repo: repo,
	}
}

func (u *GetSalesById) Execute(id string) (*domain.SaleReport, error) {
	report, err := u.repo.GetBySaleID(id)
	if err != nil {
		return nil, err
	}

	return report, nil
}
