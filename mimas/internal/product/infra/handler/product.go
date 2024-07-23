package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jonattasmoraes/mimas/internal/product/usecase"
	"github.com/jonattasmoraes/mimas/internal/utils"
)

type ProductHandler struct {
	createProduct *usecase.CreateProductUsecase
	listProducts  *usecase.ListProductsUsecase
	getById       *usecase.GetProductByIdUsecase
	deleteProduct *usecase.DeleteProductUsecase
}

func NewProductHandler(
	createProduct *usecase.CreateProductUsecase,
	listProducts *usecase.ListProductsUsecase,
	getById *usecase.GetProductByIdUsecase,
	deleteProduct *usecase.DeleteProductUsecase,
) *ProductHandler {
	return &ProductHandler{
		createProduct: createProduct,
		listProducts:  listProducts,
		getById:       getById,
		deleteProduct: deleteProduct,
	}
}

func (h *ProductHandler) CreateProduct(ctx *gin.Context) {
	var request usecase.ProductRequestDTO
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.createProduct.Execute(&request)
	if err != nil {
		if err == usecase.ErrCategoryNotFound || err == usecase.ErrUnitNotFound {
			utils.SendError(ctx, http.StatusBadRequest, err.Error())
			return
		}

		utils.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendSuccess(ctx, "product created", response, http.StatusCreated)
}

func (h *ProductHandler) ListProducts(ctx *gin.Context) {
	request := ctx.DefaultQuery("page", "1")

	page, err := strconv.Atoi(request)
	if err != nil {
		utils.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.listProducts.Execute(page)
	if err != nil {
		if err == usecase.ErrInvalidPageNumber {
			utils.SendError(ctx, http.StatusBadRequest, err.Error())
			return
		}
		if err == usecase.ErrProductsNotFound {
			utils.SendError(ctx, http.StatusNotFound, err.Error())
			return
		}
		utils.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendSuccess(ctx, "list products", response, http.StatusOK)
}

func (h *ProductHandler) FindProductById(ctx *gin.Context) {
	id := ctx.Param("id")

	response, err := h.getById.Execute(id)
	if err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendSuccess(ctx, "find product by id", response, http.StatusOK)
}

func (h *ProductHandler) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	response, err := h.deleteProduct.Execute(ctx, id)
	if err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendSuccess(ctx, "product deleted", response, http.StatusOK)
}
