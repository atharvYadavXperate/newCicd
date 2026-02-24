package custiomeerror

import "errors"

var (
	ErrDatabaseConnectionFailed = errors.New("Failed to make connection with database")
)
