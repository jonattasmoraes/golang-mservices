package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonattasmoraes/mimas/internal/product/usecase"
	"github.com/jonattasmoraes/mimas/internal/utils"
)

type CategoryHandler struct {
	createCategory *usecase.CreateCategoryUseCase
	deleteCategory *usecase.DeleteCategoryUseCase
	listCategories *usecase.ListCategoriesUsecase
}

func NewCategoryHandler(
	createCategory *usecase.CreateCategoryUseCase,
	deleteCategory *usecase.DeleteCategoryUseCase,
	listCategories *usecase.ListCategoriesUsecase,
) *CategoryHandler {
	return &CategoryHandler{
		createCategory: createCategory,
		deleteCategory: deleteCategory,
		listCategories: listCategories,
	}
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var input usecase.CategoryRequestDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.createCategory.Execute(&input)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendSuccess(c, "create category", response, http.StatusCreated)
}

func (h *CategoryHandler) ListCategories(c *gin.Context) {
	response, err := h.listCategories.Execute()
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendSuccess(c, "list categories", response, http.StatusOK)
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	err := h.deleteCategory.Execute(id)
	if err != nil {
		if err == usecase.ErrUnitNotFound {
			utils.SendError(c, http.StatusNotFound, err.Error())
			return
		}
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendSuccess(c, "unit deleted successfully", nil, http.StatusOK)
}
