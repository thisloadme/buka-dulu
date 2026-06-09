package domain

type Evidence struct {
	ID            string `db:"id" json:"id"`
	VentureID     string `db:"venture_id" json:"venture_id"`
	MissionID     string `db:"mission_id" json:"mission_id"`
	UploaderID    string `db:"uploader_id" json:"uploader_id"`
	EvidenceType  string `db:"evidence_type" json:"evidence_type"`
	StorageURL    string `db:"storage_url" json:"storage_url,omitempty"`
	TextContent   string `db:"text_content" json:"text_content,omitempty"`
	ThumbnailURL  string `db:"thumbnail_url" json:"thumbnail_url,omitempty"`
	FileSizeBytes int    `db:"file_size_bytes" json:"file_size_bytes,omitempty"`
	MimeType      string `db:"mime_type" json:"mime_type,omitempty"`
	SubmittedAt   string `db:"submitted_at" json:"submitted_at"`
}

type EvidenceReview struct {
	ID              string `db:"id" json:"id"`
	EvidenceID      string `db:"evidence_id" json:"evidence_id"`
	ReviewerType    string `db:"reviewer_type" json:"reviewer_type"`
	Verdict         string `db:"verdict" json:"verdict"`
	Score           float64 `db:"score" json:"score,omitempty"`
	Rationale       string `db:"rationale" json:"rationale"`
	NextAction      string `db:"next_action" json:"next_action,omitempty"`
	OverriddenBy    string `db:"overridden_by" json:"overridden_by,omitempty"`
	ProcessingTimeMs int   `db:"processing_time_ms" json:"processing_time_ms,omitempty"`
	CreatedAt       string `db:"created_at" json:"created_at"`
}

type EvidenceWithReview struct {
	Evidence
	Review *EvidenceReview `json:"review,omitempty"`
}

type UploadEvidenceRequest struct {
	MissionID    string `json:"mission_id"`
	EvidenceType string `json:"evidence_type"`
	TextContent  string `json:"text_content,omitempty"`
	LinkURL      string `json:"link_url,omitempty"`
}

type OverrideReviewRequest struct {
	Verdict   string `json:"verdict"`
	Rationale string `json:"rationale"`
}
