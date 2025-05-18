package error

import "errors"

// Authentication errors
var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidPhoneNumber = errors.New("invalid phone number")
	ErrInvalidPassword    = errors.New("invalid password")
	ErrInvalidEmail       = errors.New("invalid email")
	ErrInvalidUserID      = errors.New("invalid user ID")
	ErrInvalidUser        = errors.New("invalid user")
	ErrEmailAlreadyExists = errors.New("email already registered")
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrUserNotFound       = errors.New("user not found")
	ErrAuthHeaderNotFound = errors.New("authorization header not found")
)

// Request validation errors
var (
	ErrRequestBodyRequired = errors.New("invalid request: request body is required")
	ErrInvalidRequestBody  = errors.New("invalid request body")
)

// Token Error
var (
	ErrTokenNotFound           = errors.New("token not found")
	ErrInvalidToken            = errors.New("invalid token")
	ErrTokenExpired            = errors.New("token has expired")
	ErrTokenMismatch           = errors.New("token hash mismatch")
	ErrTokenAlreadyBlacklisted = errors.New("token already blacklisted")
)
