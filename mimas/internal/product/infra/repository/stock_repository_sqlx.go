package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/jonattasmoraes/mimas/internal/product/domain"
)

type StockSqlx struct {
	writer *sqlx.DB
	reader *sqlx.DB
}

func NewStockSqlxRepository(writer, reader *sqlx.DB) domain.StockRepository {
	return &StockSqlx{writer: writer, reader: reader}
}

func (s *StockSqlx) IncreaseStock(stock *domain.Stock) (*domain.Stock, error) {
	query := `
	UPDATE stock
	SET quantity = quantity + $1
	WHERE product_id = $2
	RETURNING quantity
	`

	var result domain.Stock
	err := s.writer.QueryRow(query, stock.Quantity, stock.ProductID).Scan(&result.Quantity)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *StockSqlx) DecreaseStock(stock *domain.Stock) (*domain.Stock, error) {
	query := `
	UPDATE stock
	SET quantity = quantity - $1
	WHERE product_id = $2
	RETURNING quantity
	`

	var result domain.Stock
	err := s.writer.QueryRow(query, stock.Quantity, stock.ProductID).Scan(&result.Quantity)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
