package engine

import "github.com/riyantobudi/bukadulu/internal/domain"

// ScoreComponents holds all scoring factors
type ScoreComponents struct {
	ClarityScore       float64
	FocusScore         float64
	EconomicsScore     float64
	ExecutionScore     float64
	EvidenceScore      float64
	MarketResponseScore float64
}

// Weights for each component
const (
	WeightClarity        = 0.10
	WeightFocus          = 0.10
	WeightEconomics      = 0.25
	WeightExecution      = 0.20
	WeightEvidence       = 0.25
	WeightMarketResponse = 0.10
)

// CalculateTotal returns weighted total score (0-100)
func CalculateTotal(comp ScoreComponents) float64 {
	total := comp.ClarityScore*WeightClarity +
		comp.FocusScore*WeightFocus +
		comp.EconomicsScore*WeightEconomics +
		comp.ExecutionScore*WeightExecution +
		comp.EvidenceScore*WeightEvidence +
		comp.MarketResponseScore*WeightMarketResponse
	return total
}

// MarginThresholds for economics scoring
const (
	MarginSehat     = 40.0
	MarginTipis     = 20.0
	MarginBerbahaya = 0.0
)

// CalculateClarityScore evaluates idea clarity
func CalculateClarityScore(idea *domain.Idea) float64 {
	if idea == nil || idea.OneLineConcept == nil {
		return 0
	}
	score := 50.0
	if len(*idea.OneLineConcept) > 30 {
		score += 20
	}
	if idea.TargetCustomer != nil && len(*idea.TargetCustomer) > 10 {
		score += 15
	}
	if idea.KeyAssumptions != nil && len(*idea.KeyAssumptions) > 10 {
		score += 15
	}
	if score > 100 {
		score = 100
	}
	return score
}

// CalculateFocusScore evaluates menu focus
func CalculateFocusScore(activeMenus int, hasHero bool) float64 {
	switch {
	case activeMenus == 0:
		return 0
	case activeMenus <= 2 && hasHero:
		return 100
	case activeMenus <= 2:
		return 80
	case activeMenus == 3 && hasHero:
		return 75
	case activeMenus == 3:
		return 60
	default:
		return 30
	}
}

// CalculateEconomicsScore from margin percentage
func CalculateEconomicsScore(marginPercent float64) float64 {
	switch {
	case marginPercent >= MarginSehat:
		return 100
	case marginPercent >= MarginTipis:
		return 50 + (marginPercent-MarginTipis)/(MarginSehat-MarginTipis)*50
	case marginPercent > 0:
		return (marginPercent / MarginTipis) * 50
	default:
		return 0
	}
}

// CalculateExecutionScore from mission completion rate
func CalculateExecutionScore(completed, total int) float64 {
	if total == 0 {
		return 0
	}
	rate := float64(completed) / float64(total)
	return rate * 100
}

// CalculateEvidenceScore from evidence verdicts
func CalculateEvidenceScore(validCount, weakCount, total int) float64 {
	if total == 0 {
		return 0
	}
	score := (float64(validCount)*100 + float64(weakCount)*50) / float64(total)
	if score > 100 {
		score = 100
	}
	return score
}

// CalculateMarketResponseScore estimates market response
func CalculateMarketResponseScore(validEvidenceCount int) float64 {
	switch {
	case validEvidenceCount >= 3:
		return 80
	case validEvidenceCount == 2:
		return 60
	case validEvidenceCount == 1:
		return 40
	default:
		return 0
	}
}

// DecisionThresholds
const (
	DecisionThresholdContinue = 70.0
	DecisionThresholdRepeat   = 40.0
	DecisionEvidenceMinimum   = 60.0
	DecisionEconomicsMinimum  = 40.0
)

// GenerateDecision applies the decision matrix
func GenerateDecision(totalScore, evidenceScore, economicsScore float64) domain.DecisionType {
	lowEvidence := evidenceScore < DecisionEvidenceMinimum
	lowEconomics := economicsScore < DecisionEconomicsMinimum

	switch {
	case totalScore >= DecisionThresholdContinue && !lowEvidence && !lowEconomics:
		return domain.DecisionContinue
	case totalScore >= DecisionThresholdRepeat && !lowEconomics:
		return domain.DecisionContinue // continue with warning
	case totalScore >= DecisionThresholdRepeat && lowEconomics && lowEvidence:
		return domain.DecisionStop
	case totalScore >= DecisionThresholdRepeat && lowEvidence:
		return domain.DecisionRepeat
	case totalScore >= 20 && lowEconomics:
		return domain.DecisionPivot
	default:
		return domain.DecisionStop
	}
}
