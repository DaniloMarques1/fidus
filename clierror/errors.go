package clierror

import "errors"

func ErrInvalidToken() error {
	return errors.New("You need to authenticate")
}

func ErrInvalidParameters() error {
	return errors.New("Invalid parameters see help for usage")
}
