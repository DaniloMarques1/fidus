package authenticate

import (
	"fmt"
	"os"

	"github.com/danilomarques1/fidus/app"
	"github.com/danilomarques1/fidus/terminal"
	"github.com/spf13/cobra"
)

var AuthenticateCmd = &cobra.Command{
	Use:   "authenticate",
	Short: "Authenticate a master to create password",
	Long:  "You will need to get authenticated before you can store/retrieve passwords",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	email := cmd.Flag("email").Value.String()
	fmt.Print("Password: ")
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

func init() {
	AuthenticateCmd.PersistentFlags().String("email", "", "Master email")
	AuthenticateCmd.MarkPersistentFlagRequired("email")
}
