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
		 VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		m.ID, m.VentureID, m.Name, m.Description, m.Status, boolToInt(m.IsHero), m.ComplexityScore, m.ComplexityFactors, m.SortOrder, m.CreatedAt, m.UpdatedAt,
	)
	return err
}

func (r *MenuRepository) FindByVenture(ventureID string) ([]*domain.Menu, error) {
	rows, err := r.db.Query(
		`SELECT id, venture_id, name, description, status, is_hero, complexity_score, complexity_factors, sort_order, created_at, updated_at
		 FROM menus WHERE venture_id = ? ORDER BY sort_order`, ventureID,
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
		 FROM menus WHERE id = ?`, id,
	).Scan(&m.ID, &m.VentureID, &m.Name, &m.Description, &m.Status, &m.IsHero, &m.ComplexityScore, &m.ComplexityFactors, &m.SortOrder, &m.CreatedAt, &m.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	return m, err
}

func (r *MenuRepository) Update(m *domain.Menu) error {
	_, err := r.db.Exec(
		`UPDATE menus SET name=?, description=?, status=?, is_hero=?, complexity_score=?, complexity_factors=?, sort_order=?, updated_at=? WHERE id=?`,
		m.Name, m.Description, m.Status, boolToInt(m.IsHero), m.ComplexityScore, m.ComplexityFactors, m.SortOrder, m.UpdatedAt, m.ID,
	)
	return err
}

func (r *MenuRepository) Delete(id string) error {
	_, err := r.db.Exec(`DELETE FROM menus WHERE id=?`, id)
	return err
}

func (r *MenuRepository) CountActive(ventureID string) (int, error) {
	var count int
	err := r.db.QueryRow(`SELECT COUNT(*) FROM menus WHERE venture_id=? AND status='active'`, ventureID).Scan(&count)
	return count, err
}
