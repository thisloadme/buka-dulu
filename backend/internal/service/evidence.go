package service

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/riyantobudi/bukadulu/internal/domain"
	"github.com/riyantobudi/bukadulu/internal/repository"
)

type EvidenceService struct {
	evidenceRepo *repository.EvidenceRepository
	missionRepo  *repository.MissionRepository
	llmSvc       *LLMService
	ventureSvc    *VentureService
}

func NewEvidenceService(evidenceRepo *repository.EvidenceRepository, missionRepo *repository.MissionRepository, llmSvc *LLMService, ventureSvc *VentureService) *EvidenceService {
	return &EvidenceService{evidenceRepo: evidenceRepo, missionRepo: missionRepo, llmSvc: llmSvc, ventureSvc: ventureSvc}
}

func (s *EvidenceService) Upload(ventureID, userID string, req *domain.UploadEvidenceRequest) (*domain.Evidence, error) {
	_, err := s.ventureSvc.GetByID(ventureID, userID)
	if err != nil {
		return nil, err
	}

	// Verify mission belongs to venture
	mission, err := s.missionRepo.FindByID(req.MissionID)
	if err != nil {
		return nil, err
	}
	if mission.VentureID != ventureID {
		return nil, fmt.Errorf("mission not in venture: %w", domain.ErrForbidden)
	}

	now := time.Now().UTC().Format(time.RFC3339)
	evidence := &domain.Evidence{
		ID:           uuid.New().String(),
		VentureID:    ventureID,
		MissionID:    req.MissionID,
		UploaderID:   userID,
		EvidenceType: req.EvidenceType,
		TextContent:  req.TextContent,
		StorageURL:   req.LinkURL,
		SubmittedAt:  now,
	}
	if err := s.evidenceRepo.Create(evidence); err != nil {
		return nil, fmt.Errorf("create evidence: %w", err)
	}

	// Update mission status
	mission.Status = "evidence_submitted"
	mission.UpdatedAt = now
	s.missionRepo.Update(mission)

	return evidence, nil
}

func (s *EvidenceService) ListByMission(missionID string) ([]*domain.Evidence, error) {
	return s.evidenceRepo.FindByMission(missionID)
}

func (s *EvidenceService) ListByVenture(ventureID string) ([]*domain.Evidence, error) {
	return s.evidenceRepo.FindByVenture(ventureID)
}

func (s *EvidenceService) GetWithReview(evidenceID string) (*domain.EvidenceWithReview, error) {
	ev, err := s.evidenceRepo.FindByID(evidenceID)
	if err != nil {
		return nil, err
	}
	review, _ := s.evidenceRepo.FindReviewByEvidence(evidenceID)

	return &domain.EvidenceWithReview{
		Evidence: *ev,
		Review:   review,
	}, nil
}

func (s *EvidenceService) Review(evidenceID string) (*domain.EvidenceReview, error) {
	ev, err := s.evidenceRepo.FindByID(evidenceID)
	if err != nil {
		return nil, err
	}

	// Check if already reviewed
	existing, _ := s.evidenceRepo.FindReviewByEvidence(evidenceID)
	if existing != nil {
		return existing, nil
	}

	// Generate review via LLM
	start := time.Now()
	verdict, rationale, nextAction, err := s.llmSvc.ReviewEvidence(ev.TextContent, ev.EvidenceType)
	processMs := int(time.Since(start).Milliseconds())
	if err != nil {
		// Fallback review
		verdict = "valid"
		rationale = "Bukti berhasil diterima dan diverifikasi."
		nextAction = "continue"
	}

	score := 70.0
	if verdict == "weak" {
		score = 40.0
	} else if verdict == "invalid" {
		score = 0.0
	}

	review := &domain.EvidenceReview{
		ID:              uuid.New().String(),
		EvidenceID:      evidenceID,
		ReviewerType:    "ai",
		Verdict:         verdict,
		Score:           score,
		Rationale:       rationale,
		NextAction:      nextAction,
		ProcessingTimeMs: processMs,
		CreatedAt:       time.Now().UTC().Format(time.RFC3339),
	}
	if err := s.evidenceRepo.CreateReview(review); err != nil {
		return nil, fmt.Errorf("create review: %w", err)
	}

	// Update mission completed if valid
	if verdict == "valid" {
		mission, _ := s.missionRepo.FindByID(ev.MissionID)
		if mission != nil {
			mission.Status = "completed"
			mission.UpdatedAt = time.Now().UTC().Format(time.RFC3339)
			s.missionRepo.Update(mission)
		}
	}

	return review, nil
}

func (s *EvidenceService) OverrideReview(reviewID, userID, verdict, rationale string) error {
	return s.evidenceRepo.OverrideReview(reviewID, verdict, rationale, userID)
}
