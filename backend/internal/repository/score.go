package repository

import (
	"database/sql"
	"github.com/riyantobudi/bukadulu/internal/domain"
)

type ScoreRepository struct {
	db *sql.DB
}

func NewScoreRepository(db *sql.DB) *ScoreRepository {
	return &ScoreRepository{db: db}
}

func (r *ScoreRepository) Save(s *domain.Score) error {
	_, err := r.db.Exec(
		`INSERT INTO scores (id, venture_id, clarity_score, focus_score, economics_score, execution_score, evidence_score, market_response_score, total_score, is_final, created_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		s.ID, s.VentureID, s.ClarityScore, s.FocusScore, s.EconomicsScore, s.ExecutionScore, s.EvidenceScore, s.MarketResponseScore, s.TotalScore, s.IsFinal, s.CreatedAt,
	)
	return err
}

func (r *ScoreRepository) FindLatest(ventureID string) (*domain.Score, error) {
	s := &domain.Score{}
	err := r.db.QueryRow(
		`SELECT id, venture_id, clarity_score, focus_score, economics_score, execution_score, evidence_score, market_response_score, total_score, is_final, created_at
		 FROM scores WHERE venture_id=$1 ORDER BY created_at DESC LIMIT 1`, ventureID,
	).Scan(&s.ID, &s.VentureID, &s.ClarityScore, &s.FocusScore, &s.EconomicsScore, &s.ExecutionScore, &s.EvidenceScore, &s.MarketResponseScore, &s.TotalScore, &s.IsFinal, &s.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	return s, err
}

func (r *ScoreRepository) SaveDecision(d *domain.Decision) error {
	_, err := r.db.Exec(
		`INSERT INTO decisions (id, venture_id, decision, rationale, score_snapshot, triggered_by, created_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		d.ID, d.VentureID, string(d.Decision), d.Rationale, d.ScoreSnapshot, d.TriggeredBy, d.CreatedAt,
	)
	return err
}

func (r *ScoreRepository) FindDecision(ventureID string) (*domain.Decision, error) {
	d := &domain.Decision{}
	err := r.db.QueryRow(
		`SELECT id, venture_id, decision, rationale, score_snapshot, triggered_by, created_at
		 FROM decisions WHERE venture_id=$1 ORDER BY created_at DESC LIMIT 1`, ventureID,
	).Scan(&d.ID, &d.VentureID, &d.Decision, &d.Rationale, &d.ScoreSnapshot, &d.TriggeredBy, &d.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	return d, err
}
