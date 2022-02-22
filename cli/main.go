package main

import (
	"fmt"
	cli "github.com/jawher/mow.cli"
	"os"
)

// Global options available to any of the commands
var filename *string

func main() {
	app := cli.App("vault", "Password Keeper")

	// Define our top-level global options
	filename = app.StringOpt("f file", os.Getenv("HOME")+"/.safe", "Path to safe")

	// Define our command structure for usage like this:
	app.Command("list", "list accounts", cmdList)
	app.Command("creds", "display account credentials", cmdCreds)
	app.Command("config", "manage accounts", func(config *cli.Cmd) {
		config.Command("list", "list accounts", cmdList)
		config.Command("add", "add an account", cmdAdd)
		config.Command("remove", "remove an account(s)", cmdRemove)
	})

	app.Run(os.Args)
}

// Sample use: vault list OR vault config list
func cmdList(cmd *cli.Cmd) {
	cmd.Action = func() {
		fmt.Printf("list the contents of the safe here")
	}
}

// Sample use: vault creds reddit.com
func cmdCreds(cmd *cli.Cmd) {
	cmd.Spec = "ACCOUNT"
	account := cmd.StringArg("ACCOUNT", "", "Name of account")
	cmd.Action = func() {
		fmt.Printf("display account info for %s\n", *account)
	}
}

// Sample use: vault config add reddit.com -u username -p password
func cmdAdd(cmd *cli.Cmd) {
	cmd.Spec = "ACCOUNT [ -u=<username> ] [ -p=<password> ]"
	var (
		account  = cmd.StringArg("ACCOUNT", "", "Account name")
		username = cmd.StringOpt("u username", "admin", "Account username")
		password = cmd.StringOpt("p password", "admin", "Account password")
	)
	cmd.Action = func() {
		fmt.Printf("Adding account %s:%s@%s", *username, *password, *account)
	}
}

// Sample use: vault config remove reddit.com twitter.com
func cmdRemove(cmd *cli.Cmd) {
	cmd.Spec = "ACCOUNT..."
	var (
		accounts = cmd.StringsArg("ACCOUNT", nil, "Account names to remove")
	)
	cmd.Action = func() {
		fmt.Printf("Deleting accounts: %v", *accounts)
	}
}
