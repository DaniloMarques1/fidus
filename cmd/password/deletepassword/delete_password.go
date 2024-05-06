package deletepassword

import (
	"fmt"
	"os"

	"github.com/danilomarques1/fidus/app"
	"github.com/spf13/cobra"
)

var DeletePasswordCmd = &cobra.Command{
	Use:   "delete",
	Short: "To delete a password",
	Long:  "Deletes a stored password",
	Run:   run,
}

func init() {
	DeletePasswordCmd.PersistentFlags().String("key", "", "The password key you want to delete")
	DeletePasswordCmd.MarkPersistentFlagRequired("key")
}

func run(cmd *cobra.Command, args []string) {
	key := cmd.Flag("key").Value.String()
	deletePassword := app.NewDeletePassword()
	if err := deletePassword.Execute(key); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Deleted successfuly")
}
