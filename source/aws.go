package source

import (
	"log"
	"path"

	"github.com/charliekenney23/credenv/pkg/env"
	"github.com/urfave/cli"
	"gopkg.in/ini.v1"
)

type awsProfile struct {
	AccessKeyID     string `ini:"aws_access_key_id"`
	SecretAccessKey string `ini:"aws_secret_access_key"`
}

func (p *awsProfile) MapToEnvironment() env.Environment {
	return env.Environment{
		"AWS_ACCESS_KEY_ID":     p.AccessKeyID,
		"AWS_SECRET_ACCESS_KEY": p.SecretAccessKey,
	}
}

// AWS sources AWS environment credentials.
type AWS struct {
}

// GetEnvironment parses the AWS environment credentials and maps
// them to an environment variable set
func (a AWS) GetEnvironment(c *cli.Context) env.Environment {
	profile := c.String("profile")
	if profile == "" {
		profile = "default"
	}

	cred := mustReadFile(path.Join(env.GetHomeDir(), ".aws", "credentials"))

	f, err := ini.Load(cred)
	if err != nil {
		log.Fatalf("Error parsing aws credentials: %v\n", err)
	}

	sec, err := f.GetSection(profile)
	if err != nil {
		log.Fatalf("No such profile `%s` found in aws credentials\n", profile)
	}

	prof := new(awsProfile)

	if err := sec.MapTo(&prof); err != nil {
		log.Fatalf("Error unmarshalling aws profile from credentials: %v\n", err)
	}

	return prof.MapToEnvironment()
}
