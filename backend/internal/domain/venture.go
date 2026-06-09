package domain

type VentureStage string

const (
	StageDraft             VentureStage = "draft"
	StageIdeaDefined       VentureStage = "idea_defined"
	StageCustomerDefined   VentureStage = "customer_defined"
	StageSKUFocused        VentureStage = "sku_focused"
	StageCostEvaluated     VentureStage = "cost_evaluated"
	StageMissionActive     VentureStage = "mission_active"
	StageEvidenceSubmitted VentureStage = "evidence_submitted"
	StageEvidenceReviewed  VentureStage = "evidence_reviewed"
	StageReadyToDecide     VentureStage = "ready_to_decide"
	StageContinue          VentureStage = "continue"
	StageRepeat            VentureStage = "repeat"
	StagePivot             VentureStage = "pivot"
	StageStop              VentureStage = "stop"
)

type Venture struct {
	ID             string       `db:"id" json:"id"`
	OwnerUserID    string       `db:"owner_user_id" json:"owner_user_id"`
	Name           string       `db:"name" json:"name"`
	Category       string       `db:"category" json:"category,omitempty"`
	Region         string       `db:"region" json:"region,omitempty"`
	Stage          VentureStage `db:"stage" json:"stage"`
	CurrentVersion int          `db:"current_version" json:"current_version"`
	CreatedAt      string       `db:"created_at" json:"created_at"`
	UpdatedAt      string       `db:"updated_at" json:"updated_at"`
}

type CreateVentureRequest struct {
	Name     string `json:"name"`
	Category string `json:"category,omitempty"`
	Region   string `json:"region,omitempty"`
}

type UpdateVentureRequest struct {
	Name     *string `json:"name,omitempty"`
	Category *string `json:"category,omitempty"`
	Region   *string `json:"region,omitempty"`
}
