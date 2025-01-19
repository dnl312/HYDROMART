package repo

import "merchant/model"

type MerchantInterface interface {
	// list method
	GetAllProduct(string) (*[]model.Product, error)
	AddProduct(*model.Product) error
	UpdateProduct(*model.Product) error
	DeleteProduct(string) error
}
