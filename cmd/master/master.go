package master

import (
	"fmt"
	"os"
	"syscall"

	"github.com/danilomarques1/fidus/app"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

func RegisterCommand(cmd *cobra.Command, args []string) {
	name := cmd.Flag("name").Value.String()
	email := cmd.Flag("email").Value.String()
	password, err := readUserPassword()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	registerMaster := app.NewRegisterMaster()
	if err := registerMaster.Execute(name, email, password); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Master registered")
}

func Authenticate(cmd *cobra.Command, args []string) {
	email := cmd.Flag("email").Value.String()
	password, err := readUserPassword()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	authenticateMaster := app.NewAuthenticateMaster()
	if err := authenticateMaster.Execute(email, password); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Authenticated")
}

func readUserPassword() (string, error) {
	fmt.Print("Password: ")
	b, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}
	fmt.Println()
	return string(b), nil
}
