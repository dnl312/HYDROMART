package model

type CreateOrder struct {
	ProductID string `json:"product_id" validate:"required"`
	Qty       int    `json:"qty" validate:"required"`
}

type UpdateOrder struct {
	TransactionID string `json:"transaction_id" validate:"required"`
	Qty           int    `json:"qty" validate:"required"`
}

type DeleteOrder struct {
	TransactionID string `json:"transaction_id" validate:"required"`
}