package validate

import "regexp"

func Email(value string) bool {
	rg := "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"
	pattern, err := regexp.Compile(rg)
	if err != nil {
		return false
	}

	return pattern.MatchString(value)
	return true
}
