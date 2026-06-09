package domain

type Score struct {
	ID                 string  `db:"id" json:"id"`
	VentureID          string  `db:"venture_id" json:"venture_id"`
	ClarityScore       float64 `db:"clarity_score" json:"clarity_score"`
	FocusScore         float64 `db:"focus_score" json:"focus_score"`
	EconomicsScore     float64 `db:"economics_score" json:"economics_score"`
	ExecutionScore     float64 `db:"execution_score" json:"execution_score"`
	EvidenceScore      float64 `db:"evidence_score" json:"evidence_score"`
	MarketResponseScore float64 `db:"market_response_score" json:"market_response_score"`
	TotalScore         float64 `db:"total_score" json:"total_score"`
	IsFinal            bool    `db:"is_final" json:"is_final"`
	CreatedAt          string  `db:"created_at" json:"created_at"`
}

type DecisionType string

const (
	DecisionContinue DecisionType = "continue"
	DecisionRepeat   DecisionType = "repeat"
	DecisionPivot    DecisionType = "pivot"
	DecisionStop     DecisionType = "stop"
)

type Decision struct {
	ID            string       `db:"id" json:"id"`
	VentureID     string       `db:"venture_id" json:"venture_id"`
	Decision      DecisionType `db:"decision" json:"decision"`
	Rationale     string       `db:"rationale" json:"rationale"`
	ScoreSnapshot string       `db:"score_snapshot" json:"score_snapshot"`
	TriggeredBy   string       `db:"triggered_by" json:"triggered_by"`
	CreatedAt     string       `db:"created_at" json:"created_at"`
}
