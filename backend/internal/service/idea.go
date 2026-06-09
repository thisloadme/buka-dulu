package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/riyantobudi/bukadulu/internal/domain"
	"github.com/riyantobudi/bukadulu/internal/repository"
)

type IdeaService struct {
	ideaRepo   *repository.IdeaRepository
	llmService *LLMService
	ventureSvc *VentureService
}

func NewIdeaService(ideaRepo *repository.IdeaRepository, llmService *LLMService, ventureSvc *VentureService) *IdeaService {
	return &IdeaService{
		ideaRepo:   ideaRepo,
		llmService: llmService,
		ventureSvc: ventureSvc,
	}
}

func (s *IdeaService) Capture(ventureID, userID, rawInput string) (*domain.Idea, error) {
	_, err := s.ventureSvc.GetByID(ventureID, userID)
	if err != nil {
		return nil, err
	}

	if len(rawInput) < 20 {
		return nil, fmt.Errorf("raw input must be at least 20 characters: %w", domain.ErrInvalidInput)
	}

	existing, err := s.ideaRepo.FindByVenture(ventureID)
	if err == nil && existing != nil && !existing.IsLocked {
		existing.RawInput = rawInput
		existing.Status = "pending"
		existing.UpdatedAt = time.Now().UTC().Format(time.RFC3339)
		if err := s.ideaRepo.Update(existing); err != nil {
			return nil, fmt.Errorf("update idea: %w", err)
		}
		return existing, nil
	}

	now := time.Now().UTC().Format(time.RFC3339)
	idea := &domain.Idea{
		ID:        uuid.New().String(),
		VentureID: ventureID,
		RawInput:  rawInput,
		Version:   1,
		IsLocked:  false,
		Status:    "pending",
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := s.ideaRepo.Create(idea); err != nil {
		return nil, fmt.Errorf("create idea: %w", err)
	}
	return idea, nil
}

func (s *IdeaService) Process(ventureID, userID string) (*domain.Idea, error) {
	idea, err := s.ideaRepo.FindByVenture(ventureID)
	if err != nil {
		return nil, err
	}

	if idea.IsLocked {
		return nil, fmt.Errorf("idea already locked: %w", domain.ErrInvalidInput)
	}

	idea.Status = "processing"
	s.ideaRepo.Update(idea)

	concept, err := s.llmService.StructureIdea(idea.RawInput)
	if err != nil {
		idea.Status = "failed"
		s.ideaRepo.Update(idea)
		return nil, fmt.Errorf("AI structuring failed: %w", err)
	}

	assumptionsJSON, _ := json.Marshal(concept.KeyAssumptions)
	risksJSON, _ := json.Marshal(concept.EarlyRisks)
	assumptionsStr := string(assumptionsJSON)
	risksStr := string(risksJSON)

	idea.OneLineConcept = &concept.OneLineConcept
	idea.TargetCustomer = &concept.TargetCustomer
	idea.ValueProposition = &concept.ValueProposition
	idea.KeyAssumptions = &assumptionsStr
	idea.EarlyRisks = &risksStr
	idea.Status = "done"
	idea.UpdatedAt = time.Now().UTC().Format(time.RFC3339)

	if err := s.ideaRepo.Update(idea); err != nil {
		return nil, fmt.Errorf("update idea: %w", err)
	}

	return idea, nil
}

func (s *IdeaService) GetByVenture(ventureID, userID string) (*domain.Idea, error) {
	_, err := s.ventureSvc.GetByID(ventureID, userID)
	if err != nil {
		return nil, err
	}
	return s.ideaRepo.FindByVenture(ventureID)
}

func (s *IdeaService) Update(ventureID, userID string, req *domain.UpdateIdeaRequest) (*domain.Idea, error) {
	idea, err := s.GetByVenture(ventureID, userID)
	if err != nil {
		return nil, err
	}

	if idea.IsLocked {
		return nil, fmt.Errorf("idea is locked, cannot edit: %w", domain.ErrInvalidInput)
	}

	if req.OneLineConcept != nil {
		idea.OneLineConcept = req.OneLineConcept
	}
	if req.TargetCustomer != nil {
		idea.TargetCustomer = req.TargetCustomer
	}
	if req.ValueProposition != nil {
		idea.ValueProposition = req.ValueProposition
	}
	if req.KeyAssumptions != nil {
		idea.KeyAssumptions = req.KeyAssumptions
	}
	if req.EarlyRisks != nil {
		idea.EarlyRisks = req.EarlyRisks
	}
	idea.UpdatedAt = time.Now().UTC().Format(time.RFC3339)

	if err := s.ideaRepo.Update(idea); err != nil {
		return nil, fmt.Errorf("update idea: %w", err)
	}
	return idea, nil
}

func (s *IdeaService) Confirm(ventureID, userID string) (*domain.Idea, error) {
	idea, err := s.GetByVenture(ventureID, userID)
	if err != nil {
		return nil, err
	}

	if idea.Status != "done" && idea.OneLineConcept == nil {
		return nil, fmt.Errorf("idea must be processed first: %w", domain.ErrStageGate)
	}

	idea.IsLocked = true
	idea.UpdatedAt = time.Now().UTC().Format(time.RFC3339)
	if err := s.ideaRepo.Update(idea); err != nil {
		return nil, fmt.Errorf("lock idea: %w", err)
	}

	if _, err := s.ventureSvc.TransitionStage(ventureID, userID, domain.StageIdeaDefined); err != nil {
		return nil, err
	}

	return idea, nil
}
