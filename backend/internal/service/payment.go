package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/riyantobudi/bukadulu/internal/domain"
	"github.com/riyantobudi/bukadulu/internal/repository"
)

// KlikQrisConfig holds credentials for the KlikQris dynamic QRIS API.
type KlikQrisConfig struct {
	APIKey     string
	MerchantID string
	BaseURL    string
}

// PaymentConfig bundles freemium + gateway settings.
type PaymentConfig struct {
	KlikQris      KlikQrisConfig
	IdeaPriceIDR  int
	FreeQuota     int
}

// PaymentService handles freemium quota + KlikQris orders.
type PaymentService struct {
	orderRepo *repository.OrderRepository
	userRepo  *repository.UserRepository
	cfg       PaymentConfig
	client    *http.Client
}

func NewPaymentService(orderRepo *repository.OrderRepository, userRepo *repository.UserRepository, cfg PaymentConfig) *PaymentService {
	return &PaymentService{
		orderRepo: orderRepo,
		userRepo:  userRepo,
		cfg:       cfg,
		client:    &http.Client{Timeout: 20 * time.Second},
	}
}

// FreeLimit returns the configured free quota per user.
func (s *PaymentService) FreeLimit() int { return s.cfg.FreeQuota }

// Price returns the configured idea validation price in IDR.
func (s *PaymentService) Price() int { return s.cfg.IdeaPriceIDR }

// PaymentOrderResult is what the API returns after CreateOrder.
type PaymentOrderResult struct {
	OrderID     string  `json:"order_id"`
	Free        bool    `json:"free"`
	QrisImage   *string `json:"qris_image,omitempty"`
	TotalAmount *string `json:"total_amount,omitempty"`
	ExpiredAt   *string `json:"expired_at,omitempty"`
	Status      string  `json:"status"`
}

// CreateOrder checks free quota first; if used up, creates a KlikQris QRIS order.
func (s *PaymentService) CreateOrder(userID, ventureID, purpose string) (*PaymentOrderResult, error) {
	// Free quota path
	used, err := s.userRepo.GetFreeQuotaUsed(userID)
	if err != nil {
		return nil, err
	}
	if used < s.cfg.FreeQuota {
		return &PaymentOrderResult{OrderID: "free", Free: true, Status: "paid"}, nil
	}

	// Reuse an existing pending order for the same venture+purpose if any
	if existing, err := s.orderRepo.FindPaidUnfulfilledForVenture(ventureID, purpose); err == nil && existing != nil {
		return s.toResult(existing), nil
	}

	amount := s.cfg.IdeaPriceIDR
	orderID := "BDL-" + uuid.New().String()[:12]

	qrisResp, err := s.callKlikQrisCreate(orderID, amount, "Validasi Ide BukaDulu")
	if err != nil {
		return nil, err
	}

	vPtr := &ventureID
	o := &domain.Order{
		ID:              uuid.New().String(),
		UserID:          userID,
		VentureID:       vPtr,
		Purpose:         purpose,
		Amount:          amount,
		TotalAmount:     strPtr(qrisResp.Data.TotalAmount),
		Status:          "pending",
		QrisURL:         strPtr(qrisResp.Data.QrisURL),
		QrisImage:       strPtr(qrisResp.Data.QrisImage),
		Signature:       strPtr(qrisResp.Data.Signature),
		KlikQrisOrderID: strPtr(qrisResp.Data.OrderID),
		ExpiredAt:       strPtr(qrisResp.Data.ExpiredAt),
	}
	if err := s.orderRepo.Create(o); err != nil {
		return nil, fmt.Errorf("create order: %w", err)
	}

	return s.toResult(o), nil
}

// HasAccess returns true if the user may process the venture's idea (free quota
// available OR a paid unfulfilled order exists for the venture).
func (s *PaymentService) HasAccess(userID, ventureID, purpose string) (bool, error) {
	used, err := s.userRepo.GetFreeQuotaUsed(userID)
	if err != nil {
		return false, err
	}
	if used < s.cfg.FreeQuota {
		return true, nil
	}
	if _, err := s.orderRepo.FindPaidUnfulfilledForVenture(ventureID, purpose); err == nil {
		return true, nil
	}
	return false, nil
}

// Fulfill marks the venture's paid order as fulfilled and consumes free quota
// if the validation was processed on free quota.
func (s *PaymentService) Fulfill(userID, ventureID, purpose string) error {
	used, err := s.userRepo.GetFreeQuotaUsed(userID)
	if err != nil {
		return err
	}
	if used < s.cfg.FreeQuota {
		// Consumed a free slot
		return s.userRepo.IncrementFreeQuota(userID)
	}
	if o, err := s.orderRepo.FindPaidUnfulfilledForVenture(ventureID, purpose); err == nil && o != nil {
		return s.orderRepo.MarkFulfilled(o.ID)
	}
	// ponytail: no-op if neither path matches (e.g. already fulfilled)
	return nil
}

// GetOrder returns an order by id.
func (s *PaymentService) GetOrder(id string) (*domain.Order, error) {
	return s.orderRepo.FindByID(id)
}

// GetOrderStatus returns a lightweight status view for polling.
func (s *PaymentService) GetOrderStatus(id string) (status string, expiredAt *string, err error) {
	o, err := s.orderRepo.FindByID(id)
	if err != nil {
		return "", nil, err
	}
	return o.Status, o.ExpiredAt, nil
}

// HandleWebhook processes a KlikQris callback payload.
// Ponytail: signature is compared against the stored create-response signature
// (per KlikQris docs). If mismatch → ignore.
func (s *PaymentService) HandleWebhook(payload []byte) error {
	var wh struct {
		OrderID   string `json:"order_id"`
		Status    string `json:"status"`
		Signature string `json:"signature"`
	}
	if err := json.Unmarshal(payload, &wh); err != nil {
		return fmt.Errorf("invalid webhook payload: %w", err)
	}

	o, err := s.orderRepo.FindByKlikQrisOrderID(wh.OrderID)
	if err != nil {
		return err
	}

	// Signature check (skip if create-time signature missing for any reason)
	if o.Signature != nil && *o.Signature != "" && wh.Signature != "" && wh.Signature != *o.Signature {
		return fmt.Errorf("webhook signature mismatch: %w", domain.ErrUnauthorized)
	}

	switch wh.Status {
	case "PAID", "SUCCESS":
		return s.orderRepo.MarkPaid(wh.OrderID, time.Now().UTC().Format(time.RFC3339))
	case "EXPIRED":
		return s.orderRepo.MarkExpired(wh.OrderID)
	}
	return nil
}

// ── KlikQris create transaction ────────────────────────────────────────────

type klikQrisCreateResp struct {
	Status bool   `json:"status"`
	Data   struct {
		OrderID     string `json:"order_id"`
		TotalAmount string `json:"total_amount"`
		QrisURL     string `json:"qris_url"`
		QrisImage   string `json:"qris_image"`
		Signature   string `json:"signature"`
		ExpiredAt   string `json:"expired_at"`
		Status      string `json:"status"`
	} `json:"data"`
}

func (s *PaymentService) callKlikQrisCreate(orderID string, amount int, keterangan string) (*klikQrisCreateResp, error) {
	body, _ := json.Marshal(map[string]interface{}{
		"order_id":   orderID,
		"id_merchant": s.cfg.KlikQris.MerchantID,
		"amount":     amount,
		"keterangan": keterangan,
	})
	req, err := http.NewRequest(http.MethodPost, s.cfg.KlikQris.BaseURL+"/qris/create", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", s.cfg.KlikQris.APIKey)
	req.Header.Set("id_merchant", s.cfg.KlikQris.MerchantID)

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("klikqris create: %w", err)
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("klikqris create failed (%d): %s", resp.StatusCode, string(respBody))
	}
	var result klikQrisCreateResp
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("parse klikqris response: %w", err)
	}
	if !result.Status {
		return nil, fmt.Errorf("klikqris rejected order: %s", string(respBody))
	}
	return &result, nil
}

func (s *PaymentService) toResult(o *domain.Order) *PaymentOrderResult {
	return &PaymentOrderResult{
		OrderID:     o.ID,
		Free:        false,
		QrisImage:   o.QrisImage,
		TotalAmount: o.TotalAmount,
		ExpiredAt:   o.ExpiredAt,
		Status:      o.Status,
	}
}

func strPtr(s string) *string { return &s }
