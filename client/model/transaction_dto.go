package model

type CreateOrder struct {
	ProductID string `json:"ProductID" validate:"required"`
	Qty       int    `json:"Qty" validate:"required"`
}
