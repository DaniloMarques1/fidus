package main

import (
	"log"

	"github.com/danilomarques1/fidus/cmd"
	"github.com/danilomarques1/fidus/config"
)

func main() {
	cfg := config.NewConfig()
	if err := cfg.CreateConfigFolder(); err != nil {
		log.Fatal(err)
	}
	if err := cfg.SetupLogger(); err != nil {
		log.Fatal(err)
	}
	cmd.Execute()
}
