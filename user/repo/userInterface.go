package repo

import "user/model"

type UserInterface interface {
	// list method
	GetProductByID(id string) (*model.Product, error)
	CreateOrder(order *model.Transaction) error
}
