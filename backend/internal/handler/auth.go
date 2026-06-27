package handler

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/riyantobudi/bukadulu/internal/domain"
	"github.com/riyantobudi/bukadulu/internal/middleware"
	"github.com/riyantobudi/bukadulu/internal/service"
)

type AuthHandler struct {
	authSvc *service.AuthService
}

func NewAuthHandler(authSvc *service.AuthService) *AuthHandler {
	return &AuthHandler{authSvc: authSvc}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req domain.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": "invalid request body"})
		return
	}

	if req.Email == "" {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": "email is required for registration"})
		return
	}
	if req.FullName == "" {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": "full_name is required"})
		return
	}
	if len(req.Password) < 8 {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": "password must be at least 8 characters"})
		return
	}

	resp, err := h.authSvc.Register(&req)
	if err != nil {
		slog.Error("register failed", "error", err)
		writeJSON(w, http.StatusConflict, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusCreated, resp)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req domain.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": "invalid request body"})
		return
	}

	resp, err := h.authSvc.Login(&req)
	if err != nil {
		// Distinguish between "not verified" and "wrong password"
		writeJSON(w, http.StatusUnauthorized, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, resp)
}

func (h *AuthHandler) VerifyOTP(w http.ResponseWriter, r *http.Request) {
	var req domain.OTPVerifyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": "invalid request body"})
		return
	}

	if req.Email == "" || req.OTP == "" {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": "email and otp are required"})
		return
	}

	resp, err := h.authSvc.VerifyOTP(&req)
	if err != nil {
		slog.Error("verify OTP failed", "error", err)
		writeJSON(w, http.StatusUnauthorized, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, resp)
}

func (h *AuthHandler) ResendOTP(w http.ResponseWriter, r *http.Request) {
	var req domain.OTPResendRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": "invalid request body"})
		return
	}

	if req.Email == "" {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": "email is required"})
		return
	}

	if err := h.authSvc.ResendOTP(&req); err != nil {
		slog.Error("resend OTP failed", "error", err)
		writeJSON(w, http.StatusConflict, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "OTP sent to your email"})
}

func GetUserID(ctx context.Context) string {
	id, _ := ctx.Value(middleware.UserIDKey).(string)
	return id
}

func GetUserRole(ctx context.Context) string {
	role, _ := ctx.Value(middleware.UserRoleKey).(string)
	return role
}
