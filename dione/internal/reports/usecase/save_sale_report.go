package usecase

import (
	"encoding/json"
	"log"
	"time"

	"github.com/IBM/sarama"
	"github.com/jonattasmoraes/dione/internal/reports/domain"
)

type SaveSaleReportRequestDTO struct {
	SaleID      string                        `json:"SaleID"`
	UserID      string                        `json:"UserID"`
	Total       int                           `json:"total"`
	PaymentType string                        `json:"PaymentType"`
	Products    []SaveProductReportRequestDTO `json:"products"`
	SaleDate    time.Time                     `json:"date"`
}

type SaveProductReportRequestDTO struct {
	ProductID string  `json:"ProductID"`
	Name      string  `json:"name"`
	Unit      string  `json:"unit"`
	Category  string  `json:"category"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type SaveSaleReportUsecase struct {
	repo              domain.SaleReportRepository
	partitionConsumer sarama.PartitionConsumer
}

func NewSaveSaleReportUsecase(repo domain.SaleReportRepository, partition sarama.PartitionConsumer) *SaveSaleReportUsecase {
	return &SaveSaleReportUsecase{
		repo:              repo,
		partitionConsumer: partition,
	}
}

func (u *SaveSaleReportUsecase) Execute() {
	for {
		select {
		case msg := <-u.partitionConsumer.Messages():
			var dto SaveSaleReportRequestDTO
			err := json.Unmarshal(msg.Value, &dto)
			if err != nil {
				log.Printf("Failed to unmarshal message: %v", err)
				continue
			}

			report := domain.SaleReport{
				SaleID:      dto.SaleID,
				UserID:      dto.UserID,
				Total:       dto.Total,
				PaymentType: dto.PaymentType,
				Products:    []domain.ProductReport{},
				SaleDate:    dto.SaleDate,
			}

			for _, productDTO := range dto.Products {
				product := domain.ProductReport{
					ProductID: productDTO.ProductID,
					Name:      productDTO.Name,
					Unit:      productDTO.Unit,
					Category:  productDTO.Category,
					Quantity:  productDTO.Quantity,
					Price:     int(productDTO.Price),
				}
				report.Products = append(report.Products, product)
			}

			// Save sale report
			err = u.repo.SaveSale(&report)
			if err != nil {
				log.Printf("Failed to save sale report: %v", err)
			} else {
				log.Printf("Successfully saved sale report: %v", report.SaleID)
			}
		case err := <-u.partitionConsumer.Errors():
			log.Printf("Error consuming message: %v", err)
		}
	}
}
