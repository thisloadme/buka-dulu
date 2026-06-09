package service

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/riyantobudi/bukadulu/internal/domain"
	"github.com/riyantobudi/bukadulu/internal/repository"
)

type MenuService struct {
	menuRepo  *repository.MenuRepository
	llmSvc    *LLMService
	ventureSvc *VentureService
}

func NewMenuService(menuRepo *repository.MenuRepository, llmSvc *LLMService, ventureSvc *VentureService) *MenuService {
	return &MenuService{menuRepo: menuRepo, llmSvc: llmSvc, ventureSvc: ventureSvc}
}

func (s *MenuService) Create(ventureID, userID string, req *domain.CreateMenuRequest) (*domain.Menu, error) {
	_, err := s.ventureSvc.GetByID(ventureID, userID)
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC().Format(time.RFC3339)
	menu := &domain.Menu{
		ID:        uuid.New().String(),
		VentureID: ventureID,
		Name:      req.Name,
		Description: req.Description,
		Status:    "candidate",
		SortOrder: 0,
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := s.menuRepo.Create(menu); err != nil {
		return nil, fmt.Errorf("create menu: %w", err)
	}
	return menu, nil
}

func (s *MenuService) List(ventureID, userID string) ([]*domain.Menu, error) {
	_, err := s.ventureSvc.GetByID(ventureID, userID)
	if err != nil {
		return nil, err
	}
	return s.menuRepo.FindByVenture(ventureID)
}

func (s *MenuService) Update(menuID, userID string, req *domain.UpdateMenuRequest) (*domain.Menu, error) {
	menu, err := s.menuRepo.FindByID(menuID)
	if err != nil {
		return nil, err
	}

	// Verify ownership
	if _, err := s.ventureSvc.GetByID(menu.VentureID, userID); err != nil {
		return nil, err
	}

	if req.Name != nil {
		menu.Name = *req.Name
	}
	if req.Description != nil {
		menu.Description = *req.Description
	}
	if req.Status != nil {
		// Enforce max 3 active SKU
		if *req.Status == "active" {
			count, err := s.menuRepo.CountActive(menu.VentureID)
			if err != nil {
				return nil, err
			}
			if count >= 3 {
				return nil, fmt.Errorf("maximum 3 active SKU: %w", domain.ErrInvalidInput)
			}
		}
		menu.Status = *req.Status
	}
	if req.IsHero != nil {
		menu.IsHero = *req.IsHero
	}
	menu.UpdatedAt = time.Now().UTC().Format(time.RFC3339)

	if err := s.menuRepo.Update(menu); err != nil {
		return nil, fmt.Errorf("update menu: %w", err)
	}
	return menu, nil
}

func (s *MenuService) Delete(menuID, userID string) error {
	menu, err := s.menuRepo.FindByID(menuID)
	if err != nil {
		return err
	}
	if _, err := s.ventureSvc.GetByID(menu.VentureID, userID); err != nil {
		return err
	}
	return s.menuRepo.Delete(menuID)
}

func (s *MenuService) Focus(ventureID, userID string) error {
	menus, err := s.List(ventureID, userID)
	if err != nil {
		return err
	}

	activeCount := 0
	for _, m := range menus {
		if m.Status == "active" {
			activeCount++
		}
	}
	if activeCount < 1 {
		return fmt.Errorf("at least 1 active SKU required: %w", domain.ErrStageGate)
	}
	if activeCount > 3 {
		return fmt.Errorf("maximum 3 active SKU: %w", domain.ErrInvalidInput)
	}

	// Defer all non-active menus
	for _, m := range menus {
		if m.Status == "candidate" {
			m.Status = "deferred"
			m.UpdatedAt = time.Now().UTC().Format(time.RFC3339)
			s.menuRepo.Update(m)
		}
	}

	if _, err := s.ventureSvc.TransitionStage(ventureID, userID, domain.StageSKUFocused); err != nil {
		return err
	}
	return nil
}
