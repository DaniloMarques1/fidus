package reset

import (
	"fmt"
	"os"

	"github.com/danilomarques1/fidus/app"
	"github.com/danilomarques1/fidus/terminal"
	"github.com/spf13/cobra"
)

var ResetMasterPasswordCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset master password.",
	Long:  "You can use this command to change your password and you must use this command when you password expires.",
	Run:   run,
}

func init() {
	ResetMasterPasswordCmd.PersistentFlags().String("email", "", "Master email")
	ResetMasterPasswordCmd.MarkPersistentFlagRequired("email")
}

func run(cmd *cobra.Command, args []string) {
	email := cmd.Flag("email").Value.String()
	fmt.Print("Current password: ")
	currentPassword, err := terminal.ReadUserPassword()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Print("New password: ")
	nPassword, err := terminal.ReadUserPassword()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Print("Confirm new password: ")
	cPassword, err := terminal.ReadUserPassword()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if nPassword != cPassword {
		fmt.Println("Password needs to match")
		os.Exit(1)
	}

	resetPassword := app.NewResetMasterPassword()
	if err := resetPassword.Execute(email, currentPassword, nPassword); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Password reseted")
}
