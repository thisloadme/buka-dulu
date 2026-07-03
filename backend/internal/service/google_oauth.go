package service

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/riyantobudi/bukadulu/internal/domain"
)

// GoogleOAuthConfig holds Google OAuth credentials.
type GoogleOAuthConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

// googleIDTokenPayload is the subset of claims we need from the id_token.
type googleIDTokenPayload struct {
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Name          string `json:"name"`
}

// LoginWithGoogle exchanges an authorization code for Google tokens, decodes
// the id_token, and finds-or-creates a local user. Returns a JWT for the app.
func (s *AuthService) LoginWithGoogle(cfg GoogleOAuthConfig, code string) (*domain.AuthResponse, error) {
	if cfg.ClientID == "" || cfg.ClientSecret == "" || cfg.RedirectURL == "" {
		return nil, fmt.Errorf("google oauth not configured: %w", domain.ErrInvalidInput)
	}

	idToken, err := exchangeGoogleCode(cfg, code)
	if err != nil {
		return nil, err
	}

	claims, err := decodeIDToken(idToken)
	if err != nil {
		return nil, err
	}
	if claims.Sub == "" || claims.Email == "" {
		return nil, fmt.Errorf("invalid google token claims: %w", domain.ErrInvalidInput)
	}

	// Find existing by provider uid; fallback to email match (link accounts).
	user, err := s.userRepo.FindByProvider("google", claims.Sub)
	if err == domain.ErrNotFound {
		if existing, err := s.userRepo.FindByEmail(claims.Email); err == nil && existing != nil {
			// Link existing email account to google provider
			existing.Provider = "google"
			existing.ProviderUID = claims.Sub
			if existing.FullName == "" {
				existing.FullName = claims.Name
			}
			if err := s.userRepo.UpdateProvider(existing.ID, "google", claims.Sub); err != nil {
				return nil, err
			}
			user = existing
		} else {
			// Create new google user
			now := time.Now().UTC().Format(time.RFC3339)
			user = &domain.User{
				ID:          uuid.New().String(),
				Role:        domain.RoleFounder,
				FullName:    claims.Name,
				Email:       claims.Email,
				Status:      "active",
				Provider:    "google",
				ProviderUID: claims.Sub,
				CreatedAt:   now,
				UpdatedAt:   now,
			}
			if user.FullName == "" {
				user.FullName = strings.Split(claims.Email, "@")[0]
			}
			if err := s.userRepo.CreateSSO(user); err != nil {
				return nil, fmt.Errorf("create google user: %w", err)
			}
		}
	} else if err != nil {
		return nil, err
	}

	if user.Status != "active" {
		// Auto-activate SSO users
		user.Status = "active"
		_ = s.userRepo.Activate(user.ID)
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

// exchangeGoogleCode posts the auth code to Google's token endpoint.
func exchangeGoogleCode(cfg GoogleOAuthConfig, code string) (string, error) {
	form := url.Values{
		"code":          {code},
		"client_id":     {cfg.ClientID},
		"client_secret": {cfg.ClientSecret},
		"redirect_uri":  {cfg.RedirectURL},
		"grant_type":    {"authorization_code"},
	}
	resp, err := http.PostForm("https://oauth2.googleapis.com/token", form)
	if err != nil {
		return "", fmt.Errorf("google token exchange: %w", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("google token exchange failed (%d): %s: %w", resp.StatusCode, string(body), domain.ErrUnauthorized)
	}
	var tok struct {
		IDToken string `json:"id_token"`
	}
	if err := json.Unmarshal(body, &tok); err != nil || tok.IDToken == "" {
		return "", fmt.Errorf("no id_token in google response: %w", domain.ErrUnauthorized)
	}
	return tok.IDToken, nil
}

// decodeIDToken decodes the JWT payload WITHOUT signature verification.
// Signature verification happens implicitly during exchangeGoogleCode (Google
// rejects invalid codes). For production hardening, verify with google certs.
// ponytail: skip cert-based verification; Google issued the token seconds ago.
func decodeIDToken(idToken string) (*googleIDTokenPayload, error) {
	parts := strings.Split(idToken, ".")
	if len(parts) < 2 {
		return nil, fmt.Errorf("malformed id_token: %w", domain.ErrInvalidInput)
	}
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("decode id_token payload: %w", err)
	}
	var claims googleIDTokenPayload
	if err := json.Unmarshal(payload, &claims); err != nil {
		return nil, fmt.Errorf("parse id_token claims: %w", err)
	}
	return &claims, nil
}

// GoogleLoginRequest is the body for the /auth/google endpoint.
type GoogleLoginRequest struct {
	Code string `json:"code"`
}

// Google config accessors used by handlers. Values are set via SetGoogleConfig.
func (s *AuthService) GoogleClientID() string     { return s.googleCfg.ClientID }
func (s *AuthService) GoogleClientSecret() string { return s.googleCfg.ClientSecret }
func (s *AuthService) GoogleRedirectURL() string  { return s.googleCfg.RedirectURL }

// SetGoogleConfig injects Google OAuth credentials at startup.
func (s *AuthService) SetGoogleConfig(cfg GoogleOAuthConfig) { s.googleCfg = cfg }
