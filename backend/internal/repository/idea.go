package repository

import (
	"database/sql"
	"github.com/riyantobudi/bukadulu/internal/domain"
)

type IdeaRepository struct {
	db *sql.DB
}

func NewIdeaRepository(db *sql.DB) *IdeaRepository {
	return &IdeaRepository{db: db}
}

func (r *IdeaRepository) Create(idea *domain.Idea) error {
	_, err := r.db.Exec(
		`INSERT INTO ideas (id, venture_id, raw_input, version, is_locked, status, ai_raw_input, created_at, updated_at)
		 VALUES (?, ?, ?, ?, ?, ?, ?, datetime('now'), datetime('now'))`,
		idea.ID, idea.VentureID, idea.RawInput, idea.Version, boolToInt(idea.IsLocked), idea.Status, idea.AiRawInput,
	)
	return err
}

func (r *IdeaRepository) FindByVenture(ventureID string) (*domain.Idea, error) {
	idea := &domain.Idea{}
	var aiRawInput, aiRawOutput *string
	err := r.db.QueryRow(
		`SELECT id, venture_id, raw_input, one_line_concept, target_customer, value_proposition,
		        key_assumptions, early_risks, version, is_locked, status, ai_raw_input, ai_raw_output, created_at, updated_at
		 FROM ideas WHERE venture_id = ? ORDER BY version DESC LIMIT 1`, ventureID,
	).Scan(&idea.ID, &idea.VentureID, &idea.RawInput, &idea.OneLineConcept, &idea.TargetCustomer,
		&idea.ValueProposition, &idea.KeyAssumptions, &idea.EarlyRisks,
		&idea.Version, &idea.IsLocked, &idea.Status, &aiRawInput, &aiRawOutput, &idea.CreatedAt, &idea.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	if aiRawInput != nil {
		idea.AiRawInput = aiRawInput
	}
	if aiRawOutput != nil {
		idea.AiRawOutput = aiRawOutput
	}
	return idea, err
}

func (r *IdeaRepository) Update(idea *domain.Idea) error {
	_, err := r.db.Exec(
		`UPDATE ideas SET one_line_concept=?, target_customer=?, value_proposition=?,
		 key_assumptions=?, early_risks=?, is_locked=?, status=?, ai_raw_output=?,
		 updated_at=datetime('now') WHERE id=?`,
		idea.OneLineConcept, idea.TargetCustomer, idea.ValueProposition,
		idea.KeyAssumptions, idea.EarlyRisks, boolToInt(idea.IsLocked), idea.Status, idea.AiRawOutput, idea.ID,
	)
	return err
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
