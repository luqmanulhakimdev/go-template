package middlewares

import "errors"

var (
	ErrUnknownError          = errors.New("ERR_UNKNOWN")
	ErrUnauthorizedError     = errors.New("ERR_UNAUTHORIZED")
)
