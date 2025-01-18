package controller

import (
	"context"
	pb "merchant/pb/merchantpb"
	"merchant/repo"
)

type Merchant struct {
	pb.UnimplementedMerchantServiceServer
	Repository repo.MerchantInterface
}

func NewMerchantController(r repo.MerchantInterface) Merchant {
	return Merchant{
		Repository: r,
	}
}

func (s *Merchant) ShowAllProduct(ctx context.Context, req *pb.ShowAllProductRequest) (*pb.ShowAllProductResponse, error) {
	allProduct, err := s.Repository.GetAllProduct()
	if err != nil {
		return nil, err
	}

	allProductResponse := pb.ShowAllProductResponse{
		Products: []*pb.ProductResponse{},
	}
	for _, product := range *allProduct {
		productResponse := pb.ProductResponse{
			Id:         product.ProductID,
			MerchantId: product.MerchantID,
			Name:       product.ProductName,
			Price:      product.Price,
			Stock:      int32(product.Stock),
			Category:   product.Category,
		}
		allProductResponse.Products = append(allProductResponse.Products, &productResponse)
	}

	return &allProductResponse, nil
}
