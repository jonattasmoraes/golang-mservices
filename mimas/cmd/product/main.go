package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/jonattasmoraes/mimas/internal/config"
	productpb "github.com/jonattasmoraes/mimas/internal/product/infra/gen/reports"
	"github.com/jonattasmoraes/mimas/internal/product/infra/handler"
	"github.com/jonattasmoraes/mimas/internal/product/infra/repository"
	"github.com/jonattasmoraes/mimas/internal/product/infra/router"
	"github.com/jonattasmoraes/mimas/internal/product/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	dsn := os.Getenv("DSN")
	dione := os.Getenv("GRPC_DIONE")

	writer, err := config.GetWriterSqlx(dsn)
	if err != nil {
		panic(err)
	}

	reader, err := config.GetReaderSqlx(dsn)
	if err != nil {
		panic(err)
	}

	defer writer.Close()
	defer reader.Close()

	config.StartMigrations(writer)

	reportConn := connectGRPC(dione)
	defer reportConn.Close()
	reportClient := productpb.NewReportsServiceClient(reportConn)

	repo := repository.NewSqlxRepository(writer, reader)
	catRepo := repository.NewSqlxCategoryRepository(writer, reader)
	unitRepo := repository.NewUnitSqlxRepository(writer, reader)
	stockRepo := repository.NewStockSqlxRepository(writer, reader)

	createProduct := usecase.NewCreateProductUsecase(repo, catRepo, unitRepo)
	listProducts := usecase.NewListProductsUsecase(repo)
	getProductById := usecase.NewGetProductByIdUsecase(repo)
	deleteProduct := usecase.NewDeleteProductUsecase(repo, reportClient)

	createUnit := usecase.NewCreateUnitUseCase(unitRepo)
	listUnits := usecase.NewListUnitsUsecase(unitRepo)
	deleteUnit := usecase.NewDeleteUnitUsecase(unitRepo, repo)

	createCategory := usecase.NewCreateCategoryUseCase(catRepo)
	listCategories := usecase.NewListCategoriesUsecase(catRepo)
	deleteCategory := usecase.NewDeleteCategoryUseCase(catRepo)

	addStock := usecase.NewAddStockUsecase(stockRepo, repo)
	stockAdjustment := usecase.NewAdjustStockUsecase(stockRepo, repo)
	saleProduct := usecase.NewSaleProductUsecase(repo, stockRepo)

	productsHandler := handler.NewProductHandler(
		createProduct,
		listProducts,
		getProductById,
		deleteProduct,
	)

	unitHandler := handler.NewUnitHandler(
		createUnit,
		deleteUnit,
		listUnits,
	)

	categoryHandler := handler.NewCategoryHandler(
		createCategory,
		deleteCategory,
		listCategories,
	)

	stockHandler := handler.NewStockHandler(
		addStock,
		stockAdjustment,
	)

	go func() {
		router.StartServer(productsHandler, unitHandler, categoryHandler, stockHandler)
	}()

	go func() {
		router.StartGrpcServer(getProductById, saleProduct)
	}()

	select {}
}

// connectGRPC connects to a gRPC server at the given address

func connectGRPC(address string) *grpc.ClientConn {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server at %s: %v", address, err)
	}
	return conn
}
