package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/riyantobudi/bukadulu/internal/domain"
	"github.com/riyantobudi/bukadulu/internal/service"
)

type EvidenceHandler struct {
	svc *service.EvidenceService
}

func NewEvidenceHandler(svc *service.EvidenceService) *EvidenceHandler {
	return &EvidenceHandler{svc: svc}
}

func (h *EvidenceHandler) Upload(w http.ResponseWriter, r *http.Request) {
	userID := GetUserID(r.Context())
	ventureID := chi.URLParam(r, "id")

	var req domain.UploadEvidenceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": "invalid request"})
		return
	}
	if req.MissionID == "" || req.EvidenceType == "" {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": "mission_id and evidence_type required"})
		return
	}

	ev, err := h.svc.Upload(ventureID, userID, &req)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusCreated, ev)
}

func (h *EvidenceHandler) ListByMission(w http.ResponseWriter, r *http.Request) {
	missionID := chi.URLParam(r, "missionId")

	ev, err := h.svc.ListByMission(missionID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	if ev == nil {
		ev = []*domain.Evidence{}
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"data": ev})
}

func (h *EvidenceHandler) GetWithReview(w http.ResponseWriter, r *http.Request) {
	evidenceID := chi.URLParam(r, "evidenceId")

	ev, err := h.svc.GetWithReview(evidenceID)
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "evidence not found"})
		return
	}
	writeJSON(w, http.StatusOK, ev)
}

func (h *EvidenceHandler) Review(w http.ResponseWriter, r *http.Request) {
	evidenceID := chi.URLParam(r, "evidenceId")

	review, err := h.svc.Review(evidenceID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, review)
}

func (h *EvidenceHandler) Override(w http.ResponseWriter, r *http.Request) {
	userID := GetUserID(r.Context())
	reviewID := chi.URLParam(r, "reviewId")

	var req domain.OverrideReviewRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": "invalid request"})
		return
	}

	if err := h.svc.OverrideReview(reviewID, userID, req.Verdict, req.Rationale); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "overridden"})
}
