package mapping

import "errors"

var (
	ErrSystem         = errors.New("system error")
	ErrNotFound       = errors.New("error not found")
	ErrBadRequest     = errors.New("bad request")
	ErrEligibility    = errors.New("error not eligible")
	ErrInvalidPayload = errors.New("invalid payload")
)
