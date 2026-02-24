package custiomeerror

import "errors"

var (
	ErrParseError = errors.New("Failed to parse body unexpected json")
)
