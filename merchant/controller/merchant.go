package controller

import (
	"context"
	"merchant/model"
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

func (s *Merchant) ShowAllProducts(ctx context.Context, req *pb.ShowAllProductRequest) (*pb.ShowAllProductResponse, error) {

	allProduct, err := s.Repository.GetAllProduct(req.MerchantId)
	if err != nil {
		return nil, err
	}

	allProductResponse := pb.ShowAllProductResponse{
		Products: []*pb.Product{},
	}
	for _, product := range *allProduct {
		productResponse := pb.Product{
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

func (s *Merchant) AddProduct(ctx context.Context, req *pb.AddProductRequest) (*pb.AddProductResponse, error) {
	product := model.Product{
		MerchantID:  req.MerchantId,
		ProductName: req.Product.Name,
		Price:       req.Product.Price,
		Stock:       int(req.Product.Stock),
		Category:    req.Product.Category,
	}
	err := s.Repository.AddProduct(&product)
	if err != nil {
		return nil, err
	}

	return &pb.AddProductResponse{
		Message: "product has been added",
	}, nil
}

func (s *Merchant) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {
	product := model.Product{
		MerchantID:  req.MerchantId,
		ProductName: req.Product.Name,
		Price:       req.Product.Price,
		Stock:       int(req.Product.Stock),
		Category:    req.Product.Category,
	}
	err := s.Repository.UpdateProduct(&product)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateProductResponse{
		Message: "product has been updated",
	}, nil
}

func (s *Merchant) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	err := s.Repository.DeleteProduct(req.ProductId)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteProductResponse{
		Message: "product has been deleted",
	}, nil
}
