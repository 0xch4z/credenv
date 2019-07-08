package main

import (
	"github.com/charliekenney23/credenv"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Action = credenv.Command
	app.RunAndExitOnError()
}
