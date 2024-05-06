package update

import (
	"fmt"
	"os"

	"github.com/danilomarques1/fidus/app"
	"github.com/danilomarques1/fidus/terminal"
	"github.com/spf13/cobra"
)

var UpdatePasswordCmd = &cobra.Command{
	Use:   "update",
	Short: "To update a password",
	Long:  "Update a stored password",
	Run:   run,
}

func init() {
	UpdatePasswordCmd.PersistentFlags().String("key", "", "The password key you want to retrieve")
	UpdatePasswordCmd.MarkPersistentFlagRequired("key")
}

func run(cmd *cobra.Command, args []string) {
	key := cmd.Flag("key").Value.String()
	fmt.Print("New password: ")
	newPassword, err := terminal.ReadUserPassword()
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
	if newPassword != confirmPassword {
		fmt.Println("Password needs to match")
		os.Exit(1)
	}

	updatePassword := app.NewUpdatePassword()
	if err := updatePassword.Execute(key, newPassword); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Updated successfuly")
}
