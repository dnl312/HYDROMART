package repo

import "merchant/model"

type MerchantInterface interface {
	// list method
	GetAllProduct(string) (*[]model.Product, error)
}
