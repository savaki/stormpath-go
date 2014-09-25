package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "stormpath"
	app.Version = "0.1"
	app.Commands = []cli.Command{
		AccountCommand,
	}

	app.Run(os.Args)
}
