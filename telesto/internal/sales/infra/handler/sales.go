package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonattasmoraes/telesto/internal/sales/usecase"
)

type SaleHandler struct {
	saleProductUsecase *usecase.SaleProductUsecase
}

func NewSaleHandler(saleProductUsecase *usecase.SaleProductUsecase) *SaleHandler {
	return &SaleHandler{
		saleProductUsecase: saleProductUsecase,
	}
}

func (h *SaleHandler) MakeSale(ctx *gin.Context) {
	var request usecase.SaleRequestDTO
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	saleProducts := makeSale(request)

	response, err := h.saleProductUsecase.Execute(ctx, request.UserID, saleProducts, request.PaymentType)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": response, "message": "purchase made successfully"})
}

func makeSale(request usecase.SaleRequestDTO) []usecase.ProductRequestDTO {
	var saleProducts []usecase.ProductRequestDTO
	for _, product := range request.Products {
		saleProduct := usecase.ProductRequestDTO{
			ProductID: product.ProductID,
			Quantity:  product.Quantity,
		}
		saleProducts = append(saleProducts, saleProduct)
	}

	return saleProducts
}
