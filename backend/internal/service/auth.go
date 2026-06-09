package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/riyantobudi/bukadulu/internal/domain"
	"github.com/riyantobudi/bukadulu/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo  *repository.UserRepository
	jwtSecret string
	jwtExpiry int
}

func NewAuthService(userRepo *repository.UserRepository, jwtSecret string, jwtExpiry int) *AuthService {
	return &AuthService{userRepo: userRepo, jwtSecret: jwtSecret, jwtExpiry: jwtExpiry}
}

func (s *AuthService) Register(req *domain.RegisterRequest) (*domain.AuthResponse, error) {
	if req.Email != "" {
		if _, err := s.userRepo.FindByEmail(req.Email); err == nil {
			return nil, fmt.Errorf("email already registered: %w", domain.ErrDuplicateEntry)
		}
	}
	if req.Phone != "" {
		if _, err := s.userRepo.FindByPhone(req.Phone); err == nil {
			return nil, fmt.Errorf("phone already registered: %w", domain.ErrDuplicateEntry)
		}
	}

	if req.Role == "" {
		req.Role = domain.RoleFounder
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}

	now := time.Now().UTC().Format(time.RFC3339)
	user := &domain.User{
		ID:           uuid.New().String(),
		Role:         req.Role,
		FullName:     req.FullName,
		Email:        req.Email,
		Phone:        req.Phone,
		PasswordHash: string(hash),
		Status:       "active",
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	token, err := s.generateToken(user.ID, user.Role)
	if err != nil {
		return nil, err
	}

	return &domain.AuthResponse{
		User:      user,
		Token:     token,
		ExpiresAt: time.Now().Add(time.Duration(s.jwtExpiry) * time.Hour).Format(time.RFC3339),
	}, nil
}

func (s *AuthService) Login(req *domain.LoginRequest) (*domain.AuthResponse, error) {
	user, err := s.userRepo.FindByEmail(req.EmailOrPhone)
	if err != nil {
		user, err = s.userRepo.FindByPhone(req.EmailOrPhone)
		if err != nil {
			return nil, fmt.Errorf("invalid credentials: %w", domain.ErrUnauthorized)
		}
	}

	if user.Status != "active" {
		return nil, fmt.Errorf("account suspended: %w", domain.ErrForbidden)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, fmt.Errorf("invalid credentials: %w", domain.ErrUnauthorized)
	}

	s.userRepo.UpdateLastLogin(user.ID)

	token, err := s.generateToken(user.ID, user.Role)
	if err != nil {
		return nil, err
	}

	return &domain.AuthResponse{
		User:      user,
		Token:     token,
		ExpiresAt: time.Now().Add(time.Duration(s.jwtExpiry) * time.Hour).Format(time.RFC3339),
	}, nil
}

func (s *AuthService) ValidateToken(tokenString string) (string, domain.UserRole, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(s.jwtSecret), nil
	})
	if err != nil {
		return "", "", fmt.Errorf("invalid token: %w", domain.ErrUnauthorized)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", "", fmt.Errorf("invalid token claims: %w", domain.ErrUnauthorized)
	}

	userID, _ := claims["sub"].(string)
	role, _ := claims["role"].(string)

	if userID == "" {
		return "", "", fmt.Errorf("invalid token: %w", domain.ErrUnauthorized)
	}

	return userID, domain.UserRole(role), nil
}

func (s *AuthService) generateToken(userID string, role domain.UserRole) (string, error) {
	claims := jwt.MapClaims{
		"sub":  userID,
		"role": string(role),
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Duration(s.jwtExpiry) * time.Hour).Unix(),
		"iss":  "bukadulu",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}
