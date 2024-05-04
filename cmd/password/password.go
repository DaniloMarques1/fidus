package password

import (
	"fmt"
	"log"
	"os"

	"github.com/danilomarques1/fidus/app"
	"github.com/danilomarques1/fidus/terminal"
	"github.com/spf13/cobra"
)

func StorePassword(cmd *cobra.Command, args []string) {
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
		log.Fatal(err)
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
