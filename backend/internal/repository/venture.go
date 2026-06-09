package repository

import (
	"database/sql"
	"github.com/riyantobudi/bukadulu/internal/domain"
)

type VentureRepository struct {
	db *sql.DB
}

func NewVentureRepository(db *sql.DB) *VentureRepository {
	return &VentureRepository{db: db}
}

func (r *VentureRepository) Create(v *domain.Venture) error {
	_, err := r.db.Exec(
		`INSERT INTO ventures (id, owner_user_id, name, category, region, stage, current_version, created_at, updated_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`,
		v.ID, v.OwnerUserID, v.Name, v.Category, v.Region, v.Stage, v.CurrentVersion,
	)
	return err
}

func (r *VentureRepository) FindByID(id string) (*domain.Venture, error) {
	v := &domain.Venture{}
	err := r.db.QueryRow(
		`SELECT id, owner_user_id, name, category, region, stage, current_version, created_at, updated_at
		 FROM ventures WHERE id = $1`, id,
	).Scan(&v.ID, &v.OwnerUserID, &v.Name, &v.Category, &v.Region, &v.Stage, &v.CurrentVersion, &v.CreatedAt, &v.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	return v, err
}

func (r *VentureRepository) FindByOwner(ownerID string) ([]*domain.Venture, error) {
	rows, err := r.db.Query(
		`SELECT id, owner_user_id, name, category, region, stage, current_version, created_at, updated_at
		 FROM ventures WHERE owner_user_id = $1 ORDER BY updated_at DESC`, ownerID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ventures []*domain.Venture
	for rows.Next() {
		v := &domain.Venture{}
		if err := rows.Scan(&v.ID, &v.OwnerUserID, &v.Name, &v.Category, &v.Region, &v.Stage, &v.CurrentVersion, &v.CreatedAt, &v.UpdatedAt); err != nil {
			return nil, err
		}
		ventures = append(ventures, v)
	}
	return ventures, nil
}

func (r *VentureRepository) Update(v *domain.Venture) error {
	_, err := r.db.Exec(
		`UPDATE ventures SET name=$1, category=$2, region=$3, stage=$4, current_version=$5, updated_at=CURRENT_TIMESTAMP
		 WHERE id=$6`,
		v.Name, v.Category, v.Region, v.Stage, v.CurrentVersion, v.ID,
	)
	return err
}

func (r *VentureRepository) UpdateStage(id string, stage domain.VentureStage) error {
	_, err := r.db.Exec(
		`UPDATE ventures SET stage=$1, updated_at=CURRENT_TIMESTAMP WHERE id=$2`,
		stage, id,
	)
	return err
}
