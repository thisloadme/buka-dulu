package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/riyantobudi/bukadulu/internal/domain"
	"github.com/riyantobudi/bukadulu/internal/service"
)

type MenuHandler struct {
	svc *service.MenuService
}

func NewMenuHandler(svc *service.MenuService) *MenuHandler {
	return &MenuHandler{svc: svc}
}

func (h *MenuHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID := GetUserID(r.Context())
	ventureID := chi.URLParam(r, "id")

	var req domain.CreateMenuRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": "invalid request"})
		return
	}
	if req.Name == "" {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": "name required"})
		return
	}

	menu, err := h.svc.Create(ventureID, userID, &req)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusCreated, menu)
}

func (h *MenuHandler) List(w http.ResponseWriter, r *http.Request) {
	userID := GetUserID(r.Context())
	ventureID := chi.URLParam(r, "id")

	menus, err := h.svc.List(ventureID, userID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	if menus == nil {
		menus = []*domain.Menu{}
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"data": menus})
}

func (h *MenuHandler) Update(w http.ResponseWriter, r *http.Request) {
	userID := GetUserID(r.Context())
	menuID := chi.URLParam(r, "menuId")

	var req domain.UpdateMenuRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": "invalid request"})
		return
	}

	menu, err := h.svc.Update(menuID, userID, &req)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, menu)
}

func (h *MenuHandler) Delete(w http.ResponseWriter, r *http.Request) {
	userID := GetUserID(r.Context())
	menuID := chi.URLParam(r, "menuId")

	if err := h.svc.Delete(menuID, userID); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusNoContent, nil)
}

func (h *MenuHandler) Focus(w http.ResponseWriter, r *http.Request) {
	userID := GetUserID(r.Context())
	ventureID := chi.URLParam(r, "id")

	if err := h.svc.Focus(ventureID, userID); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "focused"})
}
