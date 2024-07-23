package domain

type SalesRepository interface {
	Save(sale *Sale) error
}
