package domain

type Mission struct {
	ID               string `db:"id" json:"id"`
	VentureID        string `db:"venture_id" json:"venture_id"`
	Title            string `db:"title" json:"title"`
	Description      string `db:"description" json:"description"`
	MissionType      string `db:"mission_type" json:"mission_type"`
	Priority         string `db:"priority" json:"priority"`
	Status           string `db:"status" json:"status"`
	DueAt            string `db:"due_at" json:"due_at,omitempty"`
	EvidenceRequired string `db:"evidence_required" json:"evidence_required,omitempty"`
	EstimatedMinutes int    `db:"estimated_minutes" json:"estimated_minutes,omitempty"`
	CreatedBy        string `db:"created_by" json:"created_by,omitempty"`
	SortOrder        int    `db:"sort_order" json:"sort_order"`
	CreatedAt        string `db:"created_at" json:"created_at"`
	UpdatedAt        string `db:"updated_at" json:"updated_at"`
}

type CreateMissionRequest struct {
	Title            string `json:"title"`
	Description      string `json:"description"`
	MissionType      string `json:"mission_type"`
	Priority         string `json:"priority,omitempty"`
	DueAt            string `json:"due_at,omitempty"`
	EvidenceRequired string `json:"evidence_required,omitempty"`
	EstimatedMinutes int    `json:"estimated_minutes,omitempty"`
}
