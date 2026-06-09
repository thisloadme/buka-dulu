package domain

import "errors"

var (
	ErrNotFound       = errors.New("resource not found")
	ErrConflict       = errors.New("resource already exists")
	ErrInvalidInput   = errors.New("invalid input")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrForbidden      = errors.New("forbidden")
	ErrStageGate      = errors.New("stage gate not satisfied")
	ErrAIService      = errors.New("AI service unavailable")
	ErrInternal       = errors.New("internal server error")
	ErrDuplicateEntry = errors.New("duplicate entry")
	ErrTokenExpired   = errors.New("token expired")
)

type AppError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func NewAppError(code, message string, err error) *AppError {
	return &AppError{Code: code, Message: message, Err: err}
}

var (
	AppErrNotFound       = NewAppError("ERR_NF_001", "Resource not found", ErrNotFound)
	AppErrInvalidInput   = NewAppError("ERR_VAL_001", "Invalid input", ErrInvalidInput)
	AppErrUnauthorized   = NewAppError("ERR_AUTH_001", "Unauthorized", ErrUnauthorized)
	AppErrForbidden      = NewAppError("ERR_AUTH_003", "Forbidden", ErrForbidden)
	AppErrDuplicateEntry = NewAppError("ERR_VAL_002", "Duplicate entry", ErrDuplicateEntry)
	AppErrAIService      = NewAppError("ERR_LLM_001", "AI service unavailable", ErrAIService)
	AppErrStageGate      = NewAppError("ERR_GATE_001", "Stage gate not satisfied", ErrStageGate)
	AppErrTokenExpired   = NewAppError("ERR_AUTH_002", "Token expired", ErrTokenExpired)
)
