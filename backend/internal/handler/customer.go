package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/riyantobudi/bukadulu/internal/domain"
	"github.com/riyantobudi/bukadulu/internal/service"
)

type CustomerHandler struct {
	svc *service.CustomerService
}

func NewCustomerHandler(svc *service.CustomerService) *CustomerHandler {
	return &CustomerHandler{svc: svc}
}

func (h *CustomerHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID := GetUserID(r.Context())
	ventureID := chi.URLParam(r, "id")

	var req domain.CreateCustomerSegmentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": "invalid request"})
		return
	}
	if req.Name == "" {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": "name required"})
		return
	}

	seg, err := h.svc.Create(ventureID, userID, &req)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusCreated, seg)
}

func (h *CustomerHandler) Get(w http.ResponseWriter, r *http.Request) {
	userID := GetUserID(r.Context())
	ventureID := chi.URLParam(r, "id")

	seg, err := h.svc.Get(ventureID, userID)
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "segment not found"})
		return
	}
	writeJSON(w, http.StatusOK, seg)
}

func (h *CustomerHandler) Confirm(w http.ResponseWriter, r *http.Request) {
	userID := GetUserID(r.Context())
	ventureID := chi.URLParam(r, "id")

	seg, err := h.svc.Confirm(ventureID, userID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, seg)
}
