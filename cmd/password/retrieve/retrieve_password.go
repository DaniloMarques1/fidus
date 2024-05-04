package retrieve

import (
	"fmt"
	"os"

	"github.com/danilomarques1/fidus/app"
	"github.com/spf13/cobra"
)

var RetrievePasswordCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "To retrieve a password",
	Long:  "Retrieve a stored password",
	Run:   run,
}

func init() {
	RetrievePasswordCmd.PersistentFlags().String("key", "", "The password key you want to retrieve")
	RetrievePasswordCmd.MarkPersistentFlagRequired("key")
}

func run(cmd *cobra.Command, args []string) {
	key := cmd.Flag("key").Value.String()
	retrievePassword := app.NewRetrievePassword()
	password, err := retrievePassword.Execute(key)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(password)
}
