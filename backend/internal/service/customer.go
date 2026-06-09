package service

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/riyantobudi/bukadulu/internal/domain"
	"github.com/riyantobudi/bukadulu/internal/repository"
)

type CustomerService struct {
	customerRepo *repository.CustomerRepository
	ventureSvc   *VentureService
}

func NewCustomerService(customerRepo *repository.CustomerRepository, ventureSvc *VentureService) *CustomerService {
	return &CustomerService{customerRepo: customerRepo, ventureSvc: ventureSvc}
}

func (s *CustomerService) Create(ventureID, userID string, req *domain.CreateCustomerSegmentRequest) (*domain.CustomerSegment, error) {
	_, err := s.ventureSvc.GetByID(ventureID, userID)
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC().Format(time.RFC3339)
	seg := &domain.CustomerSegment{
		ID:                uuid.New().String(),
		VentureID:         ventureID,
		Name:              req.Name,
		AgeRange:          req.AgeRange,
		BuyContext:        req.BuyContext,
		BudgetRange:       req.BudgetRange,
		Location:          req.Location,
		ConsumptionMoment: req.ConsumptionMoment,
		IsTooGeneral:      isSegmentTooGeneral(req),
		CreatedAt:         now,
		UpdatedAt:         now,
	}
	if err := s.customerRepo.Create(seg); err != nil {
		return nil, fmt.Errorf("create segment: %w", err)
	}
	return seg, nil
}

func (s *CustomerService) Get(ventureID, userID string) (*domain.CustomerSegment, error) {
	_, err := s.ventureSvc.GetByID(ventureID, userID)
	if err != nil {
		return nil, err
	}
	return s.customerRepo.FindByVenture(ventureID)
}

func (s *CustomerService) Confirm(ventureID, userID string) (*domain.CustomerSegment, error) {
	seg, err := s.Get(ventureID, userID)
	if err != nil {
		return nil, err
	}
	if err := s.customerRepo.Lock(seg.ID); err != nil {
		return nil, fmt.Errorf("lock segment: %w", err)
	}
	seg.IsLocked = true

	if _, err := s.ventureSvc.TransitionStage(ventureID, userID, domain.StageCustomerDefined); err != nil {
		return nil, err
	}
	return seg, nil
}

func isSegmentTooGeneral(req *domain.CreateCustomerSegmentRequest) bool {
	// If name or age_range mentions "semua", "all", or is empty
	if req.Name == "" || req.AgeRange == "" {
		return true
	}
	generalKeywords := []string{"semua", "all", "everyone", "semua orang", "general"}
	for _, kw := range generalKeywords {
		if req.Name == kw || req.AgeRange == kw {
			return true
		}
	}
	return false
}
