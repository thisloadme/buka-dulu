package domain

type Idea struct {
	ID               string  `db:"id" json:"id"`
	VentureID        string  `db:"venture_id" json:"venture_id"`
	RawInput         string  `db:"raw_input" json:"raw_input"`
	OneLineConcept   *string `db:"one_line_concept" json:"one_line_concept,omitempty"`
	TargetCustomer   *string `db:"target_customer" json:"target_customer,omitempty"`
	ValueProposition *string `db:"value_proposition" json:"value_proposition,omitempty"`
	KeyAssumptions   *string `db:"key_assumptions" json:"key_assumptions,omitempty"`
	EarlyRisks       *string `db:"early_risks" json:"early_risks,omitempty"`
	Version          int     `db:"version" json:"version"`
	IsLocked         bool    `db:"is_locked" json:"is_locked"`
	Status           string  `db:"status" json:"status"`
	AiRawInput       *string `db:"ai_raw_input" json:"-"`
	AiRawOutput      *string `db:"ai_raw_output" json:"-"`
	CreatedAt        string  `db:"created_at" json:"created_at"`
	UpdatedAt        string  `db:"updated_at" json:"updated_at"`
}

type StructuredConcept struct {
	OneLineConcept   string   `json:"one_line_concept"`
	TargetCustomer   string   `json:"target_customer"`
	ValueProposition string   `json:"value_proposition"`
	KeyAssumptions   []string `json:"key_assumptions"`
	EarlyRisks       []string `json:"early_risks"`
}

type UpdateIdeaRequest struct {
	OneLineConcept   *string `json:"one_line_concept,omitempty"`
	TargetCustomer   *string `json:"target_customer,omitempty"`
	ValueProposition *string `json:"value_proposition,omitempty"`
	KeyAssumptions   *string `json:"key_assumptions,omitempty"`
	EarlyRisks       *string `json:"early_risks,omitempty"`
}
