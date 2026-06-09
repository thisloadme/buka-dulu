package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/riyantobudi/bukadulu/internal/domain"
	"github.com/riyantobudi/bukadulu/internal/service"
)

type VentureHandler struct {
	ventureSvc *service.VentureService
}

func NewVentureHandler(ventureSvc *service.VentureService) *VentureHandler {
	return &VentureHandler{ventureSvc: ventureSvc}
}

func (h *VentureHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID := GetUserID(r.Context())

	var req domain.CreateVentureRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": "invalid request body"})
		return
	}
	if req.Name == "" {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": "name is required"})
		return
	}

	venture, err := h.ventureSvc.Create(userID, &req)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusCreated, venture)
}

func (h *VentureHandler) List(w http.ResponseWriter, r *http.Request) {
	userID := GetUserID(r.Context())
	ventures, err := h.ventureSvc.ListByOwner(userID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	if ventures == nil {
		ventures = []*domain.Venture{}
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"data": ventures})
}

func (h *VentureHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	userID := GetUserID(r.Context())
	ventureID := chi.URLParam(r, "id")

	venture, err := h.ventureSvc.GetByID(ventureID, userID)
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "venture not found"})
		return
	}

	writeJSON(w, http.StatusOK, venture)
}

func (h *VentureHandler) Update(w http.ResponseWriter, r *http.Request) {
	userID := GetUserID(r.Context())
	ventureID := chi.URLParam(r, "id")

	var req domain.UpdateVentureRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": "invalid request body"})
		return
	}

	venture, err := h.ventureSvc.Update(ventureID, userID, &req)
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, venture)
}
