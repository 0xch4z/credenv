package source

import (
	"encoding/json"
	"log"
	"path"

	"github.com/charliekenney23/credenv/pkg/env"
	"github.com/urfave/cli"
)

// AccessTokenSet represents a mapping of services to their
// respective API tokens.
type AccessTokenSet map[string]string

type pulumiCredentials struct {
	Current      string         `json:"current"`
	AccessTokens AccessTokenSet `json:"accessTokens"`
}

// Pulumi sources Pulumi environment credentials.
type Pulumi struct {
}

// GetEnvironment parses the Pulumi environment credentials and
// maps them to an environment variable set.
func (p Pulumi) GetEnvironment(c *cli.Context) env.Environment {
	service := c.String("service")
	if service == "" {
		service = "https://api.pulumi.com"
	}

	cred := mustReadFile(path.Join(env.GetHomeDir(), ".pulumi", "credentials.json"))

	var pc pulumiCredentials
	if err := json.Unmarshal(cred, &pc); err != nil {
		log.Fatalf("Error parsing pulumi credentials: %v\n", err)
	}

	tok, ok := pc.AccessTokens[service]
	if !ok {
		log.Fatalf("No token present for Pulumi service '%s'\n", service)
	}

	return env.Environment{
		"PULUMI_ACCESS_TOKEN": tok,
	}
}
