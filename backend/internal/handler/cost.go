package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/riyantobudi/bukadulu/internal/domain"
	"github.com/riyantobudi/bukadulu/internal/service"
)

type CostHandler struct {
	svc *service.CostService
}

func NewCostHandler(svc *service.CostService) *CostHandler {
	return &CostHandler{svc: svc}
}

func (h *CostHandler) AddIngredient(w http.ResponseWriter, r *http.Request) {
	userID := GetUserID(r.Context())
	ventureID := chi.URLParam(r, "id")

	var req domain.CreateIngredientRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": "invalid request"})
		return
	}

	ing, err := h.svc.AddIngredient(ventureID, userID, &req)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusCreated, ing)
}

func (h *CostHandler) ListIngredients(w http.ResponseWriter, r *http.Request) {
	ventureID := chi.URLParam(r, "id")
	menuID := r.URL.Query().Get("menu_id")
	if menuID == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "menu_id required"})
		return
	}

	ingredients, err := h.svc.GetIngredients(ventureID, menuID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	if ingredients == nil {
		ingredients = []*domain.Ingredient{}
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"data": ingredients})
}

func (h *CostHandler) DeleteIngredient(w http.ResponseWriter, r *http.Request) {
	ingredientID := chi.URLParam(r, "ingredientId")
	if err := h.svc.DeleteIngredient(ingredientID); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusNoContent, nil)
}

func (h *CostHandler) AddPackaging(w http.ResponseWriter, r *http.Request) {
	userID := GetUserID(r.Context())
	ventureID := chi.URLParam(r, "id")

	var req domain.CreatePackagingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": "invalid request"})
		return
	}

	p, err := h.svc.AddPackaging(ventureID, userID, &req)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusCreated, p)
}

func (h *CostHandler) ListPackaging(w http.ResponseWriter, r *http.Request) {
	menuID := r.URL.Query().Get("menu_id")
	if menuID == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "menu_id required"})
		return
	}

	items, err := h.svc.GetPackaging(menuID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	if items == nil {
		items = []*domain.PackagingCost{}
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"data": items})
}

func (h *CostHandler) Calculate(w http.ResponseWriter, r *http.Request) {
	userID := GetUserID(r.Context())
	ventureID := chi.URLParam(r, "id")
	menuID := chi.URLParam(r, "menuId")

	var cfg domain.CostConfigRequest
	json.NewDecoder(r.Body).Decode(&cfg) // optional body

	cs, err := h.svc.Calculate(ventureID, userID, menuID, &cfg)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, cs)
}

func (h *CostHandler) GetSummary(w http.ResponseWriter, r *http.Request) {
	ventureID := chi.URLParam(r, "id")
	menuID := chi.URLParam(r, "menuId")

	cs, err := h.svc.GetSummary(ventureID, menuID)
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "summary not found"})
		return
	}
	writeJSON(w, http.StatusOK, cs)
}

func (h *CostHandler) GetAllSummaries(w http.ResponseWriter, r *http.Request) {
	ventureID := chi.URLParam(r, "id")

	summaries, err := h.svc.GetAllSummaries(ventureID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	if summaries == nil {
		summaries = []*domain.CostSummary{}
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"data": summaries})
}

func (h *CostHandler) Confirm(w http.ResponseWriter, r *http.Request) {
	userID := GetUserID(r.Context())
	ventureID := chi.URLParam(r, "id")

	if err := h.svc.Confirm(ventureID, userID); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "cost_locked"})
}
