package source

import (
	"log"
	"path"

	"github.com/charliekenney23/credenv/pkg/env"
	"github.com/urfave/cli"
	"gopkg.in/ini.v1"
)

type linodeMetadata struct {
	DefaultUser string `ini:"default-user"`
}

type linodeUser struct {
	Token string `ini:"token"`
}

func (p *linodeUser) MapToEnvironment() env.Environment {
	return env.Environment{
		"LINODE_TOKEN":     p.Token,
		"LINODE_API_TOKEN": p.Token,
	}
}

// Linode sources Linode environment credentials.
type Linode struct {
}

// GetEnvironment parses the Linode environment credentials and maps
// them to an environment variable set
func (l Linode) GetEnvironment(c *cli.Context) env.Environment {
	conf := mustReadFile((path.Join(env.GetHomeDir(), ".config", "linode-cli")))

	f, err := ini.Load(conf)
	if err != nil {
		log.Fatalf("Error parsing linode credentials: %v\n", err)
	}

	user := ""
	meta := new(linodeMetadata)
	metaSec, _ := f.GetSection("DEFAULT")

	if err := metaSec.MapTo(&meta); err == nil {
		user = meta.DefaultUser
	}

	if u := c.String("user"); u != "" {
		user = u
	}

	sec, err := f.GetSection(user)
	if err != nil {
		log.Fatalf("No such user `%s` found in linode credentials\n", user)
	}

	u := new(linodeUser)
	if err := sec.MapTo(&u); err != nil {
		log.Fatalf("Error unmarshalling aws profile from credentials: %v\n", err)
	}

	return u.MapToEnvironment()
}
