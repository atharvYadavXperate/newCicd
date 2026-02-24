package helpers

import "strings"

func IsEmptyString(value string) bool {
	if TrimSpacesInString(value) == "" {
		return true
	}
	return false
}

func TrimSpacesInString(value string) string {
	return strings.TrimSpace(value)
}
