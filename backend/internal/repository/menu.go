package repository

import (
	"database/sql"
	"github.com/riyantobudi/bukadulu/internal/domain"
)

type MenuRepository struct {
	db *sql.DB
}

func NewMenuRepository(db *sql.DB) *MenuRepository {
	return &MenuRepository{db: db}
}

func (r *MenuRepository) Create(m *domain.Menu) error {
	_, err := r.db.Exec(
		`INSERT INTO menus (id, venture_id, name, description, status, is_hero, complexity_score, complexity_factors, sort_order, created_at, updated_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		m.ID, m.VentureID, m.Name, m.Description, m.Status, m.IsHero, m.ComplexityScore, m.ComplexityFactors, m.SortOrder, m.CreatedAt, m.UpdatedAt,
	)
	return err
}

func (r *MenuRepository) FindByVenture(ventureID string) ([]*domain.Menu, error) {
	rows, err := r.db.Query(
		`SELECT id, venture_id, name, description, status, is_hero, complexity_score, complexity_factors, sort_order, created_at, updated_at
		 FROM menus WHERE venture_id = $1 ORDER BY sort_order`, ventureID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var menus []*domain.Menu
	for rows.Next() {
		m := &domain.Menu{}
		if err := rows.Scan(&m.ID, &m.VentureID, &m.Name, &m.Description, &m.Status, &m.IsHero, &m.ComplexityScore, &m.ComplexityFactors, &m.SortOrder, &m.CreatedAt, &m.UpdatedAt); err != nil {
			return nil, err
		}
		menus = append(menus, m)
	}
	return menus, nil
}

func (r *MenuRepository) FindByID(id string) (*domain.Menu, error) {
	m := &domain.Menu{}
	err := r.db.QueryRow(
		`SELECT id, venture_id, name, description, status, is_hero, complexity_score, complexity_factors, sort_order, created_at, updated_at
		 FROM menus WHERE id = $1`, id,
	).Scan(&m.ID, &m.VentureID, &m.Name, &m.Description, &m.Status, &m.IsHero, &m.ComplexityScore, &m.ComplexityFactors, &m.SortOrder, &m.CreatedAt, &m.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	return m, err
}

func (r *MenuRepository) Update(m *domain.Menu) error {
	_, err := r.db.Exec(
		`UPDATE menus SET name=$1, description=$2, status=$3, is_hero=$4, complexity_score=$5, complexity_factors=$6, sort_order=$7, updated_at=CURRENT_TIMESTAMP WHERE id=$8`,
		m.Name, m.Description, m.Status, m.IsHero, m.ComplexityScore, m.ComplexityFactors, m.SortOrder, m.ID,
	)
	return err
}

func (r *MenuRepository) Delete(id string) error {
	_, err := r.db.Exec(`DELETE FROM menus WHERE id=$1`, id)
	return err
}

func (r *MenuRepository) CountActive(ventureID string) (int, error) {
	var count int
	err := r.db.QueryRow(`SELECT COUNT(*) FROM menus WHERE venture_id=$1 AND status='active'`, ventureID).Scan(&count)
	return count, err
}
