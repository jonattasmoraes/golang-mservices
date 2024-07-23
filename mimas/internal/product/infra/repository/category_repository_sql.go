package repository

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/jonattasmoraes/mimas/internal/product/domain"
)

var (
	ErrCategoryNotFound = errors.New("category not found, please enter a valid category and try again")
)

type categorySqlx struct {
	writer *sqlx.DB
	reader *sqlx.DB
}

func NewSqlxCategoryRepository(writer, reader *sqlx.DB) domain.CategoryRepository {
	return &categorySqlx{writer: writer, reader: reader}
}

func (c *categorySqlx) CreateCategory(category *domain.Category) error {
	query := `
	INSERT INTO category (name)
	VALUES ($1)
	`

	_, err := c.writer.Exec(query, category.Name)
	if err != nil {
		return err
	}
	return nil
}

func (c *categorySqlx) ListCategories() ([]*domain.Category, error) {
	query := `
	SELECT id, name
	FROM category
	`

	rows, err := c.reader.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*domain.Category
	for rows.Next() {
		var category domain.Category
		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}

	return categories, nil
}

func (c *categorySqlx) FindCategoryById(id string) (*domain.Category, error) {
	query := `
	SELECT id, name
	FROM category
	WHERE id = $1
	`

	var category domain.Category
	err := c.reader.QueryRow(query, id).Scan(&category.ID, &category.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrCategoryNotFound
		}
		return nil, err
	}

	return &category, nil
}

func (c *categorySqlx) DeleteCategory(id string) error {
	query := `
	DELETE FROM category
	WHERE id = $1
	`

	_, err := c.writer.Exec(query, id)
	if err != nil {
		if strings.Contains(err.Error(), "violates foreign key constraint") {
			return ErrAlreadyUsed
		}
		return err
	}

	return nil
}
