package repository

import (
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/jonattasmoraes/dione/internal/reports/domain"
)

var (
	ErrProductNotFound = errors.New("product not found, please enter a valid product and try again")
)

type ReportSqlx struct {
	writer *sqlx.DB
	reader *sqlx.DB
}

func NewSalesRepository(writer, reader *sqlx.DB) domain.SaleReportRepository {
	return &ReportSqlx{writer: writer, reader: reader}
}

func (r *ReportSqlx) SaveSale(saleReport *domain.SaleReport) error {
	query := `
	INSERT INTO sale_reports (user_id, sale_id, product_id, name, unit, category, quantity, price, total, payment_type, sale_date)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`

	for _, product := range saleReport.Products {
		_, err := r.writer.Exec(
			query,
			saleReport.UserID,
			saleReport.SaleID,
			product.ProductID,
			product.Name,
			product.Unit,
			product.Category,
			product.Quantity,
			product.Price,
			saleReport.Total,
			saleReport.PaymentType,
			saleReport.SaleDate,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *ReportSqlx) GetBySaleID(saleID string) (*domain.SaleReport, error) {
	query := `
	SELECT
		user_id, sale_id, product_id, name, unit, category, quantity, price, total, payment_type, sale_date
	FROM sale_reports
	WHERE sale_id = $1
	`

	rows, err := r.reader.Queryx(query, saleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		userID      string
		total       int
		paymentType string
		saleDate    time.Time
		productMap  = make(map[string]domain.ProductReport)
	)

	for rows.Next() {
		var (
			productID string
			name      string
			unit      string
			category  string
			quantity  int
			price     int
		)

		err := rows.Scan(
			&userID,
			&saleID,
			&productID,
			&name,
			&unit,
			&category,
			&quantity,
			&price,
			&total,
			&paymentType,
			&saleDate,
		)
		if err != nil {
			return nil, err
		}

		product := domain.ProductReport{
			ProductID: productID,
			Name:      name,
			Unit:      unit,
			Category:  category,
			Quantity:  quantity,
			Price:     price,
		}

		productMap[productID] = product
	}

	if len(productMap) == 0 {
		return nil, ErrProductNotFound
	}

	var products []domain.ProductReport
	for _, product := range productMap {
		products = append(products, product)
	}

	report := &domain.SaleReport{
		UserID:      userID,
		SaleID:      saleID,
		Total:       total,
		PaymentType: paymentType,
		SaleDate:    saleDate,
		Products:    products,
	}

	return report, nil
}


func (r *ReportSqlx) GetSalesByUserId(userID string) ([]*domain.SaleReport, error) {
	query := `
	SELECT
		sale_id, product_id, name, unit, category, quantity, price, total, payment_type, sale_date
	FROM sale_reports
	WHERE user_id = $1
	ORDER BY sale_date
	`

	rows, err := r.reader.Queryx(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	salesMap := make(map[string]*domain.SaleReport)
	for rows.Next() {
		var (
			saleID      string
			productID   string
			name        string
			unit        string
			category    string
			quantity    int
			price       int
			total       int
			paymentType string
			saleDate    time.Time
		)

		err := rows.Scan(
			&saleID,
			&productID,
			&name,
			&unit,
			&category,
			&quantity,
			&price,
			&total,
			&paymentType,
			&saleDate,
		)
		if err != nil {
			return nil, err
		}

		product := domain.ProductReport{
			ProductID: productID,
			Name:      name,
			Unit:      unit,
			Category:  category,
			Quantity:  quantity,
			Price:     price,
		}

		if sale, exists := salesMap[saleID]; exists {
			sale.Products = append(sale.Products, product)
		} else {
			salesMap[saleID] = &domain.SaleReport{
				UserID:      userID,
				SaleID:      saleID,
				Total:       total,
				PaymentType: paymentType,
				SaleDate:    saleDate,
				Products:    []domain.ProductReport{product},
			}
		}
	}

	if len(salesMap) == 0 {
		return nil, ErrProductNotFound
	}

	var sales []*domain.SaleReport
	for _, sale := range salesMap {
		sales = append(sales, sale)
	}

	return sales, nil
}
