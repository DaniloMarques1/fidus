package terminal

import (
	"fmt"
	"syscall"

	"golang.org/x/term"
)

func ReadUserPassword() (string, error) {
	fmt.Print("Password: ")
	b, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}
	fmt.Println()
	return string(b), nil
}
