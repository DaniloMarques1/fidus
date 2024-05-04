package register

import (
	"fmt"
	"os"

	"github.com/danilomarques1/fidus/app"
	"github.com/danilomarques1/fidus/terminal"
	"github.com/spf13/cobra"
)

var RegisterCmd = &cobra.Command{
	Use:   "register",
	Short: "Create a new master account",
	Long:  "Command to create a new master account. You need to provide the email, name and a password",
	Run:   run,
}

func init() {
	RegisterCmd.PersistentFlags().String("name", "", "Master name")
	RegisterCmd.PersistentFlags().String("email", "", "Master email")
	RegisterCmd.MarkPersistentFlagRequired("name")
	RegisterCmd.MarkPersistentFlagRequired("email")
}

func run(cmd *cobra.Command, args []string) {
	name := cmd.Flag("name").Value.String()
	email := cmd.Flag("email").Value.String()
	fmt.Print("Password: ")
	password, err := terminal.ReadUserPassword()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Print("Confirm password: ")
	confirmPassword, err := terminal.ReadUserPassword()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if password != confirmPassword {
		fmt.Println("Password needs to match")
		os.Exit(1)
	}
	registerMaster := app.NewRegisterMaster()
	if err := registerMaster.Execute(name, email, password); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Master registered")

}
