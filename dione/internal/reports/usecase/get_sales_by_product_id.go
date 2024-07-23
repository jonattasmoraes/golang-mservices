package usecase

import "github.com/jonattasmoraes/dione/internal/reports/domain"

type GetSalesByProductId struct {
	repo domain.SaleReportRepository
}

func NewGetSalesByProductId(repo domain.SaleReportRepository) *GetSalesByProductId {
	return &GetSalesByProductId{
		repo: repo,
	}
}

func (u *GetSalesByProductId) Execute(productID string) (*SaveSaleReportRequestDTO, error) {
	reports, err := u.repo.GetSalesByProductId(productID)

	if err != nil {
		return nil, err
	}

	reponse := &SaveSaleReportRequestDTO{
		SaleID:      reports.SaleID,
		UserID:      reports.UserID,
		Total:       reports.Total,
		PaymentType: reports.PaymentType,
		SaleDate:    reports.SaleDate,
		Products:    []SaveProductReportRequestDTO{},
	}
	for _, product := range reports.Products {
		reponse.Products = append(reponse.Products, SaveProductReportRequestDTO{
			ProductID: product.ProductID,
			Name:      product.Name,
			Unit:      product.Unit,
			Category:  product.Category,
			Quantity:  product.Quantity,
			Price:     float64(product.Price),
		})
	}

	return reponse, nil
}
