package models

import "errors"

var (
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateUsername = errors.New("models: username already exists")
	ErrNoRecord = errors.New("models: record not found")
)