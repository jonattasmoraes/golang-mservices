package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/jonattasmoraes/telesto/internal/sales/domain"
)

type SalesSqlx struct {
	writer *sqlx.DB
	reader *sqlx.DB
}

func NewSalesRepository(writer, reader *sqlx.DB) domain.SalesRepository {
	return &SalesSqlx{writer: writer, reader: reader}
}

func (s *SalesSqlx) Save(sale *domain.Sale) error {
	insertSalesQuery := `
		INSERT INTO sales (id, user_id, total, payment_type, date)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := s.writer.Exec(insertSalesQuery, sale.SaleID, sale.UserID, sale.Total, sale.PaymentType, sale.Date)
	if err != nil {
		return err
	}

	insertSalesProductsQuery := `
		INSERT INTO sale_products (sale_id, product_id, quantity, price)
		VALUES ($1, $2, $3, $4)
	`
	for _, sp := range sale.Products {
		_, err := s.writer.Exec(insertSalesProductsQuery, sale.SaleID, sp.ProductID, sp.Quantity, sp.Price)
		if err != nil {
			return err
		}
	}

	return nil
}
