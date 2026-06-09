package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/riyantobudi/bukadulu/internal/service"
)

type MentorHandler struct {
	svc *service.MentorService
}

func NewMentorHandler(svc *service.MentorService) *MentorHandler {
	return &MentorHandler{svc: svc}
}

func (h *MentorHandler) ListMentees(w http.ResponseWriter, r *http.Request) {
	userID := GetUserID(r.Context())
	mentees, err := h.svc.ListMentees(userID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"data": mentees})
}

func (h *MentorHandler) AddComment(w http.ResponseWriter, r *http.Request) {
	userID := GetUserID(r.Context())
	ventureID := chi.URLParam(r, "id")

	var req struct {
		MissionID  string `json:"mission_id"`
		EvidenceID string `json:"evidence_id"`
		Content    string `json:"content"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": "invalid request"})
		return
	}
	if req.Content == "" {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": "content required"})
		return
	}

	comment, err := h.svc.AddComment(userID, ventureID, req.MissionID, req.EvidenceID, req.Content)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusCreated, comment)
}
