package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/riyantobudi/bukadulu/internal/domain"
	"github.com/riyantobudi/bukadulu/internal/service"
)

type MissionHandler struct {
	svc *service.MissionService
}

func NewMissionHandler(svc *service.MissionService) *MissionHandler {
	return &MissionHandler{svc: svc}
}

func (h *MissionHandler) Generate(w http.ResponseWriter, r *http.Request) {
	userID := GetUserID(r.Context())
	ventureID := chi.URLParam(r, "id")

	missions, err := h.svc.Generate(ventureID, userID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"data": missions})
}

func (h *MissionHandler) List(w http.ResponseWriter, r *http.Request) {
	userID := GetUserID(r.Context())
	ventureID := chi.URLParam(r, "id")

	missions, err := h.svc.List(ventureID, userID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	if missions == nil {
		missions = []*domain.Mission{}
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"data": missions})
}

func (h *MissionHandler) Accept(w http.ResponseWriter, r *http.Request) {
	userID := GetUserID(r.Context())
	missionID := chi.URLParam(r, "missionId")

	mission, err := h.svc.Accept(missionID, userID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, mission)
}

func (h *MissionHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID := GetUserID(r.Context())
	ventureID := chi.URLParam(r, "id")

	var req domain.CreateMissionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": "invalid request"})
		return
	}
	if req.Title == "" {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": "title required"})
		return
	}

	mission, err := h.svc.Create(ventureID, userID, &req)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusCreated, mission)
}
