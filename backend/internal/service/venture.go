package service

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/riyantobudi/bukadulu/internal/domain"
	"github.com/riyantobudi/bukadulu/internal/engine"
	"github.com/riyantobudi/bukadulu/internal/repository"
)

type VentureService struct {
	ventureRepo *repository.VentureRepository
}

func NewVentureService(ventureRepo *repository.VentureRepository) *VentureService {
	return &VentureService{ventureRepo: ventureRepo}
}

func (s *VentureService) Create(ownerID string, req *domain.CreateVentureRequest) (*domain.Venture, error) {
	now := time.Now().UTC().Format(time.RFC3339)
	venture := &domain.Venture{
		ID:             uuid.New().String(),
		OwnerUserID:    ownerID,
		Name:           req.Name,
		Category:       req.Category,
		Region:         req.Region,
		Stage:          domain.StageDraft,
		CurrentVersion: 1,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
	if err := s.ventureRepo.Create(venture); err != nil {
		return nil, fmt.Errorf("create venture: %w", err)
	}
	return venture, nil
}

func (s *VentureService) GetByID(id, userID string) (*domain.Venture, error) {
	venture, err := s.ventureRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if venture.OwnerUserID != userID {
		return nil, fmt.Errorf("forbidden: %w", domain.ErrForbidden)
	}
	return venture, nil
}

func (s *VentureService) ListByOwner(ownerID string) ([]*domain.Venture, error) {
	return s.ventureRepo.FindByOwner(ownerID)
}

func (s *VentureService) Update(id, userID string, req *domain.UpdateVentureRequest) (*domain.Venture, error) {
	venture, err := s.GetByID(id, userID)
	if err != nil {
		return nil, err
	}
	if req.Name != nil {
		venture.Name = *req.Name
	}
	if req.Category != nil {
		venture.Category = *req.Category
	}
	if req.Region != nil {
		venture.Region = *req.Region
	}
	venture.UpdatedAt = time.Now().UTC().Format(time.RFC3339)
	if err := s.ventureRepo.Update(venture); err != nil {
		return nil, fmt.Errorf("update venture: %w", err)
	}
	return venture, nil
}

func (s *VentureService) TransitionStage(id, userID string, nextStage domain.VentureStage) (*domain.Venture, error) {
	venture, err := s.GetByID(id, userID)
	if err != nil {
		return nil, err
	}
	if err := engine.GateCheck(venture.Stage, nextStage); err != nil {
		return nil, err
	}
	if err := s.ventureRepo.UpdateStage(id, nextStage); err != nil {
		return nil, fmt.Errorf("transition stage: %w", err)
	}
	venture.Stage = nextStage
	return venture, nil
}
