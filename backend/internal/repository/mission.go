package repository

import (
	"database/sql"
	"github.com/riyantobudi/bukadulu/internal/domain"
)

type MissionRepository struct {
	db *sql.DB
}

func NewMissionRepository(db *sql.DB) *MissionRepository {
	return &MissionRepository{db: db}
}

func (r *MissionRepository) Create(m *domain.Mission) error {
	_, err := r.db.Exec(
		`INSERT INTO missions (id, venture_id, title, description, mission_type, priority, status, due_at, evidence_required, estimated_minutes, created_by, sort_order, created_at, updated_at)
		 VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		m.ID, m.VentureID, m.Title, m.Description, m.MissionType, m.Priority, m.Status, nullIfEmpty(m.DueAt), m.EvidenceRequired, m.EstimatedMinutes, nullIfEmpty(m.CreatedBy), m.SortOrder, m.CreatedAt, m.UpdatedAt,
	)
	return err
}

func (r *MissionRepository) FindByVenture(ventureID string) ([]*domain.Mission, error) {
	rows, err := r.db.Query(
		`SELECT id, venture_id, title, description, mission_type, priority, status, due_at, evidence_required, estimated_minutes, created_by, sort_order, created_at, updated_at
		 FROM missions WHERE venture_id = ? ORDER BY sort_order, created_at`, ventureID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var missions []*domain.Mission
	for rows.Next() {
		m := &domain.Mission{}
		var dueAt, createdBy sql.NullString
		if err := rows.Scan(&m.ID, &m.VentureID, &m.Title, &m.Description, &m.MissionType, &m.Priority, &m.Status, &dueAt, &m.EvidenceRequired, &m.EstimatedMinutes, &createdBy, &m.SortOrder, &m.CreatedAt, &m.UpdatedAt); err != nil {
			return nil, err
		}
		m.DueAt = dueAt.String
		m.CreatedBy = createdBy.String
		missions = append(missions, m)
	}
	return missions, nil
}

func (r *MissionRepository) FindByID(id string) (*domain.Mission, error) {
	m := &domain.Mission{}
	var dueAt, createdBy sql.NullString
	err := r.db.QueryRow(
		`SELECT id, venture_id, title, description, mission_type, priority, status, due_at, evidence_required, estimated_minutes, created_by, sort_order, created_at, updated_at
		 FROM missions WHERE id = ?`, id,
	).Scan(&m.ID, &m.VentureID, &m.Title, &m.Description, &m.MissionType, &m.Priority, &m.Status, &dueAt, &m.EvidenceRequired, &m.EstimatedMinutes, &createdBy, &m.SortOrder, &m.CreatedAt, &m.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	m.DueAt = dueAt.String
	m.CreatedBy = createdBy.String
	return m, err
}

func (r *MissionRepository) Update(m *domain.Mission) error {
	_, err := r.db.Exec(
		`UPDATE missions SET title=?, description=?, mission_type=?, priority=?, status=?, due_at=?, evidence_required=?, estimated_minutes=?, sort_order=?, updated_at=? WHERE id=?`,
		m.Title, m.Description, m.MissionType, m.Priority, m.Status, nullIfEmpty(m.DueAt), m.EvidenceRequired, m.EstimatedMinutes, m.SortOrder, m.UpdatedAt, m.ID,
	)
	return err
}

func (r *MissionRepository) CountByStatus(ventureID, status string) (int, error) {
	var count int
	err := r.db.QueryRow(`SELECT COUNT(*) FROM missions WHERE venture_id=? AND status=?`, ventureID, status).Scan(&count)
	return count, err
}

func nullIfEmpty(s string) interface{} {
	if s == "" {
		return nil
	}
	return s
}
