package validate

import (
	"regexp"
)

func MasterPassword(value string) bool {
	rgString := regexp.MustCompile(`[a-z].*[A-Z]|[A-Z].*[a-z]`)
	rgNumber := regexp.MustCompile(`\d`)
	rgSymbol := regexp.MustCompile(`[^a-zA-Z0-9]`)

	return rgString.MatchString(value) && rgNumber.MatchString(value) && rgSymbol.MatchString(value) && len(value) >= 8
}
