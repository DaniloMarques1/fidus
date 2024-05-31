package clierror

import (
	"errors"
	"fmt"
)

func ErrInvalidToken() error {
	return errors.New("You need to authenticate")
}

func ErrInvalidParameters(msg string) error {
	return fmt.Errorf("Invalid parameters see help for usage. %v", msg)
}
