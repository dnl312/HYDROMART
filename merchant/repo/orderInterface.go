package repo

import "merchant/model"

type OrderInterface interface {
	// list method
	GetOrderByProductID([]string) (*[]model.TransactionLs, error)
	UpdateOrderStatus(model.Transaction) error
	GetOrderByID(string) (model.Transaction, error)
	GetUserByID(string) (*model.User, error)
}
