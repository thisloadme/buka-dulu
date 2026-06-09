package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/riyantobudi/bukadulu/internal/domain"
	"github.com/riyantobudi/bukadulu/internal/service"
)

type IdeaHandler struct {
	ideaSvc *service.IdeaService
}

func NewIdeaHandler(ideaSvc *service.IdeaService) *IdeaHandler {
	return &IdeaHandler{ideaSvc: ideaSvc}
}

func (h *IdeaHandler) Capture(w http.ResponseWriter, r *http.Request) {
	userID := GetUserID(r.Context())
	ventureID := chi.URLParam(r, "id")

	var req struct {
		RawInput string `json:"raw_input"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": "invalid request body"})
		return
	}
	if len(req.RawInput) < 20 {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": "raw_input must be at least 20 characters"})
		return
	}

	idea, err := h.ideaSvc.Capture(ventureID, userID, req.RawInput)
	if err != nil {
		if err.Error() == "forbidden" {
			writeJSON(w, http.StatusForbidden, map[string]string{"error": "forbidden"})
			return
		}
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusCreated, idea)
}

func (h *IdeaHandler) Process(w http.ResponseWriter, r *http.Request) {
	userID := GetUserID(r.Context())
	ventureID := chi.URLParam(r, "id")

	idea, err := h.ideaSvc.Process(ventureID, userID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, idea)
}

func (h *IdeaHandler) Get(w http.ResponseWriter, r *http.Request) {
	userID := GetUserID(r.Context())
	ventureID := chi.URLParam(r, "id")

	idea, err := h.ideaSvc.GetByVenture(ventureID, userID)
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "idea not found"})
		return
	}

	writeJSON(w, http.StatusOK, idea)
}

func (h *IdeaHandler) Update(w http.ResponseWriter, r *http.Request) {
	userID := GetUserID(r.Context())
	ventureID := chi.URLParam(r, "id")

	var req domain.UpdateIdeaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": "invalid request body"})
		return
	}

	idea, err := h.ideaSvc.Update(ventureID, userID, &req)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, idea)
}

func (h *IdeaHandler) Confirm(w http.ResponseWriter, r *http.Request) {
	userID := GetUserID(r.Context())
	ventureID := chi.URLParam(r, "id")

	idea, err := h.ideaSvc.Confirm(ventureID, userID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, idea)
}
