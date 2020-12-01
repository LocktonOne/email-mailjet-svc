package main

import (
	"os"

	"gitlab.com/tokend/notifications/email-mailjet-svc/internal/cli"
)

func main() {
	//os.Setenv("KV_VIPER_FILE", "./config.local.yaml")
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
