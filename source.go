package credenv

import (
	"github.com/charliekenney23/credenv/pkg/env"
	"github.com/urfave/cli"
)

// Source represents a environment credentials source.
type Source interface {
	GetEnvironment(c *cli.Context) env.Environment
}
