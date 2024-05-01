package cmd

import (
	"log"

	"github.com/danilomarques1/fidus/cmd/master"
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

	registerCmd := &cobra.Command{
		Use:   "register",
		Short: "Create a new master account",
		Long:  "Command to create a new master account. You need to provide the email, name and a password",
		Run:   master.RegisterCommand,
	}
	authenticateCmd := &cobra.Command{
		Use:   "authenticate",
		Short: "Authenticate a master to create password",
		Long:  "You will need to get authenticated before you can store/retrieve passwords",
		Run:   master.Authenticate,
	}

	registerCmd.PersistentFlags().String("name", "", "Master name")
	registerCmd.PersistentFlags().String("email", "", "Master email")

	authenticateCmd.PersistentFlags().String("email", "", "Master name")

	cmd.AddCommand(registerCmd)
	cmd.AddCommand(authenticateCmd)

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
