package main

import (
	"fmt"

	"github.com/charliekenney23/credenv"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Action = credenv.Command
	app.Version = fmt.Sprintf("%s (commit %s)", credenv.Version, credenv.GitCommit)
	app.RunAndExitOnError()
}
