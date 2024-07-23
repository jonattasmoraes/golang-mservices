package domain

import "errors"

var (
	ErrCategoryNameEmpty    = errors.New("category name cannot be empty")
	ErrCategoryNameTooShort = errors.New("category name must be at least 3 characters long")
)

type Category struct {
	ID   int
	Name string
}

func NewCategory(name string) (*Category, error) {
	category := &Category{
		Name: name,
	}

	return category, nil
}

func (category *Category) Validate() error {
	if category.Name == "" {
		return ErrCategoryNameEmpty
	}

	if len(category.Name) < 3 {
		return ErrCategoryNameTooShort
	}

	return nil
}
