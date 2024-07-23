package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jonattasmoraes/mimas/internal/product/infra/handler"
)

func StartServer(
	productHandler *handler.ProductHandler,
	unitHandler *handler.UnitHandler,
	categoryHandler *handler.CategoryHandler,
	stockHandler *handler.StockHandler,
) {
	router := gin.Default()

	product := router.Group("/api")
	{
		product.POST("/product", productHandler.CreateProduct)
		product.GET("/products", productHandler.ListProducts)
		product.GET("/product/:id", productHandler.FindProductById)
		product.DELETE("/product/:id", productHandler.DeleteProduct)
	}
	{
		product.POST("/unit", unitHandler.CreateUnit)
		product.GET("/units", unitHandler.ListUnits)
		product.DELETE("/unit/:id", unitHandler.DeleteUnit)
	}
	{
		product.POST("/category", categoryHandler.CreateCategory)
		product.GET("/categories", categoryHandler.ListCategories)
		product.DELETE("/category/:id", categoryHandler.DeleteCategory)
	}
	{
		product.POST("/stock", stockHandler.AddStock)
		product.PATCH("/stock", stockHandler.StockAjustment)
	}

	router.Run(":8081")
}
