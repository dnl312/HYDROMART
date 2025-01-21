package controller

import (
	"context"
	"errors"
	"log"
	pb "merchant/pb/merchantpb"
	opb "merchant/pb/orderpb"
	"merchant/repo"
	"merchant/utils"
)

type MerchantOrder struct {
	pb.UnimplementedMerchantServiceServer
	Client opb.OrderServiceClient
}

func NewMerchantOrderController(r repo.MerchantInterface) MerchantOrder {
	return MerchantOrder{}
}

func (mc MerchantOrder) ShowAllOrders(ctx context.Context, req pb.ShowAllOrderRequest) (*pb.ShowAllOrderResponse, error) {
	tokenString, err := utils.GetTokenStringFromContext(ctx)
	if err != nil {
		return nil, err
	}
	user, err := utils.RecoverUser(tokenString)
	if err != nil {
		return nil, err
	}
	userRole := user["role"].(string)
	if userRole != "merchant" {
		return nil, errors.New("user not a merchant")
	}
	merchantID := user["user_id"].(string)

	r, err := mc.Client.ShowAllOrdersByMerchant(ctx, &opb.ShowAllOrdersByMerchantRequest{MerchantId: merchantID})
	if err != nil {
		log.Printf("could not show all orders: %v", err)
		return nil, err
	}
	log.Printf("show all order Response: %v", r)

	allOrderResponse := pb.ShowAllOrderResponse{
		Orders: []*pb.Order{},
	}
	for _, order := range r.Orders {
		orderResponse := pb.Order{
			OrderId:   order.OrderId,
			UserId:    order.UserId,
			ProductId: order.ProductId,
			Qty:       order.Qty,
			Amount:    order.Amount,
			Status:    order.Status,
		}
		allOrderResponse.Orders = append(allOrderResponse.Orders, &orderResponse)
	}

	return &allOrderResponse, nil
}

func (mc MerchantOrder) ProcessOrder(ctx context.Context, req *pb.ProcessOrderRequest) error {
	tokenString, err := utils.GetTokenStringFromContext(ctx)
	if err != nil {
		return err
	}
	user, err := utils.RecoverUser(tokenString)
	if err != nil {
		return err
	}
	userRole := user["role"].(string)
	if userRole != "merchant" {
		return errors.New("user not a merchant")
	}
	merchantID := user["user_id"].(string)

	order, err := mc.Client.GetOrderById(ctx, &opb.GetOrderByIdRequest{OrderId: req.OrderId})
	if err != nil {
		log.Printf("could not process order: %v", err)
		return err
	}

	if merchantID != order.UserId {
		return errors.New("merchant don't own the product")
	}

	r, err := mc.Client.UpdateOrder(ctx, &opb.UpdateOrderRequest{
		OrderId:   order.OrderId,
		UserId:    order.UserId,
		ProductId: order.ProductId,
		Qty:       order.Qty,
		Amount:    order.Amount,
		Status:    "processed",
	})
	if err != nil {
		log.Printf("could not process order: %v", err)
		return err
	}
	log.Printf("process Response: %v", r)

	return nil
}
