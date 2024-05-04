package store

import (
	"fmt"
	"log"
	"os"

	"github.com/danilomarques1/fidus/app"
	"github.com/danilomarques1/fidus/terminal"
	"github.com/spf13/cobra"
)

var StorePasswordCmd = &cobra.Command{
	Use:   "store",
	Short: "To store a password",
	Long:  "Created a new master password",
	Run:   run,
}

func init() {
	StorePasswordCmd.PersistentFlags().String("key", "", "The password key you want to store")
	StorePasswordCmd.MarkPersistentFlagRequired("key")
}

func run(cmd *cobra.Command, args []string) {
	key := cmd.Flag("key").Value.String()
	fmt.Print("Password: ")
	pwd, err := terminal.ReadUserPassword()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Print("Confirm Password: ")
	confirmPwd, err := terminal.ReadUserPassword()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if pwd != confirmPwd {
		fmt.Println("passwords needs to match")
		os.Exit(1)
	}

	storePassword := app.NewStorePassword()
	if err := storePassword.Execute(key, pwd); err != nil {
		log.Fatal(err)
	}
}
