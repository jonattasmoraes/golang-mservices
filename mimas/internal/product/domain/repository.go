package domain

type ProductRepository interface {
	CreateProduct(product *Product) error
	FindProductById(id string) (*Product, error)
	ListProducts(page int) ([]*Product, error)
	DeleteProduct(id string) error
}

type CategoryRepository interface {
	CreateCategory(group *Category) error
	FindCategoryById(id string) (*Category, error)
	ListCategories() ([]*Category, error)
	DeleteCategory(id string) error
}

type UnitRepository interface {
	CreateUnit(unity *Unit) error
	FindUnitById(id string) (*Unit, error)
	ListUnits() ([]*Unit, error)
	DeleteUnit(id string) error
}

type StockRepository interface {
	IncreaseStock(stock *Stock) (*Stock, error)
	DecreaseStock(stock *Stock) (*Stock, error)
}
