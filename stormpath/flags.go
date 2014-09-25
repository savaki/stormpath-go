package main

import "github.com/codegangsta/cli"

const (
	FieldEmail = "email"
)

var (
	FlagEmail = cli.StringFlag{FieldEmail, "", "specify email address"}
)
