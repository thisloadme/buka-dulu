package domain

type UserRole string

const (
	RoleFounder UserRole = "founder"
	RoleMentor  UserRole = "mentor"
	RoleAdmin   UserRole = "admin"
)

type User struct {
	ID           string   `db:"id" json:"id"`
	Role         UserRole `db:"role" json:"role"`
	FullName     string   `db:"full_name" json:"full_name"`
	Email        string   `db:"email" json:"email,omitempty"`
	Phone        string   `db:"phone" json:"phone,omitempty"`
	PasswordHash string   `db:"password_hash" json:"-"`
	Status       string   `db:"status" json:"status"`
	LastLoginAt  *string  `db:"last_login_at" json:"last_login_at,omitempty"`
	CreatedAt    string   `db:"created_at" json:"created_at"`
	UpdatedAt    string   `db:"updated_at" json:"updated_at"`
}

type RegisterRequest struct {
	FullName string   `json:"full_name"`
	Email    string   `json:"email,omitempty"`
	Phone    string   `json:"phone,omitempty"`
	Password string   `json:"password"`
	Role     UserRole `json:"role,omitempty"`
}

type LoginRequest struct {
	EmailOrPhone string `json:"email_or_phone"`
	Password     string `json:"password"`
}

type AuthResponse struct {
	User      *User  `json:"user"`
	Token     string `json:"token"`
	ExpiresAt string `json:"expires_at"`
}
