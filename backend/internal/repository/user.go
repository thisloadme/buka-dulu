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
	var email, phone, otpCode, otpExp, otpVerified interface{}
	if user.Email != "" {
		email = user.Email
	}
	if user.Phone != "" {
		phone = user.Phone
	}
	if user.OTPCode != nil {
		otpCode = *user.OTPCode
	}
	if user.OTPExpiresAt != nil {
		otpExp = *user.OTPExpiresAt
	}
	if user.OTPVerifiedAt != nil {
		otpVerified = *user.OTPVerifiedAt
	}

	_, err := r.db.Exec(
		`INSERT INTO users (id, role, full_name, email, phone, password_hash, status,
		 otp_code, otp_expires_at, otp_verified_at, created_at, updated_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`,
		user.ID, string(user.Role), user.FullName, email, phone, user.PasswordHash, user.Status,
		otpCode, otpExp, otpVerified, user.CreatedAt, user.UpdatedAt,
	)
	return err
}

// userColumns returns the SELECT column list including OTP + provider + quota fields
func userColumns() string {
	return `id, role, full_name, email, phone, password_hash, status,
		otp_code, otp_expires_at, otp_verified_at, last_login_at,
		provider, provider_uid, free_quota_used, created_at, updated_at`
}

// scanUser scans a user row into a User struct, handling NULL fields
func scanUser(scanner interface {
	Scan(dest ...interface{}) error
}) (*domain.User, error) {
	u := &domain.User{}
	var email, phone, otpCode, otpExp, otpVerified, lastLogin, providerUID sql.NullString
	err := scanner.Scan(
		&u.ID, &u.Role, &u.FullName, &email, &phone, &u.PasswordHash, &u.Status,
		&otpCode, &otpExp, &otpVerified, &lastLogin,
		&u.Provider, &providerUID, &u.FreeQuotaUsed, &u.CreatedAt, &u.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	u.Email = email.String
	u.Phone = phone.String
	if otpCode.Valid {
		u.OTPCode = &otpCode.String
	}
	if otpExp.Valid {
		u.OTPExpiresAt = &otpExp.String
	}
	if otpVerified.Valid {
		u.OTPVerifiedAt = &otpVerified.String
	}
	if lastLogin.Valid {
		u.LastLoginAt = &lastLogin.String
	}
	if u.Provider == "" {
		u.Provider = "email"
	}
	u.ProviderUID = providerUID.String
	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
	row := r.db.QueryRow(
		`SELECT `+userColumns()+` FROM users WHERE email = $1`, email,
	)
	u, err := scanUser(row)
	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	return u, err
}

func (r *UserRepository) FindByPhone(phone string) (*domain.User, error) {
	row := r.db.QueryRow(
		`SELECT `+userColumns()+` FROM users WHERE phone = $1`, phone,
	)
	u, err := scanUser(row)
	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	return u, err
}

func (r *UserRepository) FindByID(id string) (*domain.User, error) {
	row := r.db.QueryRow(
		`SELECT `+userColumns()+` FROM users WHERE id = $1`, id,
	)
	u, err := scanUser(row)
	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	return u, err
}

func (r *UserRepository) UpdateLastLogin(id string) error {
	_, err := r.db.Exec(`UPDATE users SET last_login_at = CURRENT_TIMESTAMP, updated_at = CURRENT_TIMESTAMP WHERE id = $1`, id)
	return err
}

func (r *UserRepository) VerifyOTP(id string, verifiedAt string) error {
	_, err := r.db.Exec(
		`UPDATE users SET status = 'active', otp_code = NULL, otp_expires_at = NULL,
		 otp_verified_at = $1, updated_at = $1 WHERE id = $2`,
		verifiedAt, id,
	)
	return err
}

func (r *UserRepository) UpdateOTP(id string, otpCode, otpExpiresAt string) error {
	_, err := r.db.Exec(
		`UPDATE users SET otp_code = $1, otp_expires_at = $2, updated_at = $2 WHERE id = $3`,
		otpCode, otpExpiresAt, id,
	)
	return err
}

// FindByProvider finds a user by SSO provider + provider uid (e.g. google sub).
func (r *UserRepository) FindByProvider(provider, providerUID string) (*domain.User, error) {
	row := r.db.QueryRow(
		`SELECT `+userColumns()+` FROM users WHERE provider = $1 AND provider_uid = $2`,
		provider, providerUID,
	)
	u, err := scanUser(row)
	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	return u, err
}

// CreateSSO creates a user from an SSO provider (no password, already active).
func (r *UserRepository) CreateSSO(user *domain.User) error {
	_, err := r.db.Exec(
		`INSERT INTO users (id, role, full_name, email, phone, password_hash, status,
		 provider, provider_uid, created_at, updated_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		user.ID, string(user.Role), user.FullName, user.Email, nil, "", user.Status,
		user.Provider, user.ProviderUID, user.CreatedAt, user.UpdatedAt,
	)
	return err
}

// IncrementFreeQuota atomically increments the free quota counter.
func (r *UserRepository) IncrementFreeQuota(userID string) error {
	_, err := r.db.Exec(
		`UPDATE users SET free_quota_used = free_quota_used + 1, updated_at = CURRENT_TIMESTAMP WHERE id = $1`,
		userID,
	)
	return err
}

// GetFreeQuotaUsed returns how many free idea validations the user has consumed.
func (r *UserRepository) GetFreeQuotaUsed(userID string) (int, error) {
	var used int
	err := r.db.QueryRow(`SELECT free_quota_used FROM users WHERE id = $1`, userID).Scan(&used)
	if err == sql.ErrNoRows {
		return 0, domain.ErrNotFound
	}
	return used, err
}

// UpdateProvider links a SSO provider to an existing user.
func (r *UserRepository) UpdateProvider(id, provider, providerUID string) error {
	_, err := r.db.Exec(
		`UPDATE users SET provider = $1, provider_uid = $2, updated_at = CURRENT_TIMESTAMP WHERE id = $3`,
		provider, providerUID, id,
	)
	return err
}

// Activate sets user status to active (used when completing SSO / verification).
func (r *UserRepository) Activate(id string) error {
	_, err := r.db.Exec(
		`UPDATE users SET status = 'active', updated_at = CURRENT_TIMESTAMP WHERE id = $1`, id,
	)
	return err
}
