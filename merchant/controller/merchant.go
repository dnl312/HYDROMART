package controller

import (
	"context"
	"errors"
	"fmt"
	"log"
	"merchant/model"
	pb "merchant/pb/merchantpb"
	"merchant/repo"
	"merchant/service"
	"merchant/utils"

	"github.com/google/uuid"
)

type Merchant struct {
	pb.UnimplementedMerchantServiceServer
	merchantRepo repo.MerchantInterface
	orderRepo    repo.OrderInterface
	Mb         service.MessageBroker
}

func NewMerchantController(mr repo.MerchantInterface, or repo.OrderInterface, mb service.MessageBroker) Merchant {
	return Merchant{
		merchantRepo: mr,
		orderRepo:    or,
		Mb:         mb,
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
	userRole := user["role"].(string)
	if userRole != "MERCHANT" {
		return nil, errors.New("user not a merchant")
	}
	merchantID := user["user_id"].(string)

	allProduct, err := m.merchantRepo.GetAllProduct(merchantID)
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
	userRole := user["role"].(string)
	if userRole != "MERCHANT" {
		return nil, errors.New("user not a merchant")
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
	err = m.merchantRepo.AddProduct(&product)
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
	userRole := user["role"].(string)
	if userRole != "MERCHANT" {
		return nil, errors.New("user not a merchant")
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
	err = m.merchantRepo.UpdateProduct(&product)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateProductResponse{
		Message: "product has been updated",
	}, nil
}

func (m *Merchant) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	tokenString, err := utils.GetTokenStringFromContext(ctx)
	if err != nil {
		return nil, err
	}
	user, err := utils.RecoverUser(tokenString)
	if err != nil {
		return nil, err
	}
	userRole := user["role"].(string)
	if userRole != "MERCHANT" {
		return nil, errors.New("user not a merchant")
	}
	merchantID := user["user_id"].(string)

	err = m.merchantRepo.DeleteProduct(req.ProductId, merchantID)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteProductResponse{
		Message: "product has been deleted",
	}, nil
}

func (mc Merchant) ShowAllOrders(ctx context.Context, req *pb.ShowAllOrderRequest) (*pb.ShowAllOrderResponse, error) {
	tokenString, err := utils.GetTokenStringFromContext(ctx)
	if err != nil {
		return nil, err
	}
	user, err := utils.RecoverUser(tokenString)
	if err != nil {
		return nil, err
	}
	userRole := user["role"].(string)
	if userRole != "MERCHANT" {
		return nil, errors.New("user not a merchant")
	}
	merchantID := user["user_id"].(string)

	productIds, err := mc.GetProductIDs(merchantID)
	if err != nil {
		return nil, err
	}

	r, err := mc.orderRepo.GetOrderByProductID(productIds)
	if err != nil {
		log.Printf("could not show all orders: %v", err)
		return nil, err
	}
	log.Printf("show all order Response: %v", r)

	allOrderResponse := pb.ShowAllOrderResponse{
		Orders: []*pb.Order{},
	}
	for _, order := range *r {
		orderResponse := pb.Order{
			OrderId:   order.TransactionID,
			UserId:    order.UserID,
			ProductId: order.ProductID,
			Qty:       int32(order.Qty),
			Amount:    order.Amount,
			Status:    order.Status,
		}
		allOrderResponse.Orders = append(allOrderResponse.Orders, &orderResponse)
	}

	return &allOrderResponse, nil
}

func (mc Merchant) GetProductIDs(merchantID string) ([]string, error) {
	allProduct, err := mc.merchantRepo.GetAllProduct(merchantID)
	if err != nil {
		return nil, err
	}

	allProductIds := []string{}
	for _, product := range *allProduct {
		allProductIds = append(allProductIds, product.ProductID)
	}

	return allProductIds, nil
}

func (mc Merchant) ProcessOrder(ctx context.Context, req *pb.ProcessOrderRequest) (*pb.ProcessOrderResponse, error) {
	tokenString, err := utils.GetTokenStringFromContext(ctx)
	if err != nil {
		return nil, err
	}
	user, err := utils.RecoverUser(tokenString)
	if err != nil {
		return nil, err
	}
	userRole := user["role"].(string)
	if userRole != "MERCHANT" {
		return nil, errors.New("user not a merchant")
	}
	merchantID := user["user_id"].(string)

	fmt.Printf("\n\n\n%s\n\n\n", req.OrderId)

	order, err := mc.orderRepo.GetOrderByID(req.OrderId)
	if err != nil {
		log.Printf("could not process order: %v", err)
		return nil, err
	}

	product, err := mc.merchantRepo.GetProductByID(order.ProductID)
	if err != nil {
		log.Printf("could not process order: %v", err)
		return nil, err
	}

	if merchantID != product.MerchantID {
		return nil, errors.New("merchant don't own the product")
	}

	err = mc.orderRepo.UpdateOrderStatus(model.Transaction{
		TransactionID: order.TransactionID,
		UserID:        order.UserID,
		ProductID:     order.ProductID,
		Qty:           order.Qty,
		Amount:        order.Amount,
		Status:        "PROCESSED",
	})
	if err != nil {
		log.Printf("could not process order: %v", err)
		return nil, err
	}

	customer, err := mc.orderRepo.GetUserByID(order.UserID)
	if err != nil {
		log.Printf("could not process order: %v", err)
		return nil, err
	}

	err = utils.SendMail("hydromart@admin.com", customer.Email, order.ProductID, order.Qty, order.Amount, mc.Mb)
	if err != nil {
		log.Printf("could not send email %v", err)
		return nil, err
	}

	return &pb.ProcessOrderResponse{Message: "order processed"}, nil
}
