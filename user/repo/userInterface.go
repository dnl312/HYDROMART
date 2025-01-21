package repo

import "user/model"

type UserInterface interface {
	// list method
	ValidateUser(uuid string) (model.User, error)
	UpdateDeposit(order model.Transaction) error
	GetProductByID(id string) (model.Product, error)
	CreateOrder(order model.Transaction) error
	DeleteOrder(orderID string) error
}
