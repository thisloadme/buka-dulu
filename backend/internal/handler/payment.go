package handler

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/riyantobudi/bukadulu/internal/domain"
	"github.com/riyantobudi/bukadulu/internal/service"
)

// PaymentHandler handles payment + quota + history endpoints.
type PaymentHandler struct {
	svc      *service.PaymentService
	authSvc  *service.AuthService
	userRepo interface {
		GetFreeQuotaUsed(userID string) (int, error)
	}
	ventureSvc *service.VentureService
	ideaSvc    *service.IdeaService
}

func NewPaymentHandler(svc *service.PaymentService, authSvc *service.AuthService, ventureSvc *service.VentureService, ideaSvc *service.IdeaService, userRepo interface {
	GetFreeQuotaUsed(userID string) (int, error)
}) *PaymentHandler {
	return &PaymentHandler{svc: svc, authSvc: authSvc, ventureSvc: ventureSvc, ideaSvc: ideaSvc, userRepo: userRepo}
}

// GoogleLogin exchanges an OAuth code for an app JWT.
func (h *PaymentHandler) GoogleLogin(w http.ResponseWriter, r *http.Request) {
	var req service.GoogleLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Code == "" {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": "code is required"})
		return
	}
	cfg := service.GoogleOAuthConfig{
		ClientID:     h.authSvc.GoogleClientID(),
		ClientSecret: h.authSvc.GoogleClientSecret(),
		RedirectURL:  h.authSvc.GoogleRedirectURL(),
	}
	resp, err := h.authSvc.LoginWithGoogle(cfg, req.Code)
	if err != nil {
		slog.Error("google login failed", "error", err)
		writeJSON(w, http.StatusUnauthorized, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, resp)
}

// Quota reports the user's free-quota usage and price.
func (h *PaymentHandler) Quota(w http.ResponseWriter, r *http.Request) {
	userID := GetUserID(r.Context())
	used, err := h.userRepo.GetFreeQuotaUsed(userID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, domain.QuotaResponse{
		FreeUsed:  used,
		FreeLimit: h.svc.FreeLimit(),
		Price:     h.svc.Price(),
	})
}

// CreateOrder creates a free-slot OR a KlikQris QRIS order.
func (h *PaymentHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	userID := GetUserID(r.Context())
	var req domain.CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": "invalid body"})
		return
	}
	if req.Purpose == "" {
		req.Purpose = "idea_validation"
	}
	result, err := h.svc.CreateOrder(userID, req.VentureID, req.Purpose)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, result)
}

// GetOrder returns an order's full view.
func (h *PaymentHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	o, err := h.svc.GetOrder(id)
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "order not found"})
		return
	}
	writeJSON(w, http.StatusOK, o)
}

// Webhook receives KlikQris callbacks (no auth; signature-checked in service).
func (h *PaymentHandler) Webhook(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "cannot read body"})
		return
	}
	if err := h.svc.HandleWebhook(body); err != nil {
		slog.Warn("webhook processing failed", "error", err)
		// Still return 200 so KlikQris doesn't retry indefinitely
	}
	w.WriteHeader(http.StatusOK)
}

// History returns all ventures + their idea for the logged-in user.
func (h *PaymentHandler) History(w http.ResponseWriter, r *http.Request) {
	userID := GetUserID(r.Context())
	ventures, err := h.ventureSvc.ListByOwner(userID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	type historyItem struct {
		Venture *domain.Venture `json:"venture"`
		Idea    *domain.Idea    `json:"idea"`
	}
	items := make([]historyItem, 0, len(ventures))
	for _, v := range ventures {
		// Toleran: venture tanpa ide → idea nil, jangan gagal seluruh history
		idea, _ := h.ideaSvc.GetByVenture(v.ID, userID)
		if idea == nil {
			// fallback: cek repo langsung (skip ownership re-check sudah dilakukan di ListByOwner)
			idea, _ = h.ideaSvc.GetByVentureSafe(v.ID)
		}
		items = append(items, historyItem{Venture: v, Idea: idea})
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"items": items})
}
