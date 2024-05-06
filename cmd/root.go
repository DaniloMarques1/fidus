package cmd

import (
	"log"

	"github.com/danilomarques1/fidus/cmd/master/authenticate"
	"github.com/danilomarques1/fidus/cmd/master/register"
	"github.com/danilomarques1/fidus/cmd/password/deletepassword"
	"github.com/danilomarques1/fidus/cmd/password/retrieve"
	"github.com/danilomarques1/fidus/cmd/password/store"
	"github.com/danilomarques1/fidus/cmd/password/update"
	"github.com/spf13/cobra"
)

func Execute() {
	cmd := &cobra.Command{
		Short: "Password manager to store your secrets",
		Long:  "A password manager that you can create a master account and then store passwords",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	cmd.AddCommand(register.RegisterCmd)
	cmd.AddCommand(authenticate.AuthenticateCmd)
	cmd.AddCommand(store.StorePasswordCmd)
	cmd.AddCommand(retrieve.RetrievePasswordCmd)
	cmd.AddCommand(deletepassword.DeletePasswordCmd)
	cmd.AddCommand(update.UpdatePasswordCmd)

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
