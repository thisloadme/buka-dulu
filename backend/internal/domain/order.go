package domain

// Order represents a payment order (KlikQris dynamic QRIS).
type Order struct {
	ID             string  `db:"id" json:"id"`
	UserID         string  `db:"user_id" json:"user_id"`
	VentureID      *string `db:"venture_id" json:"venture_id,omitempty"`
	Purpose        string  `db:"purpose" json:"purpose"`
	Amount         int     `db:"amount" json:"amount"`
	TotalAmount    *string `db:"total_amount" json:"total_amount,omitempty"`
	Status         string  `db:"status" json:"status"`
	QrisURL        *string `db:"qris_url" json:"qris_url,omitempty"`
	QrisImage      *string `db:"qris_image" json:"qris_image,omitempty"`
	Signature      *string `db:"signature" json:"signature,omitempty"`
	KlikQrisOrderID *string `db:"klikqris_order_id" json:"klikqris_order_id,omitempty"`
	ExpiredAt      *string `db:"expired_at" json:"expired_at,omitempty"`
	PaidAt         *string `db:"paid_at" json:"paid_at,omitempty"`
	Fulfilled      bool    `db:"fulfilled" json:"fulfilled"`
	CreatedAt      string  `db:"created_at" json:"created_at"`
	UpdatedAt      string  `db:"updated_at" json:"updated_at"`
}

// CreateOrderRequest is the body for creating a payment order.
type CreateOrderRequest struct {
	VentureID string `json:"venture_id"`
	Purpose   string `json:"purpose"`
}

// QuotaResponse reports the user's free-quota usage and pricing.
type QuotaResponse struct {
	FreeUsed  int `json:"free_used"`
	FreeLimit int `json:"free_limit"`
	Price     int `json:"price"`
}
