package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonattasmoraes/mimas/internal/product/usecase"
	"github.com/jonattasmoraes/mimas/internal/utils"
)

type StockHandler struct {
	createStock *usecase.AddStockUsecase
	ajustStock  *usecase.AdjustStockUsecase
}

func NewStockHandler(createStock *usecase.AddStockUsecase, ajustStock *usecase.AdjustStockUsecase) *StockHandler {
	return &StockHandler{
		createStock: createStock,
		ajustStock:  ajustStock,
	}
}

func (h *StockHandler) AddStock(c *gin.Context) {
	var input usecase.AddStockRequestDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.createStock.Execute(&input)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendSuccess(c, "add stock", response, http.StatusCreated)
}

func (h *StockHandler) StockAjustment(c *gin.Context) {
	var input usecase.AdjustStockRequestDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.ajustStock.Execute(&input)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendSuccess(c, "stock adjustment", response, http.StatusCreated)
}
