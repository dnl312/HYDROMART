package model

type ShowAllProductRequest struct {
	MerchantId string `json:"merchant_id" validate:"required,merchant_id"`
}

type ShowAllProductResponse struct {
	Products []ProductResponse
}

type ProductResponse struct {
	ProductID  string
	MerchantId string
	Name       string
	Price      string
	Stock      int
	Category   string
}

type AddProductRequest struct {
	product Product
}
