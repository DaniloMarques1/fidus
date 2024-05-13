package keys

import (
	"fmt"
	"os"

	"github.com/danilomarques1/fidus/app"
	"github.com/spf13/cobra"
)

var Keys = &cobra.Command{
	Use:   "keys",
	Short: "Retrieve all keys",
	Long:  "It will print all stored keys",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	keys := app.NewKeys()
	retrievedKeys, err := keys.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for idx, key := range retrievedKeys {
		fmt.Printf("%v- %v\n", idx+1, key)
	}
}
