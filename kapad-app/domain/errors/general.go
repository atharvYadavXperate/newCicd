package custiomeerror

import "errors"

var (
	ErrParseError     = errors.New("Failed to parse body unexpected json")
	ErrRequiredFields = errors.New("All fields are required")
)
