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
		 VALUES ($1, $2, $3, $4, $5, $6, $7, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`,
		idea.ID, idea.VentureID, idea.RawInput, idea.Version, idea.IsLocked, idea.Status, idea.AiRawInput,
	)
	return err
}

func (r *IdeaRepository) FindByVenture(ventureID string) (*domain.Idea, error) {
	idea := &domain.Idea{}
	var aiRawInput, aiRawOutput *string
	err := r.db.QueryRow(
		`SELECT id, venture_id, raw_input, one_line_concept, target_customer, value_proposition,
		        key_assumptions, early_risks, version, is_locked, status, ai_raw_input, ai_raw_output, created_at, updated_at
		 FROM ideas WHERE venture_id = $1 ORDER BY version DESC LIMIT 1`, ventureID,
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
		`UPDATE ideas SET one_line_concept=$1, target_customer=$2, value_proposition=$3,
		 key_assumptions=$4, early_risks=$5, is_locked=$6, status=$7, ai_raw_output=$8,
		 updated_at=CURRENT_TIMESTAMP WHERE id=$9`,
		idea.OneLineConcept, idea.TargetCustomer, idea.ValueProposition,
		idea.KeyAssumptions, idea.EarlyRisks, idea.IsLocked, idea.Status, idea.AiRawOutput, idea.ID,
	)
	return err
}
