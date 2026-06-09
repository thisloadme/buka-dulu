package service

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/riyantobudi/bukadulu/internal/domain"
	"github.com/riyantobudi/bukadulu/internal/repository"
)

type MissionService struct {
	missionRepo *repository.MissionRepository
	llmSvc      *LLMService
	ventureSvc   *VentureService
}

func NewMissionService(missionRepo *repository.MissionRepository, llmSvc *LLMService, ventureSvc *VentureService) *MissionService {
	return &MissionService{missionRepo: missionRepo, llmSvc: llmSvc, ventureSvc: ventureSvc}
}

func (s *MissionService) Generate(ventureID, userID string) ([]*domain.Mission, error) {
	_, err := s.ventureSvc.GetByID(ventureID, userID)
	if err != nil {
		return nil, err
	}

	// Check existing missions
	existing, _ := s.missionRepo.FindByVenture(ventureID)
	if len(existing) > 0 {
		return existing, nil // already generated
	}

	// Generate missions via LLM (or mock)
	mockMissions := []struct {
		Title, Description, MissionType, Priority string
		Minutes                                   int
	}{
		{"Polling ke Calon Pembeli", "Tanya 10 orang di sekitar lokasi target: apakah mereka tertarik dengan produkmu? Catat umur, budget, dan alasan mereka.", "polling", "high", 30},
		{"Wawancara 3 Responden", "Wawancara 3 orang dari target customer. Tanya tentang kebiasaan makan, budget, dan produk ideal mereka.", "interview", "high", 45},
		{"Observasi Lokasi", "Datang ke lokasi target di jam sibuk. Catat berapa banyak orang yang lalu lalang, dan usaha sejenis di sekitar.", "observation", "medium", 60},
	}

	now := time.Now().UTC().Format(time.RFC3339)
	var missions []*domain.Mission
	for i, m := range mockMissions {
		due := time.Now().AddDate(0, 0, i+1).UTC().Format(time.RFC3339)
		mission := &domain.Mission{
			ID:               uuid.New().String(),
			VentureID:        ventureID,
			Title:            m.Title,
			Description:      m.Description,
			MissionType:      m.MissionType,
			Priority:         m.Priority,
			Status:           "pending",
			DueAt:            due,
			EstimatedMinutes: m.Minutes,
			SortOrder:        i,
			CreatedAt:        now,
			UpdatedAt:        now,
		}
		if err := s.missionRepo.Create(mission); err != nil {
			return nil, fmt.Errorf("create mission: %w", err)
		}
		missions = append(missions, mission)
	}

	return missions, nil
}

func (s *MissionService) List(ventureID, userID string) ([]*domain.Mission, error) {
	_, err := s.ventureSvc.GetByID(ventureID, userID)
	if err != nil {
		return nil, err
	}
	return s.missionRepo.FindByVenture(ventureID)
}

func (s *MissionService) Accept(missionID, userID string) (*domain.Mission, error) {
	mission, err := s.missionRepo.FindByID(missionID)
	if err != nil {
		return nil, err
	}
	if _, err := s.ventureSvc.GetByID(mission.VentureID, userID); err != nil {
		return nil, err
	}

	mission.Status = "accepted"
	mission.UpdatedAt = time.Now().UTC().Format(time.RFC3339)
	if err := s.missionRepo.Update(mission); err != nil {
		return nil, err
	}

	// First mission accepted → transition stage
	if mission.Status == "accepted" {
		s.ventureSvc.TransitionStage(mission.VentureID, userID, domain.StageMissionActive)
	}
	return mission, nil
}

func (s *MissionService) Create(ventureID, userID string, req *domain.CreateMissionRequest) (*domain.Mission, error) {
	_, err := s.ventureSvc.GetByID(ventureID, userID)
	if err != nil {
		return nil, err
	}
	now := time.Now().UTC().Format(time.RFC3339)
	mission := &domain.Mission{
		ID:               uuid.New().String(),
		VentureID:        ventureID,
		Title:            req.Title,
		Description:      req.Description,
		MissionType:      req.MissionType,
		Priority:         req.Priority,
		Status:           "pending",
		DueAt:            req.DueAt,
		EvidenceRequired: req.EvidenceRequired,
		EstimatedMinutes: req.EstimatedMinutes,
		CreatedBy:        userID,
		SortOrder:        99,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
	if err := s.missionRepo.Create(mission); err != nil {
		return nil, fmt.Errorf("create mission: %w", err)
	}
	return mission, nil
}
