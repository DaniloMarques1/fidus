package master

import (
	"fmt"
	"os"

	"github.com/danilomarques1/fidus/app"
	"github.com/danilomarques1/fidus/terminal"
	"github.com/spf13/cobra"
)

func RegisterCommand(cmd *cobra.Command, args []string) {
	name := cmd.Flag("name").Value.String()
	email := cmd.Flag("email").Value.String()
	password, err := terminal.ReadUserPassword()
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
	password, err := terminal.ReadUserPassword()
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
