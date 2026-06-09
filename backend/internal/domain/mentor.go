package domain

type MentorComment struct {
	ID         string `json:"id"`
	MentorID   string `json:"mentor_id"`
	VentureID  string `json:"venture_id"`
	MissionID  string `json:"mission_id,omitempty"`
	EvidenceID string `json:"evidence_id,omitempty"`
	Content    string `json:"content"`
	IsDeleted  bool   `json:"is_deleted"`
	CreatedAt  string `json:"created_at"`
}

type Mentor struct {
	ID             string `json:"id"`
	MentorID       string `json:"mentor_id"`
	FounderID      string `json:"founder_id"`
	VentureID      string `json:"venture_id"`
	Status         string `json:"status"`
	CreatedAt      string `json:"created_at"`
}
