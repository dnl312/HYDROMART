package controller

import (
	"context"
	"merchant/model"
	pb "merchant/pb/merchantpb"
	"merchant/repo"
	"merchant/utils"

	"github.com/google/uuid"
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

func (m *Merchant) ShowAllProducts(ctx context.Context, req *pb.ShowAllProductRequest) (*pb.ShowAllProductResponse, error) {
	tokenString, err := utils.GetTokenStringFromContext(ctx)
	if err != nil {
		return nil, err
	}
	user, err := utils.RecoverUser(tokenString)
	if err != nil {
		return nil, err
	}
	merchantID := user["user_id"].(string)

	allProduct, err := m.Repository.GetAllProduct(merchantID)
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

func (m *Merchant) AddProduct(ctx context.Context, req *pb.AddProductRequest) (*pb.AddProductResponse, error) {
	tokenString, err := utils.GetTokenStringFromContext(ctx)
	if err != nil {
		return nil, err
	}
	user, err := utils.RecoverUser(tokenString)
	if err != nil {
		return nil, err
	}
	merchantID := user["user_id"].(string)

	product := model.Product{
		ProductID:   uuid.New().String(),
		MerchantID:  merchantID,
		ProductName: req.Product.Name,
		Price:       req.Product.Price,
		Stock:       int(req.Product.Stock),
		Category:    req.Product.Category,
	}
	err = m.Repository.AddProduct(&product)
	if err != nil {
		return nil, err
	}

	return &pb.AddProductResponse{
		Message: "the product has been added",
	}, nil
}

func (m *Merchant) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {
	tokenString, err := utils.GetTokenStringFromContext(ctx)
	if err != nil {
		return nil, err
	}
	user, err := utils.RecoverUser(tokenString)
	if err != nil {
		return nil, err
	}
	merchantID := user["user_id"].(string)

	product := model.Product{
		ProductID:   req.Product.Id,
		MerchantID:  merchantID,
		ProductName: req.Product.Name,
		Price:       req.Product.Price,
		Stock:       int(req.Product.Stock),
		Category:    req.Product.Category,
	}
	err = m.Repository.UpdateProduct(&product)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateProductResponse{
		Message: "product has been updated",
	}, nil
}

func (m *Merchant) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	err := m.Repository.DeleteProduct(req.ProductId)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteProductResponse{
		Message: "product has been deleted",
	}, nil
}
