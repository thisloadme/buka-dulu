package service

import (
	"fmt"
	"math"
	"time"

	"github.com/google/uuid"
	"github.com/riyantobudi/bukadulu/internal/domain"
	"github.com/riyantobudi/bukadulu/internal/repository"
)

type CostService struct {
	ingredientRepo *repository.IngredientRepository
	ventureSvc     *VentureService
}

func NewCostService(ingredientRepo *repository.IngredientRepository, ventureSvc *VentureService) *CostService {
	return &CostService{ingredientRepo: ingredientRepo, ventureSvc: ventureSvc}
}

func (s *CostService) AddIngredient(ventureID, userID string, req *domain.CreateIngredientRequest) (*domain.Ingredient, error) {
	_, err := s.ventureSvc.GetByID(ventureID, userID)
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC().Format(time.RFC3339)
	ing := &domain.Ingredient{
		ID:        uuid.New().String(),
		VentureID: ventureID,
		MenuID:    req.MenuID,
		Name:      req.Name,
		Unit:      req.Unit,
		Quantity:  req.Quantity,
		UnitPrice: req.UnitPrice,
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := s.ingredientRepo.Create(ing); err != nil {
		return nil, fmt.Errorf("create ingredient: %w", err)
	}
	return ing, nil
}

func (s *CostService) GetIngredients(ventureID, menuID string) ([]*domain.Ingredient, error) {
	return s.ingredientRepo.FindByMenu(menuID)
}

func (s *CostService) DeleteIngredient(ingredientID string) error {
	return s.ingredientRepo.Delete(ingredientID)
}

func (s *CostService) AddPackaging(ventureID, userID string, req *domain.CreatePackagingRequest) (*domain.PackagingCost, error) {
	_, err := s.ventureSvc.GetByID(ventureID, userID)
	if err != nil {
		return nil, err
	}

	p := &domain.PackagingCost{
		ID:              uuid.New().String(),
		VentureID:       ventureID,
		MenuID:          req.MenuID,
		Name:            req.Name,
		UnitPrice:       req.UnitPrice,
		QuantityPerPcs:  req.QuantityPerPcs,
		CreatedAt:       time.Now().UTC().Format(time.RFC3339),
	}
	if err := s.ingredientRepo.CreatePackaging(p); err != nil {
		return nil, fmt.Errorf("create packaging: %w", err)
	}
	return p, nil
}

func (s *CostService) GetPackaging(menuID string) ([]*domain.PackagingCost, error) {
	return s.ingredientRepo.FindPackagingByMenu(menuID)
}

func (s *CostService) Calculate(ventureID, userID, menuID string, cfg *domain.CostConfigRequest) (*domain.CostSummary, error) {
	_, err := s.ventureSvc.GetByID(ventureID, userID)
	if err != nil {
		return nil, err
	}

	// Get all ingredients for this menu
	ingredients, err := s.ingredientRepo.FindByMenu(menuID)
	if err != nil {
		return nil, err
	}

	// Get packaging costs
	packaging, err := s.ingredientRepo.FindPackagingByMenu(menuID)
	if err != nil {
		return nil, err
	}

	// Calculate total ingredient cost per portion
	var totalIngredients float64
	for _, ing := range ingredients {
		totalIngredients += ing.Quantity * ing.UnitPrice
	}

	// Calculate total packaging cost per portion
	var totalPackaging float64
	for _, p := range packaging {
		totalPackaging += p.UnitPrice * p.QuantityPerPcs
	}

	// Labor and overhead
	labor := 0.0
	overhead := 0.0
	if cfg != nil {
		if cfg.LaborPerUnit != nil {
			labor = *cfg.LaborPerUnit
		}
		if cfg.OverheadPerUnit != nil {
			overhead = *cfg.OverheadPerUnit
		}
	}

	// HPP = ingredients + packaging + labor + overhead
	hpp := totalIngredients + totalPackaging + labor + overhead

	// Target margin (default 40%)
	targetMargin := 40.0
	if cfg != nil && cfg.TargetMargin != nil {
		targetMargin = *cfg.TargetMargin
	}

	// Suggested price: HPP / (1 - margin/100)
	suggestedPrice := math.Ceil(hpp / (1 - targetMargin/100) / 100) * 100
	if suggestedPrice < hpp*1.2 {
		suggestedPrice = math.Ceil(hpp*1.4/100) * 100
	}

	// Gross margin
	grossMargin := ((suggestedPrice - hpp) / suggestedPrice) * 100

	// Margin status
	marginStatus := "berbahaya"
	if grossMargin >= 40 {
		marginStatus = "sehat"
	} else if grossMargin >= 20 {
		marginStatus = "tipis"
	}

	// Break-even: assume monthly overhead estimate
	monthlyOverhead := 500000.0 // Rp 500k default monthly overhead
	breakEven := int(math.Ceil(monthlyOverhead / suggestedPrice))
	if breakEven < 1 {
		breakEven = 1
	}

	now := time.Now().UTC().Format(time.RFC3339)
	cs := &domain.CostSummary{
		ID:             uuid.New().String(),
		VentureID:      ventureID,
		MenuID:         menuID,
		HppPerPorsi:    math.Round(hpp*100) / 100,
		SuggestedPrice: suggestedPrice,
		TargetMargin:   targetMargin,
		GrossMargin:    math.Round(grossMargin*100) / 100,
		MarginStatus:   marginStatus,
		BreakEvenUnit:  breakEven,
		LaborPerUnit:   labor,
		OverheadPerUnit: overhead,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	if err := s.ingredientRepo.SaveCostSummary(cs); err != nil {
		return nil, fmt.Errorf("save cost summary: %w", err)
	}

	return cs, nil
}

func (s *CostService) GetSummary(ventureID, menuID string) (*domain.CostSummary, error) {
	return s.ingredientRepo.FindCostSummary(ventureID, menuID)
}

func (s *CostService) GetAllSummaries(ventureID string) ([]*domain.CostSummary, error) {
	return s.ingredientRepo.FindAllCostSummaries(ventureID)
}

func (s *CostService) Confirm(ventureID, userID string) error {
	if _, err := s.ventureSvc.TransitionStage(ventureID, userID, domain.StageCostEvaluated); err != nil {
		return err
	}
	return s.ingredientRepo.LockCost(ventureID)
}
