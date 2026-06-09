package repository

import (
	"database/sql"
	"github.com/riyantobudi/bukadulu/internal/domain"
)

type EvidenceRepository struct {
	db *sql.DB
}

func NewEvidenceRepository(db *sql.DB) *EvidenceRepository {
	return &EvidenceRepository{db: db}
}

func (r *EvidenceRepository) Create(e *domain.Evidence) error {
	_, err := r.db.Exec(
		`INSERT INTO evidences (id, venture_id, mission_id, uploader_id, evidence_type, storage_url, text_content, file_size_bytes, mime_type, submitted_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		e.ID, e.VentureID, e.MissionID, e.UploaderID, e.EvidenceType, e.StorageURL, e.TextContent, e.FileSizeBytes, e.MimeType, e.SubmittedAt,
	)
	return err
}

func scanEvidence(scanner interface{ Scan(dest ...interface{}) error }) (*domain.Evidence, error) {
	e := &domain.Evidence{}
	var storageURL, textContent, thumbnailURL, mimeType sql.NullString
	var fileSize sql.NullInt64
	err := scanner.Scan(&e.ID, &e.VentureID, &e.MissionID, &e.UploaderID, &e.EvidenceType, &storageURL, &textContent, &thumbnailURL, &fileSize, &mimeType, &e.SubmittedAt)
	if err != nil {
		return nil, err
	}
	e.StorageURL = storageURL.String
	e.TextContent = textContent.String
	e.ThumbnailURL = thumbnailURL.String
	e.MimeType = mimeType.String
	if fileSize.Valid {
		e.FileSizeBytes = int(fileSize.Int64)
	}
	return e, nil
}

func (r *EvidenceRepository) FindByVenture(ventureID string) ([]*domain.Evidence, error) {
	rows, err := r.db.Query(
		`SELECT id, venture_id, mission_id, uploader_id, evidence_type, storage_url, text_content, thumbnail_url, file_size_bytes, mime_type, submitted_at
		 FROM evidences WHERE venture_id=$1 ORDER BY submitted_at DESC`, ventureID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*domain.Evidence
	for rows.Next() {
		e, err := scanEvidence(rows)
		if err != nil {
			return nil, err
		}
		items = append(items, e)
	}
	return items, nil
}

func (r *EvidenceRepository) FindByMission(missionID string) ([]*domain.Evidence, error) {
	rows, err := r.db.Query(
		`SELECT id, venture_id, mission_id, uploader_id, evidence_type, storage_url, text_content, thumbnail_url, file_size_bytes, mime_type, submitted_at
		 FROM evidences WHERE mission_id=$1 ORDER BY submitted_at DESC`, missionID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*domain.Evidence
	for rows.Next() {
		e, err := scanEvidence(rows)
		if err != nil {
			return nil, err
		}
		items = append(items, e)
	}
	return items, nil
}

func (r *EvidenceRepository) FindByID(id string) (*domain.Evidence, error) {
	row := r.db.QueryRow(
		`SELECT id, venture_id, mission_id, uploader_id, evidence_type, storage_url, text_content, thumbnail_url, file_size_bytes, mime_type, submitted_at
		 FROM evidences WHERE id=$1`, id,
	)
	e, err := scanEvidence(row)
	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	return e, err
}

// Evidence Review

func (r *EvidenceRepository) CreateReview(rv *domain.EvidenceReview) error {
	_, err := r.db.Exec(
		`INSERT INTO evidence_reviews (id, evidence_id, reviewer_type, verdict, score, rationale, next_action, processing_time_ms, created_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		rv.ID, rv.EvidenceID, rv.ReviewerType, rv.Verdict, rv.Score, rv.Rationale, rv.NextAction, rv.ProcessingTimeMs, rv.CreatedAt,
	)
	return err
}

func (r *EvidenceRepository) FindReviewByEvidence(evidenceID string) (*domain.EvidenceReview, error) {
	rv := &domain.EvidenceReview{}
	var nextAction, overriddenBy sql.NullString
	err := r.db.QueryRow(
		`SELECT id, evidence_id, reviewer_type, verdict, score, rationale, next_action, overridden_by, processing_time_ms, created_at
		 FROM evidence_reviews WHERE evidence_id=$1 ORDER BY created_at DESC LIMIT 1`, evidenceID,
	).Scan(&rv.ID, &rv.EvidenceID, &rv.ReviewerType, &rv.Verdict, &rv.Score, &rv.Rationale, &nextAction, &overriddenBy, &rv.ProcessingTimeMs, &rv.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	rv.NextAction = nextAction.String
	rv.OverriddenBy = overriddenBy.String
	return rv, err
}

func (r *EvidenceRepository) OverrideReview(id string, verdict, rationale, overriddenBy string) error {
	_, err := r.db.Exec(
		`UPDATE evidence_reviews SET verdict=$1, rationale=$2, overridden_by=$3, reviewer_type='human_override' WHERE id=$4`,
		verdict, rationale, overriddenBy, id,
	)
	return err
}
