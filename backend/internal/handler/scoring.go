package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/riyantobudi/bukadulu/internal/service"
)

type ScoringHandler struct {
	svc *service.ScoringService
}

func NewScoringHandler(svc *service.ScoringService) *ScoringHandler {
	return &ScoringHandler{svc: svc}
}

func (h *ScoringHandler) GetScore(w http.ResponseWriter, r *http.Request) {
	userID := GetUserID(r.Context())
	ventureID := chi.URLParam(r, "id")

	score, err := h.svc.GetLatest(ventureID, userID)
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "score not found"})
		return
	}
	writeJSON(w, http.StatusOK, score)
}

func (h *ScoringHandler) Calculate(w http.ResponseWriter, r *http.Request) {
	userID := GetUserID(r.Context())
	ventureID := chi.URLParam(r, "id")

	score, err := h.svc.Calculate(ventureID, userID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, score)
}

func (h *ScoringHandler) GenerateDecision(w http.ResponseWriter, r *http.Request) {
	userID := GetUserID(r.Context())
	ventureID := chi.URLParam(r, "id")

	decision, err := h.svc.GenerateDecision(ventureID, userID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	// Also return the score
	score, _ := h.svc.GetLatest(ventureID, userID)
	writeJSON(w, http.StatusOK, map[string]interface{}{
		"decision": decision,
		"score":    score,
	})
}

func (h *ScoringHandler) GetDecision(w http.ResponseWriter, r *http.Request) {
	userID := GetUserID(r.Context())
	ventureID := chi.URLParam(r, "id")

	decision, err := h.svc.GetDecision(ventureID, userID)
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "decision not found"})
		return
	}
	writeJSON(w, http.StatusOK, decision)
}
