package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/IBM/sarama"
	"github.com/jonattasmoraes/telesto/internal/sales/domain"
	productpb "github.com/jonattasmoraes/telesto/internal/sales/infra/gen/product"
	userpb "github.com/jonattasmoraes/telesto/internal/sales/infra/gen/user"
)

type SaleRequestDTO struct {
	UserID      string              `json:"user_id"`
	PaymentType string              `json:"payment_type"`
	Products    []ProductRequestDTO `json:"products"`
}

type ProductRequestDTO struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type ProductResponseDTO struct {
	ProductID string
	Name      string
	Unit      string
	Category  string
	Quantity  int
	Price     int
}

type SaleResponseDTO struct {
	UserID      string
	SaleID      string
	PaymentType string
	Date        string
	Products    []ProductResponseDTO
	Total       int
}

type SaleProductUsecase struct {
	salesRepo     domain.SalesRepository
	userClient    userpb.UserServiceClient
	productClient productpb.ProductServiceClient
	kafkaProducer sarama.SyncProducer
}

func NewSaleProductUsecase(
	salesRepo domain.SalesRepository,
	userClient userpb.UserServiceClient,
	productClient productpb.ProductServiceClient,
	kafkaProducer sarama.SyncProducer) *SaleProductUsecase {
	return &SaleProductUsecase{
		salesRepo:     salesRepo,
		userClient:    userClient,
		productClient: productClient,
		kafkaProducer: kafkaProducer,
	}
}

func (u *SaleProductUsecase) Execute(ctx context.Context, userID string, reqs []ProductRequestDTO, paymentType string) (*SaleResponseDTO, error) {
	_, err := u.userClient.GetUserByID(ctx, &userpb.GetUserRequest{Id: userID})
	if err != nil {
		return nil, err
	}

	var products []domain.SaleProduct
	total := 0

	for _, req := range reqs {
		productResp, err := u.productClient.GetProductByID(ctx, &productpb.GetProductRequest{Id: req.ProductID})
		if err != nil {
			return nil, err
		}

		if req.Quantity > int(productResp.Stock) {
			return nil, errors.New("quantity exceeds stock for product " + req.ProductID)
		}

		product := domain.SaleProduct{
			ProductID: req.ProductID,
			Name:      productResp.Name,
			Unit:      productResp.Unit,
			Category:  productResp.Category,
			Quantity:  req.Quantity,
			Price:     int(productResp.Price),
		}
		products = append(products, product)

		total += product.Price * product.Quantity
	}

	sale, err := domain.NewSale(userID, products, paymentType)
	if err != nil {
		return nil, err
	}

	if sale == nil {
		return nil, errors.New("invalid sale data")
	}
	sale.Total = total

	_, err = u.productClient.SaleProducts(ctx, &productpb.SaleProductsRequest{
		Products: productsToPB(products),
	})
	if err != nil {
		return nil, err
	}

	err = u.salesRepo.Save(sale)
	if err != nil {
		return nil, err
	}

	err = u.sendSaleToKafka(sale)
	if err != nil {
		return nil, err
	}

	response := &SaleResponseDTO{
		UserID:      sale.UserID,
		SaleID:      sale.SaleID,
		PaymentType: sale.PaymentType,
		Date:        sale.Date.Format(time.DateTime),
		Products:    productsToDTO(products),
		Total:       sale.Total,
	}

	return response, nil
}

func (u *SaleProductUsecase) sendSaleToKafka(sale *domain.Sale) error {
	saleData, err := json.Marshal(sale)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: "sales",
		Value: sarama.StringEncoder(saleData),
	}

	_, _, err = u.kafkaProducer.SendMessage(msg)
	return err
}

func productsToDTO(products []domain.SaleProduct) []ProductResponseDTO {
	var dtos []ProductResponseDTO
	for _, product := range products {
		dtos = append(dtos, ProductResponseDTO{
			ProductID: product.ProductID,
			Name:      product.Name,
			Unit:      product.Unit,
			Category:  product.Category,
			Quantity:  product.Quantity,
			Price:     product.Price,
		})
	}
	return dtos
}

func productsToPB(products []domain.SaleProduct) []*productpb.SaleProductRequest {
	var pbProducts []*productpb.SaleProductRequest
	for _, product := range products {
		pbProducts = append(pbProducts, &productpb.SaleProductRequest{
			ProductId: product.ProductID,
			Quantity:  int32(product.Quantity),
		})
	}
	return pbProducts
}
