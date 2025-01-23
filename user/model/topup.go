package model

type UserTopUp struct {
	TempID string `gorm:"unique"`
	OrderID string `gorm:"unique"`
	UserID string `gorm:"not null"`
}

type TopupStatus struct {
		StatusCode             string `json:"status_code"`
		TransactionID          string `json:"transaction_id"`
		GrossAmount            string `json:"gross_amount"`
		Currency               string `json:"currency"`
		OrderID                string `json:"order_id"`
		PaymentType            string `json:"payment_type"`
		SignatureKey           string `json:"signature_key"`
		TransactionStatus      string `json:"transaction_status"`
		FraudStatus            string `json:"fraud_status"`
		StatusMessage          string `json:"status_message"`
		MerchantID             string `json:"merchant_id"`
		TransactionTime        string `json:"transaction_time"`
		SettlementTime         string `json:"settlement_time"`
		ExpiryTime             string `json:"expiry_time"`
		ChannelResponseCode    string `json:"channel_response_code"`
		ChannelResponseMessage string `json:"channel_response_message"`
		Bank                   string `json:"bank"`
		ApprovalCode           string `json:"approval_code"`
		MaskedCard             string `json:"masked_card"`
		CardType               string `json:"card_type"`
		OnUs                   bool   `json:"on_us"`
	}