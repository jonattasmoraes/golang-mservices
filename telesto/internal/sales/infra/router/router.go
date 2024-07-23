package router

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jonattasmoraes/telesto/internal/sales/infra/handler"
)

func StartRouter(handler *handler.SaleHandler) {
	server := gin.Default()

	server.POST("/sales", handler.MakeSale)

	port := getServerPort()

	if err := server.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func getServerPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}
	return port
}
