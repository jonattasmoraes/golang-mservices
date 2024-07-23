package repository

import (
	"database/sql"
	"errors"
	"log"
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
		if err == sql.ErrNoRows {
			return nil, ErrProductNotFound
		}
		return nil, err
	}
	defer rows.Close()

	var (
		userID      string
		total       int
		paymentType string
		saleDate    time.Time
		products    []domain.ProductReport
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
		user_id, sale_id, product_id, name, unit, category, quantity, price, total, payment_type, sale_date
	FROM sale_reports
	WHERE user_id = $1
	`

	rows, err := r.reader.Queryx(query, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
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

	var sales []*domain.SaleReport
	for _, sale := range salesMap {
		sales = append(sales, sale)
	}

	return sales, nil
}

func (r *ReportSqlx) GetSalesByProductId(productID string) (*domain.SaleReport, error) {
	query := `
	SELECT
		user_id, sale_id, product_id, name, unit, category, quantity, price, total, payment_type, sale_date
	FROM sale_reports
	WHERE product_id = $1
	`

	rows, err := r.reader.Queryx(query, productID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrProductNotFound
		}
		return nil, err
	}
	defer rows.Close()

	var (
		userID      string
		saleID      string
		total       int
		paymentType string
		saleDate    time.Time
		products    []domain.ProductReport
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

		// Log para verificar os valores escaneados
		log.Printf("Scanned values: userID=%s, saleID=%s, productID=%s, name=%s, unit=%s, category=%s, quantity=%d, price=%d, total=%d, paymentType=%s, saleDate=%v",
			userID, saleID, productID, name, unit, category, quantity, price, total, paymentType, saleDate)

		product := domain.ProductReport{
			ProductID: productID,
			Name:      name,
			Unit:      unit,
			Category:  category,
			Quantity:  quantity,
			Price:     price,
		}
		products = append(products, product)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	report := &domain.SaleReport{
		UserID:      userID,
		SaleID:      saleID,
		Total:       total,
		PaymentType: paymentType,
		SaleDate:    saleDate,
		Products:    products,
	}

	// Log para verificar o relatório final
	log.Printf("Final report: %+v", report)

	return report, nil
}
