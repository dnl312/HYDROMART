package repo

import "merchant/model"

type MerchantInterface interface {
	// list method
	GetAllProduct() (*[]model.Product, error)
}
