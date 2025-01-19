package model

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
	Name     string
	Price    float64
	Stock    int
	Category string
}
