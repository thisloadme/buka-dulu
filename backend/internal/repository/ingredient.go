package repository

import (
	"database/sql"
	"github.com/riyantobudi/bukadulu/internal/domain"
)

type IngredientRepository struct {
	db *sql.DB
}

func NewIngredientRepository(db *sql.DB) *IngredientRepository {
	return &IngredientRepository{db: db}
}

func (r *IngredientRepository) Create(i *domain.Ingredient) error {
	_, err := r.db.Exec(
		`INSERT INTO ingredients (id, venture_id, menu_id, name, unit, quantity, unit_price, created_at, updated_at)
		 VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		i.ID, i.VentureID, i.MenuID, i.Name, i.Unit, i.Quantity, i.UnitPrice, i.CreatedAt, i.UpdatedAt,
	)
	return err
}

func (r *IngredientRepository) FindByMenu(menuID string) ([]*domain.Ingredient, error) {
	rows, err := r.db.Query(
		`SELECT id, venture_id, menu_id, name, unit, quantity, unit_price, created_at, updated_at
		 FROM ingredients WHERE menu_id=? ORDER BY created_at`, menuID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*domain.Ingredient
	for rows.Next() {
		i := &domain.Ingredient{}
		if err := rows.Scan(&i.ID, &i.VentureID, &i.MenuID, &i.Name, &i.Unit, &i.Quantity, &i.UnitPrice, &i.CreatedAt, &i.UpdatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	return items, nil
}

func (r *IngredientRepository) FindByIDs(ids []string) ([]*domain.Ingredient, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	q := `SELECT id, venture_id, menu_id, name, unit, quantity, unit_price, created_at, updated_at FROM ingredients WHERE id IN (`
	for i := range ids {
		if i > 0 {
			q += ","
		}
		q += "?"
	}
	q += ")"

	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id
	}

	rows, err := r.db.Query(q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*domain.Ingredient
	for rows.Next() {
		it := &domain.Ingredient{}
		if err := rows.Scan(&it.ID, &it.VentureID, &it.MenuID, &it.Name, &it.Unit, &it.Quantity, &it.UnitPrice, &it.CreatedAt, &it.UpdatedAt); err != nil {
			return nil, err
		}
		items = append(items, it)
	}
	return items, nil
}

func (r *IngredientRepository) Delete(id string) error {
	_, err := r.db.Exec(`DELETE FROM ingredients WHERE id=?`, id)
	return err
}

// Packaging

func (r *IngredientRepository) CreatePackaging(p *domain.PackagingCost) error {
	_, err := r.db.Exec(
		`INSERT INTO packaging_costs (id, venture_id, menu_id, name, unit_price, quantity_per_pcs, created_at)
		 VALUES (?, ?, ?, ?, ?, ?, ?)`,
		p.ID, p.VentureID, p.MenuID, p.Name, p.UnitPrice, p.QuantityPerPcs, p.CreatedAt,
	)
	return err
}

func (r *IngredientRepository) FindPackagingByMenu(menuID string) ([]*domain.PackagingCost, error) {
	rows, err := r.db.Query(
		`SELECT id, venture_id, menu_id, name, unit_price, quantity_per_pcs, created_at
		 FROM packaging_costs WHERE menu_id=?`, menuID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*domain.PackagingCost
	for rows.Next() {
		p := &domain.PackagingCost{}
		if err := rows.Scan(&p.ID, &p.VentureID, &p.MenuID, &p.Name, &p.UnitPrice, &p.QuantityPerPcs, &p.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, p)
	}
	return items, nil
}

// Cost Summary

func (r *IngredientRepository) SaveCostSummary(cs *domain.CostSummary) error {
	_, err := r.db.Exec(
		`INSERT INTO cost_summaries (id, venture_id, menu_id, hpp_per_porsi, suggested_price, target_margin, gross_margin, margin_status, break_even_unit, labor_per_unit, overhead_per_unit, is_locked, created_at, updated_at)
		 VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		 ON CONFLICT(venture_id, menu_id) DO UPDATE SET
		 hpp_per_porsi=excluded.hpp_per_porsi, suggested_price=excluded.suggested_price, target_margin=excluded.target_margin,
		 gross_margin=excluded.gross_margin, margin_status=excluded.margin_status, break_even_unit=excluded.break_even_unit,
		 labor_per_unit=excluded.labor_per_unit, overhead_per_unit=excluded.overhead_per_unit, updated_at=excluded.updated_at`,
		cs.ID, cs.VentureID, cs.MenuID, cs.HppPerPorsi, cs.SuggestedPrice, cs.TargetMargin,
		cs.GrossMargin, cs.MarginStatus, cs.BreakEvenUnit, cs.LaborPerUnit, cs.OverheadPerUnit,
		boolToInt(cs.IsLocked), cs.CreatedAt, cs.UpdatedAt,
	)
	return err
}

func (r *IngredientRepository) FindCostSummary(ventureID, menuID string) (*domain.CostSummary, error) {
	cs := &domain.CostSummary{}
	err := r.db.QueryRow(
		`SELECT id, venture_id, menu_id, hpp_per_porsi, suggested_price, target_margin, gross_margin, margin_status, break_even_unit, labor_per_unit, overhead_per_unit, is_locked, created_at, updated_at
		 FROM cost_summaries WHERE venture_id=? AND menu_id=?`, ventureID, menuID,
	).Scan(&cs.ID, &cs.VentureID, &cs.MenuID, &cs.HppPerPorsi, &cs.SuggestedPrice, &cs.TargetMargin, &cs.GrossMargin, &cs.MarginStatus, &cs.BreakEvenUnit, &cs.LaborPerUnit, &cs.OverheadPerUnit, &cs.IsLocked, &cs.CreatedAt, &cs.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	return cs, err
}

func (r *IngredientRepository) FindAllCostSummaries(ventureID string) ([]*domain.CostSummary, error) {
	rows, err := r.db.Query(
		`SELECT id, venture_id, menu_id, hpp_per_porsi, suggested_price, target_margin, gross_margin, margin_status, break_even_unit, labor_per_unit, overhead_per_unit, is_locked, created_at, updated_at
		 FROM cost_summaries WHERE venture_id=?`, ventureID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*domain.CostSummary
	for rows.Next() {
		cs := &domain.CostSummary{}
		if err := rows.Scan(&cs.ID, &cs.VentureID, &cs.MenuID, &cs.HppPerPorsi, &cs.SuggestedPrice, &cs.TargetMargin, &cs.GrossMargin, &cs.MarginStatus, &cs.BreakEvenUnit, &cs.LaborPerUnit, &cs.OverheadPerUnit, &cs.IsLocked, &cs.CreatedAt, &cs.UpdatedAt); err != nil {
			return nil, err
		}
		items = append(items, cs)
	}
	return items, nil
}

func (r *IngredientRepository) LockCost(ventureID string) error {
	_, err := r.db.Exec(`UPDATE cost_summaries SET is_locked=1, updated_at=datetime('now') WHERE venture_id=?`, ventureID)
	return err
}
