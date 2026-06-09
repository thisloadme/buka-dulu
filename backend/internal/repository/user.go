package repository

import (
	"database/sql"
	"github.com/riyantobudi/bukadulu/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *domain.User) error {
	var email, phone interface{}
	if user.Email != "" {
		email = user.Email
	}
	if user.Phone != "" {
		phone = user.Phone
	}
	_, err := r.db.Exec(
		`INSERT INTO users (id, role, full_name, email, phone, password_hash, status, created_at, updated_at)
		 VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		user.ID, string(user.Role), user.FullName, email, phone, user.PasswordHash, user.Status, user.CreatedAt, user.UpdatedAt,
	)
	return err
}

// scanUser scans a user row into a User struct, handling NULL fields
func scanUser(scanner interface {
	Scan(dest ...interface{}) error
}) (*domain.User, error) {
	u := &domain.User{}
	var email, phone, lastLogin sql.NullString
	err := scanner.Scan(&u.ID, &u.Role, &u.FullName, &email, &phone, &u.PasswordHash, &u.Status, &lastLogin, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	u.Email = email.String
	u.Phone = phone.String
	if lastLogin.Valid {
		u.LastLoginAt = &lastLogin.String
	}
	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
	row := r.db.QueryRow(
		`SELECT id, role, full_name, email, phone, password_hash, status, last_login_at, created_at, updated_at
		 FROM users WHERE email = ?`, email,
	)
	u, err := scanUser(row)
	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	return u, err
}

func (r *UserRepository) FindByPhone(phone string) (*domain.User, error) {
	row := r.db.QueryRow(
		`SELECT id, role, full_name, email, phone, password_hash, status, last_login_at, created_at, updated_at
		 FROM users WHERE phone = ?`, phone,
	)
	u, err := scanUser(row)
	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	return u, err
}

func (r *UserRepository) FindByID(id string) (*domain.User, error) {
	row := r.db.QueryRow(
		`SELECT id, role, full_name, email, phone, password_hash, status, last_login_at, created_at, updated_at
		 FROM users WHERE id = ?`, id,
	)
	u, err := scanUser(row)
	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	return u, err
}

func (r *UserRepository) UpdateLastLogin(id string) error {
	_, err := r.db.Exec(`UPDATE users SET last_login_at = datetime('now'), updated_at = datetime('now') WHERE id = ?`, id)
	return err
}
