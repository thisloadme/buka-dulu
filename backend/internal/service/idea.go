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
	paymentSvc *PaymentService
}

func NewIdeaService(ideaRepo *repository.IdeaRepository, llmService *LLMService, ventureSvc *VentureService, paymentSvc *PaymentService) *IdeaService {
	return &IdeaService{
		ideaRepo:   ideaRepo,
		llmService: llmService,
		ventureSvc: ventureSvc,
		paymentSvc: paymentSvc,
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

	// Paywall: only allow if user has free quota OR a paid order for this venture
	if s.paymentSvc != nil {
		ok, err := s.paymentSvc.HasAccess(userID, ventureID, "idea_validation")
		if err != nil {
			return nil, fmt.Errorf("check access: %w", err)
		}
		if !ok {
			return nil, fmt.Errorf("payment required: %w", domain.ErrPaymentRequired)
		}
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

	// Consume quota / fulfill order after successful AI processing
	if s.paymentSvc != nil {
		_ = s.paymentSvc.Fulfill(userID, ventureID, "idea_validation")
	}

	return idea, nil
}

func (s *IdeaService) GetByVenture(ventureID, userID string) (*domain.Idea, error) {
	_, err := s.ventureSvc.GetByID(ventureID, userID)
	if err != nil {
		return nil, err
	}
	idea, err := s.ideaRepo.FindByVenture(ventureID)
	if err == domain.ErrNotFound {
		return nil, nil
	}
	return idea, err
}

// GetByVentureSafe returns the idea for a venture without ownership check.
// Use only when the caller has already filtered by owner (e.g. ListByOwner).
// Returns nil, nil if no idea exists (instead of an error) for graceful aggregation.
func (s *IdeaService) GetByVentureSafe(ventureID string) (*domain.Idea, error) {
	idea, err := s.ideaRepo.FindByVenture(ventureID)
	if err == domain.ErrNotFound {
		return nil, nil
	}
	return idea, err
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

// Refine runs one AI iteration on the idea using the user's instruction,
// preserves conversation history between turns, and applies the updated concept.
// Only allowed after the idea has been processed once (status 'done').
func (s *IdeaService) Refine(ventureID, userID, instruction string) (*domain.Idea, *RefineIdeaResult, error) {
	idea, err := s.GetByVenture(ventureID, userID)
	if err != nil {
		return nil, nil, err
	}
	if idea == nil {
		return nil, nil, fmt.Errorf("idea not found: %w", domain.ErrNotFound)
	}
	if idea.IsLocked {
		return nil, nil, fmt.Errorf("idea is locked, cannot refine: %w", domain.ErrInvalidInput)
	}
	if idea.Status != "done" || idea.OneLineConcept == nil {
		return nil, nil, fmt.Errorf("idea must be processed first: %w", domain.ErrStageGate)
	}

	// Reconstruct conversation history
	history := []Message{}
	if idea.RefineHistory != nil && *idea.RefineHistory != "" {
		_ = json.Unmarshal([]byte(*idea.RefineHistory), &history)
	}

	currentConcept := ""
	if idea.OneLineConcept != nil {
		currentConcept = *idea.OneLineConcept
	}

	result, err := s.llmService.RefineIdea(idea.RawInput, currentConcept, history, instruction)
	if err != nil {
		return nil, nil, fmt.Errorf("AI refine failed: %w", err)
	}

	// Apply updated fields
	idea.OneLineConcept = &result.OneLineConcept
	idea.TargetCustomer = &result.TargetCustomer
	idea.ValueProposition = &result.ValueProposition
	assumptionsJSON, _ := json.Marshal(result.KeyAssumptions)
	risksJSON, _ := json.Marshal(result.EarlyRisks)
	assumptionsStr := string(assumptionsJSON)
	risksStr := string(risksJSON)
	idea.KeyAssumptions = &assumptionsStr
	idea.EarlyRisks = &risksStr

	// Append this turn to history: user instruction + assistant summary
	history = append(history,
		Message{Role: "user", Content: instruction},
		Message{Role: "assistant", Content: result.Summary + "\n\nKonsep diperbarui: " + result.OneLineConcept},
	)
	// Ponytail: cap history to last 20 turns to bound token growth
	if len(history) > 20 {
		history = history[len(history)-20:]
	}
	historyJSON, _ := json.Marshal(history)
	historyStr := string(historyJSON)
	idea.RefineHistory = &historyStr
	idea.UpdatedAt = time.Now().UTC().Format(time.RFC3339)

	if err := s.ideaRepo.Update(idea); err != nil {
		return nil, nil, fmt.Errorf("update idea after refine: %w", err)
	}

	return idea, result, nil
}
