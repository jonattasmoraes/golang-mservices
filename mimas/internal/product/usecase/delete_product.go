package usecase

import (
	"context"
	"errors"
	"log"

	"github.com/jonattasmoraes/mimas/internal/product/domain"
	reportpb "github.com/jonattasmoraes/mimas/internal/product/infra/gen/reports"
)

var (
	ErrProductHasNoReports = errors.New("product has sales reports and cannot be deleted")
)

type DeleteProductUsecase struct {
	repo         domain.ProductRepository
	reportClient reportpb.ReportsServiceClient
}

func NewDeleteProductUsecase(
	repo domain.ProductRepository,
	reportClient reportpb.ReportsServiceClient,
) *DeleteProductUsecase {
	return &DeleteProductUsecase{
		repo:         repo,
		reportClient: reportClient,
	}
}

func (u *DeleteProductUsecase) Execute(ctx context.Context, id string) (*ProductResponseDTO, error) {
	// Check if product exists
	product, err := u.repo.FindProductById(id)
	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, ErrProductNotFound
	}

	// Check if product has sales reports
	prod, err := u.reportClient.GetSalesByProductId(ctx, &reportpb.GetSalesByProductIdRequest{ProductId: product.ID})

	if err != nil {
		return nil, err
	}

	if prod.Products != nil {
		return nil, ErrProductHasNoReports
	}

	// Delete product
	err = u.repo.DeleteProduct(id)
	if err != nil {
		// Log error for debugging purposes
		log.Printf("Error deleting product: %v", err)
		return nil, err
	}

	// Return product response
	response := &ProductResponseDTO{
		ID:        product.ID,
		Name:      product.Name,
		Price:     product.Price,
		Unit:      product.Unit.Name,
		Category:  product.Category.Name,
		Stock:     product.Stock.Quantity,
		CreatedAt: product.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: product.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return response, nil
}
