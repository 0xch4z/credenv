package credenv

import (
	"fmt"
	"log"

	"github.com/charliekenney23/credenv/pkg/env"
	"github.com/charliekenney23/credenv/source"
	"github.com/urfave/cli"
)

// Command handles user input.
func Command(c *cli.Context) {
	args := c.Args()
	src := args.Get(0)

	if src == "" {
		log.Fatal("source is required")
	}

	s := getEnvironmentSource(c, src)
	e := s.GetEnvironment(c)
	printEnvironment(e)
}

// printEnvironment prints the environment to stdout
func printEnvironment(e env.Environment) {
	for name, value := range e {
		fmt.Printf("export %s='%s'\n", name, value)
	}
}

// getEnvironmentSource maps the source arg to a Source
func getEnvironmentSource(c *cli.Context, src string) Source {
	var s Source
	switch src {
	case "aws":
		s = &source.AWS{}
	default:
		log.Fatalf("Unsupported source `%s`\n", src)
	}

	return s
}
