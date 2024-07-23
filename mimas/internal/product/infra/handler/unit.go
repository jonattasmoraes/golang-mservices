package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonattasmoraes/mimas/internal/product/usecase"
	"github.com/jonattasmoraes/mimas/internal/utils"
)

type UnitHandler struct {
	createUnit *usecase.CreateUnitUseCase
	deleteUnit *usecase.DeleteUnitUsecase
	listUnits  *usecase.ListUnitsUsecase
}

func NewUnitHandler(
	createUnit *usecase.CreateUnitUseCase,
	deleteUnit *usecase.DeleteUnitUsecase,
	listUnits *usecase.ListUnitsUsecase,
) *UnitHandler {
	return &UnitHandler{
		createUnit: createUnit,
		deleteUnit: deleteUnit,
		listUnits:  listUnits,
	}
}

func (u *UnitHandler) CreateUnit(c *gin.Context) {
	var input usecase.UnitRequestDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	response, err := u.createUnit.Execute(&input)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendSuccess(c, "create unit", response, http.StatusCreated)
}

func (u *UnitHandler) ListUnits(c *gin.Context) {
	response, err := u.listUnits.Execute()
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendSuccess(c, "list units", response, http.StatusOK)
}

func (u *UnitHandler) DeleteUnit(c *gin.Context) {
	id := c.Param("id")

	err := u.deleteUnit.Execute(id)
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
