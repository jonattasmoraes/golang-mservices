package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jonattasmoraes/dione/internal/reports/infra/handler"
)

func StartRoutes(
	reportsHandler *handler.ReportsHandler,
) {
	server := gin.Default()

	reportPath := server.Group("/api/v1/reports")

	{
		reportPath.GET("/:id", reportsHandler.GetSalesById)
		reportPath.GET("/:id/purchases", reportsHandler.GetSalesByUserId)
	}

	server.Run(":8084")
}
