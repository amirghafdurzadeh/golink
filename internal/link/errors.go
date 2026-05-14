package link

import "errors"

var (
	ErrCodeAlreadyExists = errors.New("code already exists")
	ErrNotFound          = errors.New("not found")
)
