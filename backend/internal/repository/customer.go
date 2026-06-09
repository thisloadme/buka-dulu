package repository

import (
	"database/sql"
	"github.com/riyantobudi/bukadulu/internal/domain"
)

type CustomerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) Create(c *domain.CustomerSegment) error {
	_, err := r.db.Exec(
		`INSERT INTO customer_segments (id, venture_id, name, age_range, buy_context, budget_range, location, consumption_moment, is_too_general, is_locked, created_at, updated_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`,
		c.ID, c.VentureID, c.Name, c.AgeRange, c.BuyContext, c.BudgetRange, c.Location, c.ConsumptionMoment, c.IsTooGeneral, c.IsLocked, c.CreatedAt, c.UpdatedAt,
	)
	return err
}

func (r *CustomerRepository) FindByVenture(ventureID string) (*domain.CustomerSegment, error) {
	c := &domain.CustomerSegment{}
	err := r.db.QueryRow(
		`SELECT id, venture_id, name, age_range, buy_context, budget_range, location, consumption_moment, is_too_general, is_locked, created_at, updated_at
		 FROM customer_segments WHERE venture_id = $1 ORDER BY created_at DESC LIMIT 1`, ventureID,
	).Scan(&c.ID, &c.VentureID, &c.Name, &c.AgeRange, &c.BuyContext, &c.BudgetRange, &c.Location, &c.ConsumptionMoment, &c.IsTooGeneral, &c.IsLocked, &c.CreatedAt, &c.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	return c, err
}

func (r *CustomerRepository) Update(c *domain.CustomerSegment) error {
	_, err := r.db.Exec(
		`UPDATE customer_segments SET name=$1, age_range=$2, buy_context=$3, budget_range=$4, location=$5, consumption_moment=$6, is_too_general=$7, is_locked=$8, updated_at=CURRENT_TIMESTAMP WHERE id=$9`,
		c.Name, c.AgeRange, c.BuyContext, c.BudgetRange, c.Location, c.ConsumptionMoment, c.IsTooGeneral, c.IsLocked, c.ID,
	)
	return err
}

func (r *CustomerRepository) Lock(id string) error {
	_, err := r.db.Exec(`UPDATE customer_segments SET is_locked=TRUE, updated_at=CURRENT_TIMESTAMP WHERE id=$1`, id)
	return err
}
