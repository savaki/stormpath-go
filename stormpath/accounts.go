package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/savaki/stormpath-go/accounts"
	"github.com/savaki/stormpath-go/auth"
	"log"
	"os"
)

var AccountCommand cli.Command = cli.Command{
	Name:  "account",
	Usage: "manage your stormpath accounts",
	Subcommands: []cli.Command{
		{
			Name:   "search",
			Usage:  "find accounts matching the specified criteria",
			Flags:  []cli.Flag{FlagEmail},
			Action: accountSearch,
		},
	},
}

func accountSearch(ctx *cli.Context) {
	apiKey, err := auth.EnvAuth()
	if err != nil {
		log.Fatalln(err)
	}

	uri := os.Getenv("STORMPATH_URL")
	if uri == "" {
		log.Fatalln("Please set the STORMPATH_URL with your application url")
	}

	api := accounts.FromUrl(apiKey, uri)
	email := ctx.String(FieldEmail)
	users, err := api.Search(email)
	if err != nil {
		log.Fatalln(err)
	}

	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(data))
}
