package repository

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/jonattasmoraes/mimas/internal/product/domain"
)

var (
	ErrProductNotFound = errors.New("product not found, please enter a valid product and try again")
)

type ProductSqlx struct {
	writer *sqlx.DB
	reader *sqlx.DB
}

func NewSqlxRepository(writer, reader *sqlx.DB) domain.ProductRepository {
	return &ProductSqlx{writer: writer, reader: reader}
}

func (r *ProductSqlx) CreateProduct(product *domain.Product) error {
	query := `
	INSERT INTO product (id, name, price, unit_id, category_id, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.writer.Exec(query, product.ID, product.Name, product.Price, product.UnitID, product.CategoryID, product.CreatedAt, product.UpdatedAt)
	if err != nil {
		return err
	}

	query = `
	INSERT INTO stock (product_id, quantity)
	VALUES ($1, $2)
	`

	_, err = r.writer.Exec(query, product.ID, product.Stock.Quantity)
	if err != nil {
		return err
	}

	return nil
}

func (r *ProductSqlx) ListProducts(page int) ([]*domain.Product, error) {
	limit := 10
	offset := (page - 1) * limit

	query := `
	SELECT
		p.id, p.name, p.price, p.unit_id, p.category_id, p.created_at, p.updated_at,
		c.id AS category_id, c.name AS category_name,
		u.id AS unit_id, u.name AS unit_name,
		s.product_id, s.quantity
	FROM product p
	LEFT JOIN category c ON p.category_id = c.id
	LEFT JOIN unit u ON p.unit_id = u.id
	LEFT JOIN stock s ON p.id = s.product_id
	WHERE p.deleted_at IS NULL
	LIMIT $1 OFFSET $2
	`

	rows, err := r.reader.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*domain.Product

	for rows.Next() {
		var product domain.Product
		var category domain.Category
		var unit domain.Unit
		var stock domain.Stock

		err := rows.Scan(
			&product.ID, &product.Name, &product.Price, &product.UnitID, &product.CategoryID, &product.CreatedAt, &product.UpdatedAt,
			&category.ID, &category.Name,
			&unit.ID, &unit.Name,
			&stock.ProductID, &stock.Quantity,
		)
		if err != nil {
			return nil, err
		}

		product.Category = category
		product.Unit = unit
		product.Stock = stock

		products = append(products, &product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductSqlx) FindProductById(id string) (*domain.Product, error) {
	query := `
	SELECT
		p.id, p.name, p.price, p.unit_id, p.category_id, p.created_at, p.updated_at,
		c.id AS category_id, c.name AS category_name,
		u.id AS unit_id, u.name AS unit_name,
		s.product_id, s.quantity
	FROM product p
	LEFT JOIN category c ON p.category_id = c.id
	LEFT JOIN unit u ON p.unit_id = u.id
	LEFT JOIN stock s ON p.id = s.product_id
	WHERE p.id = $1
	AND p.deleted_at IS NULL
	`

	var product domain.Product
	var category domain.Category
	var unit domain.Unit
	var stock domain.Stock

	err := r.reader.QueryRow(query, id).Scan(
		&product.ID, &product.Name, &product.Price, &product.UnitID, &product.CategoryID, &product.CreatedAt, &product.UpdatedAt,
		&category.ID, &category.Name,
		&unit.ID, &unit.Name,
		&stock.ProductID, &stock.Quantity,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrProductNotFound
		}
		return nil, err
	}

	product.Category = category
	product.Unit = unit
	product.Stock = stock

	return &product, nil
}

func (r *ProductSqlx) DeleteProduct(id string) error {
	query := `
	UPDATE product
	SET deleted_at = NOW()
	WHERE id = $1
	`

	_, err := r.writer.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
