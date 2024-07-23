package domain

type SaleReportRepository interface {
	SaveSale(saleReport *SaleReport) error
	GetBySaleID(saleID string) (*SaleReport, error)
	GetSalesByUserId(userID string) ([]*SaleReport, error)
	GetSalesByProductId(productID string) (*SaleReport, error)
}
