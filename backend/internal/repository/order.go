package repository

import (
	"database/sql"
	"time"

	"github.com/riyantobudi/bukadulu/internal/domain"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(o *domain.Order) error {
	_, err := r.db.Exec(
		`INSERT INTO orders (id, user_id, venture_id, purpose, amount, total_amount, status,
		 qris_url, qris_image, signature, klikqris_order_id, expired_at, fulfilled, created_at, updated_at)
		 VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15)`,
		o.ID, o.UserID, o.VentureID, o.Purpose, o.Amount, o.TotalAmount, o.Status,
		o.QrisURL, o.QrisImage, o.Signature, o.KlikQrisOrderID, o.ExpiredAt, o.Fulfilled,
		time.Now().UTC().Format(time.RFC3339), time.Now().UTC().Format(time.RFC3339),
	)
	return err
}

func (r *OrderRepository) FindByID(id string) (*domain.Order, error) {
	o := &domain.Order{}
	err := r.db.QueryRow(
		`SELECT id, user_id, venture_id, purpose, amount, total_amount, status,
		        qris_url, qris_image, signature, klikqris_order_id, expired_at, paid_at, fulfilled,
		        created_at, updated_at
		 FROM orders WHERE id = $1`, id,
	).Scan(&o.ID, &o.UserID, &o.VentureID, &o.Purpose, &o.Amount, &o.TotalAmount, &o.Status,
		&o.QrisURL, &o.QrisImage, &o.Signature, &o.KlikQrisOrderID, &o.ExpiredAt, &o.PaidAt, &o.Fulfilled,
		&o.CreatedAt, &o.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	return o, err
}

func (r *OrderRepository) FindByKlikQrisOrderID(klikqrisID string) (*domain.Order, error) {
	o := &domain.Order{}
	err := r.db.QueryRow(
		`SELECT id, user_id, venture_id, purpose, amount, total_amount, status,
		        qris_url, qris_image, signature, klikqris_order_id, expired_at, paid_at, fulfilled,
		        created_at, updated_at
		 FROM orders WHERE klikqris_order_id = $1`, klikqrisID,
	).Scan(&o.ID, &o.UserID, &o.VentureID, &o.Purpose, &o.Amount, &o.TotalAmount, &o.Status,
		&o.QrisURL, &o.QrisImage, &o.Signature, &o.KlikQrisOrderID, &o.ExpiredAt, &o.PaidAt, &o.Fulfilled,
		&o.CreatedAt, &o.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	return o, err
}

// FindPendingPaidForVenture returns a paid, unfulfilled order for the venture (if any).
func (r *OrderRepository) FindPaidUnfulfilledForVenture(ventureID, purpose string) (*domain.Order, error) {
	o := &domain.Order{}
	err := r.db.QueryRow(
		`SELECT id, user_id, venture_id, purpose, amount, total_amount, status,
		        qris_url, qris_image, signature, klikqris_order_id, expired_at, paid_at, fulfilled,
		        created_at, updated_at
		 FROM orders WHERE venture_id = $1 AND purpose = $2 AND status = 'paid' AND fulfilled = FALSE
		 ORDER BY created_at DESC LIMIT 1`, ventureID, purpose,
	).Scan(&o.ID, &o.UserID, &o.VentureID, &o.Purpose, &o.Amount, &o.TotalAmount, &o.Status,
		&o.QrisURL, &o.QrisImage, &o.Signature, &o.KlikQrisOrderID, &o.ExpiredAt, &o.PaidAt, &o.Fulfilled,
		&o.CreatedAt, &o.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	return o, err
}

func (r *OrderRepository) MarkPaid(klikqrisID, paidAt string) error {
	_, err := r.db.Exec(
		`UPDATE orders SET status = 'paid', paid_at = $1, updated_at = $1 WHERE klikqris_order_id = $2 AND status != 'paid'`,
		paidAt, klikqrisID,
	)
	return err
}

func (r *OrderRepository) MarkExpired(klikqrisID string) error {
	_, err := r.db.Exec(
		`UPDATE orders SET status = 'expired', updated_at = CURRENT_TIMESTAMP WHERE klikqris_order_id = $1 AND status = 'pending'`,
		klikqrisID,
	)
	return err
}

func (r *OrderRepository) MarkFulfilled(id string) error {
	_, err := r.db.Exec(
		`UPDATE orders SET fulfilled = TRUE, updated_at = CURRENT_TIMESTAMP WHERE id = $1`, id,
	)
	return err
}
