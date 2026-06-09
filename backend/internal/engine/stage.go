package engine

import "github.com/riyantobudi/bukadulu/internal/domain"

// AllowedTransitions maps current stage → allowed next stages
var AllowedTransitions = map[domain.VentureStage][]domain.VentureStage{
	domain.StageDraft:             {domain.StageIdeaDefined},
	domain.StageIdeaDefined:       {domain.StageCustomerDefined},
	domain.StageCustomerDefined:   {domain.StageSKUFocused},
	domain.StageSKUFocused:        {domain.StageCostEvaluated},
	domain.StageCostEvaluated:     {domain.StageMissionActive},
	domain.StageMissionActive:     {domain.StageEvidenceSubmitted},
	domain.StageEvidenceSubmitted: {domain.StageEvidenceReviewed},
	domain.StageEvidenceReviewed:  {domain.StageMissionActive, domain.StageReadyToDecide},
	domain.StageReadyToDecide:     {domain.StageContinue, domain.StageRepeat, domain.StagePivot, domain.StageStop},
}

// StageOrder is the numeric order for comparison
var StageOrder = map[domain.VentureStage]int{
	domain.StageDraft:             0,
	domain.StageIdeaDefined:       1,
	domain.StageCustomerDefined:   2,
	domain.StageSKUFocused:        3,
	domain.StageCostEvaluated:     4,
	domain.StageMissionActive:     5,
	domain.StageEvidenceSubmitted: 6,
	domain.StageEvidenceReviewed:  7,
	domain.StageReadyToDecide:     8,
	domain.StageContinue:          9,
	domain.StageRepeat:            5, // goes back to mission_active level
	domain.StagePivot:             9,
	domain.StageStop:              9,
}

// GateCheck validates if a transition is allowed
func GateCheck(current, next domain.VentureStage) error {
	allowed, ok := AllowedTransitions[current]
	if !ok {
		return domain.AppErrStageGate
	}
	for _, a := range allowed {
		if a == next {
			return nil
		}
	}
	return domain.AppErrStageGate
}

// CanTransitionForward checks if moving forward is allowed
func CanTransitionForward(current, next domain.VentureStage) bool {
	currentOrder, ok1 := StageOrder[current]
	nextOrder, ok2 := StageOrder[next]
	if !ok1 || !ok2 {
		return false
	}
	return nextOrder > currentOrder
}

// IsTerminalStage returns true if the stage is an end state
func IsTerminalStage(stage domain.VentureStage) bool {
	return stage == domain.StageContinue ||
		stage == domain.StageStop ||
		stage == domain.StagePivot
}
