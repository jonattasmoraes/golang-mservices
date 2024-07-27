package router

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/jonattasmoraes/dione/internal/reports/infra/gen"
	grpcService "github.com/jonattasmoraes/dione/internal/reports/infra/grpc"
	"github.com/jonattasmoraes/dione/internal/reports/usecase"
)

func StartGrpcServer(getSalesById *usecase.GetSalesById) {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterReportsServiceServer(grpcServer, grpcService.NewSalesReportsGrpcServer(getSalesById))

	reflection.Register(grpcServer)

	log.Printf("Listening and serving GRPC on :50053")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
