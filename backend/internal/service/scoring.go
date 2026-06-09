package service

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/riyantobudi/bukadulu/internal/domain"
	"github.com/riyantobudi/bukadulu/internal/engine"
	"github.com/riyantobudi/bukadulu/internal/repository"
)

type ScoringService struct {
	scoreRepo      *repository.ScoreRepository
	ventureRepo    *repository.VentureRepository
	ideaRepo       *repository.IdeaRepository
	menuRepo       *repository.MenuRepository
	ingredientRepo *repository.IngredientRepository
	missionRepo    *repository.MissionRepository
	evidenceRepo   *repository.EvidenceRepository
	ventureSvc     *VentureService
}

func NewScoringService(
	scoreRepo *repository.ScoreRepository,
	ventureRepo *repository.VentureRepository,
	ideaRepo *repository.IdeaRepository,
	menuRepo *repository.MenuRepository,
	ingredientRepo *repository.IngredientRepository,
	missionRepo *repository.MissionRepository,
	evidenceRepo *repository.EvidenceRepository,
	ventureSvc *VentureService,
) *ScoringService {
	return &ScoringService{
		scoreRepo: scoreRepo, ventureRepo: ventureRepo, ideaRepo: ideaRepo,
		menuRepo: menuRepo, ingredientRepo: ingredientRepo,
		missionRepo: missionRepo, evidenceRepo: evidenceRepo, ventureSvc: ventureSvc,
	}
}

func (s *ScoringService) Calculate(ventureID, userID string) (*domain.Score, error) {
	_, err := s.ventureSvc.GetByID(ventureID, userID)
	if err != nil {
		return nil, err
	}

	idea, _ := s.ideaRepo.FindByVenture(ventureID)
	clarity := engine.CalculateClarityScore(idea)

	menus, _ := s.menuRepo.FindByVenture(ventureID)
	activeCount := 0
	hasHero := false
	for _, m := range menus {
		if m.Status == "active" {
			activeCount++
			if m.IsHero {
				hasHero = true
			}
		}
	}
	focus := engine.CalculateFocusScore(activeCount, hasHero)

	summaries, _ := s.ingredientRepo.FindAllCostSummaries(ventureID)
	economics := 0.0
	if len(summaries) > 0 {
		economics = engine.CalculateEconomicsScore(summaries[0].GrossMargin)
	}

	missions, _ := s.missionRepo.FindByVenture(ventureID)
	completed, total := 0, 0
	for _, m := range missions {
		total++
		if m.Status == "completed" {
			completed++
		}
	}
	execution := engine.CalculateExecutionScore(completed, total)

	evidenceList, _ := s.evidenceRepo.FindByVenture(ventureID)
	validCount, weakCount := 0, 0
	for _, e := range evidenceList {
		review, rerr := s.evidenceRepo.FindReviewByEvidence(e.ID)
		if rerr == nil {
			switch review.Verdict {
			case "valid":
				validCount++
			case "weak":
				weakCount++
			}
		}
	}
	evScore := engine.CalculateEvidenceScore(validCount, weakCount, len(evidenceList))
	marketResp := engine.CalculateMarketResponseScore(validCount)

	comp := engine.ScoreComponents{
		ClarityScore: clarity, FocusScore: focus, EconomicsScore: economics,
		ExecutionScore: execution, EvidenceScore: evScore, MarketResponseScore: marketResp,
	}
	totalScore := engine.CalculateTotal(comp)

	now := time.Now().UTC().Format(time.RFC3339)
	score := &domain.Score{
		ID: uuid.New().String(), VentureID: ventureID,
		ClarityScore: clarity, FocusScore: focus, EconomicsScore: economics,
		ExecutionScore: execution, EvidenceScore: evScore,
		MarketResponseScore: marketResp, TotalScore: totalScore,
		CreatedAt: now,
	}
	if err := s.scoreRepo.Save(score); err != nil {
		return nil, err
	}
	return score, nil
}

func (s *ScoringService) GetLatest(ventureID, userID string) (*domain.Score, error) {
	_, err := s.ventureSvc.GetByID(ventureID, userID)
	if err != nil {
		return nil, err
	}
	return s.scoreRepo.FindLatest(ventureID)
}

func (s *ScoringService) GenerateDecision(ventureID, userID string) (*domain.Decision, error) {
	score, err := s.Calculate(ventureID, userID)
	if err != nil {
		return nil, err
	}

	decision := engine.GenerateDecision(score.TotalScore, score.EvidenceScore, score.EconomicsScore)

	var rationale string
	switch decision {
	case domain.DecisionContinue:
		rationale = "Ide kamu layak lanjut ke fase uji jual! Bukti cukup, margin sehat, dan target pasar jelas."
	case domain.DecisionRepeat:
		rationale = "Evidence masih kurang kuat. Ulangi misi dengan bukti yang lebih meyakinkan."
	case domain.DecisionPivot:
		rationale = "Margin kurang sehat atau evidence tidak cukup. Pertimbangkan pivot target pasar atau produk."
	case domain.DecisionStop:
		rationale = "Berdasarkan evidence yang ada, ide ini kurang layak dilanjutkan dalam bentuk sekarang. Ini bukan kegagalan."
	}

	snapshot := map[string]interface{}{
		"total": score.TotalScore, "clarity": score.ClarityScore, "focus": score.FocusScore,
		"economics": score.EconomicsScore, "execution": score.ExecutionScore,
		"evidence": score.EvidenceScore, "market_response": score.MarketResponseScore,
	}
	snapshotJSON, _ := json.Marshal(snapshot)

	now := time.Now().UTC().Format(time.RFC3339)
	d := &domain.Decision{
		ID: uuid.New().String(), VentureID: ventureID,
		Decision: decision, Rationale: rationale,
		ScoreSnapshot: string(snapshotJSON), TriggeredBy: "system", CreatedAt: now,
	}
	if err := s.scoreRepo.SaveDecision(d); err != nil {
		return nil, err
	}

	var stage domain.VentureStage
	switch decision {
	case domain.DecisionContinue:
		stage = domain.StageContinue
	case domain.DecisionRepeat:
		stage = domain.StageRepeat
	case domain.DecisionPivot:
		stage = domain.StagePivot
	case domain.DecisionStop:
		stage = domain.StageStop
	}
	s.ventureRepo.UpdateStage(ventureID, stage)

	return d, nil
}

func (s *ScoringService) GetDecision(ventureID, userID string) (*domain.Decision, error) {
	_, err := s.ventureSvc.GetByID(ventureID, userID)
	if err != nil {
		return nil, err
	}
	return s.scoreRepo.FindDecision(ventureID)
}
