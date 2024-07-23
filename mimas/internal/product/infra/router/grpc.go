package router

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/jonattasmoraes/mimas/internal/product/infra/gen"
	grpcService "github.com/jonattasmoraes/mimas/internal/product/infra/grpc"
	"github.com/jonattasmoraes/mimas/internal/product/usecase"
)

func StartGrpcServer(getProductById *usecase.GetProductByIdUsecase, saleProduct *usecase.SaleProductUsecase) {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterProductServiceServer(grpcServer, grpcService.NewProductGrpcServer(
		getProductById,
		saleProduct,
	))

	reflection.Register(grpcServer)

	log.Printf("Listening and serving GRPC on :50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
