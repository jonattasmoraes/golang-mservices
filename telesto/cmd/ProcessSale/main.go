package main

import (
	"log"
	"os"

	"github.com/IBM/sarama"
	"github.com/joho/godotenv"
	"github.com/jonattasmoraes/telesto/internal/config"
	productpb "github.com/jonattasmoraes/telesto/internal/sales/infra/gen/product"
	userpb "github.com/jonattasmoraes/telesto/internal/sales/infra/gen/user"
	handler "github.com/jonattasmoraes/telesto/internal/sales/infra/handler"
	"github.com/jonattasmoraes/telesto/internal/sales/infra/repository"
	"github.com/jonattasmoraes/telesto/internal/sales/infra/router"
	"github.com/jonattasmoraes/telesto/internal/sales/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get environment variables
	dsn := os.Getenv("DSN")
	titan := os.Getenv("GRPC_TITAN")
	mimas := os.Getenv("GRPC_MIMAS")
	kafkaAddress := os.Getenv("KAFKA_ADDRESS")

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

	// Run database migrations
	config.StartMigrations(writer)

	// Initialize repositories
	salesRepo := repository.NewSalesRepository(writer, reader)

	// Connect to gRPC services
	userConn := connectGRPC(titan)
	defer userConn.Close()
	userClient := userpb.NewUserServiceClient(userConn)

	productConn := connectGRPC(mimas)
	defer productConn.Close()
	productClient := productpb.NewProductServiceClient(productConn)

	// Configure and initialize Kafka producer
	producer := initKafkaProducer(kafkaAddress)
	defer producer.Close()

	// Initialize use cases and handlers
	saleProductUsecase := usecase.NewSaleProductUsecase(salesRepo, userClient, productClient, producer)
	saleHandler := handler.NewSaleHandler(saleProductUsecase)

	// Start the router
	router.StartRouter(saleHandler)
}

// connectGRPC connects to a gRPC server at the given address
func connectGRPC(address string) *grpc.ClientConn {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server at %s: %v", address, err)
	}
	return conn
}

// initKafkaProducer initializes a Kafka producer with the given address
func initKafkaProducer(address string) sarama.SyncProducer {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{address}, kafkaConfig)
	if err != nil {
		log.Fatalf("Failed to start Kafka producer: %v", err)
	}
	return producer
}
