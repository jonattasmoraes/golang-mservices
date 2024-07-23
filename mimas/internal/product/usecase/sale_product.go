package usecase

import (
	"errors"

	"github.com/jonattasmoraes/mimas/internal/product/domain"
)

var (
	ErrStockInsufficient = errors.New("insufficient stock")
)

type SaleProductRequestDTO struct {
	ProductID string
	Quantity  int
}

type SaleProductResponseDTO struct {
	ProductID string
	Quantity  int
}

type SaleProductUsecase struct {
	prodRepo  domain.ProductRepository
	stockRepo domain.StockRepository
}

func NewSaleProductUsecase(prodRepo domain.ProductRepository, stockRepo domain.StockRepository) *SaleProductUsecase {
	return &SaleProductUsecase{
		prodRepo:  prodRepo,
		stockRepo: stockRepo,
	}
}

func (u *SaleProductUsecase) Execute(requests []SaleProductRequestDTO) ([]SaleProductResponseDTO, error) {
	var responses []SaleProductResponseDTO

	for _, req := range requests {
		// Check if product exists
		product, err := u.prodRepo.FindProductById(req.ProductID)
		if err != nil {
			return nil, err
		}

		// Check if quantity is valid
		if req.Quantity < 1 {
			return nil, ErrInvalidQuantity
		}

		// Check if stock is sufficient
		if product.Stock.Quantity < req.Quantity {
			return nil, ErrStockInsufficient
		}

		reqStock := &domain.Stock{
			ProductID: req.ProductID,
			Quantity:  req.Quantity,
		}

		// Sale product
		newQty, err := u.stockRepo.DecreaseStock(reqStock)
		if err != nil {
			return nil, err
		}

		response := SaleProductResponseDTO{
			ProductID: req.ProductID,
			Quantity:  newQty.Quantity,
		}

		responses = append(responses, response)
	}

	return responses, nil
}
