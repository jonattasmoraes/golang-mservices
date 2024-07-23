package main

import (
	"log"
	"os"

	"github.com/IBM/sarama"
	"github.com/joho/godotenv"
	"github.com/jonattasmoraes/dione/internal/config"
	handl "github.com/jonattasmoraes/dione/internal/reports/infra/handler"
	"github.com/jonattasmoraes/dione/internal/reports/infra/repository"
	"github.com/jonattasmoraes/dione/internal/reports/infra/router"
	"github.com/jonattasmoraes/dione/internal/reports/usecase"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get environment variables
	dsn := os.Getenv("DSN")
	address := os.Getenv("KAFKA_ADDRESS")

	// Initialize database connections
	writer, err := config.GetWriterSqlx(dsn)
	if err != nil {
		log.Fatalf("Failed to get writer: %v", err)
	}
	defer writer.Close()

	reader, err := config.GetReaderSqlx(dsn)
	if err != nil {
		log.Fatalf("Failed to get reader: %v", err)
	}
	defer reader.Close()

	// Run migrations
	config.StartMigrations(writer)

	// Initialize kafka consumer
	consumer := initKafkaConsumer(address)
	defer consumer.Close()

	// Initialize kafka partition consumer
	partitionConsumer := kafkaConsumer(consumer, "sales", 0, sarama.OffsetOldest)
	defer partitionConsumer.Close()

	// Initialize repository and usecase
	repo := repository.NewSalesRepository(writer, reader)
	getSaleById := usecase.NewGetSalesById(repo)
	getSalesByUser := usecase.NewGetSalesByUserIDUsecase(repo)
	getSalesByProductId := usecase.NewGetSalesByProductId(repo)

	reportHandler := handl.NewGetSaleByIdHandler(
		getSaleById,
		getSalesByUser,
		getSalesByProductId,
	)

	reportMessages := usecase.NewSaveSaleReportUsecase(repo, partitionConsumer)

	// Start gRPC server
	go func() {
		router.StartGrpcServer(getSaleById, getSalesByProductId)
	}()

	// Start consumer
	go func() {
		reportMessages.Execute()
	}()

	// Start HTTP server
	go func() {
		router.StartRoutes(reportHandler)
	}()

	select {}
}

// Initialize kafka consumer function
func initKafkaConsumer(address string) sarama.Consumer {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	brokers := []string{address}
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Fatalf("Failed to start consumer: %v", err)
	}
	return consumer
}

// Initialize kafka partition consumer function
func kafkaConsumer(consumer sarama.Consumer, topic string, partition int32, offset int64) sarama.PartitionConsumer {
	partitionConsumer, err := consumer.ConsumePartition(topic, partition, offset)
	if err != nil {
		log.Fatalf("Failed to start consumer for partition: %v", err)
	}
	return partitionConsumer
}
