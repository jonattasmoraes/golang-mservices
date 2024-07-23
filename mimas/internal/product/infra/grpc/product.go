package grpc

import (
	"context"

	pb "github.com/jonattasmoraes/mimas/internal/product/infra/gen"
	"github.com/jonattasmoraes/mimas/internal/product/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type productGrpcServer struct {
	pb.UnimplementedProductServiceServer
	productService *usecase.GetProductByIdUsecase
	saleProduct    *usecase.SaleProductUsecase
}

func NewProductGrpcServer(productService *usecase.GetProductByIdUsecase, saleProduct *usecase.SaleProductUsecase) *productGrpcServer {
	return &productGrpcServer{
		productService: productService,
		saleProduct:    saleProduct,
	}
}

func (s *productGrpcServer) GetProductByID(ctx context.Context, in *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	product, err := s.productService.Execute(in.Id)
	if err != nil {
		return nil, err
	}

	response := &pb.GetProductResponse{
		Id:       product.ID,
		Name:     product.Name,
		Price:    int64(product.Price),
		Stock:    int64(product.Stock),
		Category: product.Category,
		Unit:     product.Unit,
	}

	return response, nil
}

func (s *productGrpcServer) SaleProducts(ctx context.Context, in *pb.SaleProductsRequest) (*pb.SaleProductsResponse, error) {
	var saleRequests []usecase.SaleProductRequestDTO
	for _, product := range in.Products {
		saleRequests = append(saleRequests, usecase.SaleProductRequestDTO{
			ProductID: product.ProductId,
			Quantity:  int(product.Quantity),
		})
	}

	sales, err := s.saleProduct.Execute(saleRequests)
	if err != nil {
		if err == usecase.ErrStockInsufficient {
			return nil, status.Error(codes.ResourceExhausted, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	var saleResponses []*pb.SaleProductResponse
	for _, sale := range sales {
		saleResponses = append(saleResponses, &pb.SaleProductResponse{
			ProductId: sale.ProductID,
			Quantity:  int32(sale.Quantity),
		})
	}

	response := &pb.SaleProductsResponse{
		Products: saleResponses,
	}

	return response, nil
}
