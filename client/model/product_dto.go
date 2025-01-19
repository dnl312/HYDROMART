package model

type ShowAllProductRequest struct {
	MerchantId string
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
