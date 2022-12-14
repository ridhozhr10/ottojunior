package main

import (
	"os"

	"github.com/ridhozhr10/ottojunior/cmd/emoney-service/cliaction"
	"github.com/ridhozhr10/ottojunior/cmd/emoney-service/engine"

	_ "github.com/ridhozhr10/ottojunior/docs"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:    "port",
		EnvVars: []string{"PORT"},
		Value:   ":3000",
		Usage:   "http port",
	},
	&cli.StringFlag{
		Name:    "db-host",
		EnvVars: []string{"DB_HOST"},
		Value:   "localhost",
		Usage:   "database host",
	},
	&cli.StringFlag{
		Name:    "db-port",
		EnvVars: []string{"DB_PORT"},
		Value:   "5432",
		Usage:   "database port",
	},
	&cli.StringFlag{
		Name:    "db-name",
		EnvVars: []string{"DB_NAME"},
		Value:   "ottojunior",
		Usage:   "database username",
	},
	&cli.StringFlag{
		Name:    "db-user",
		EnvVars: []string{"DB_USER"},
		Value:   "postgres",
		Usage:   "database username",
	},
	&cli.StringFlag{
		Name:    "db-pass",
		EnvVars: []string{"DB_PASS"},
		Value:   "postgres",
		Usage:   "database password",
	},
	&cli.StringFlag{
		Name:    "biller-service-base-url",
		EnvVars: []string{"BILLER_SERVICE_BASE_URL"},
		Value:   "https://phoenix-imkas.ottodigital.id",
		Usage:   "biller service base url",
	},
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
}

// @title           ottojunior
// @version         1.0
// @description     ridho otto junior
// @termsOfService  http://swagger.io/terms/

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	app := cli.App{
		Name: "emoney-service",
		Commands: []*cli.Command{
			{
				Name:  "serve-emoney",
				Usage: "serve http api servers",
				Flags: flags,
				Action: func(c *cli.Context) error {
					config := engine.Config{
						DBHost:               c.String("db-host"),
						DBUser:               c.String("db-user"),
						DBPass:               c.String("db-pass"),
						DBPort:               c.String("db-port"),
						DBName:               c.String("db-name"),
						Port:                 c.String("port"),
						BillerServiceBaseURL: c.String("biller-service-base-url"),
					}
					return engine.New(config)
				},
			},
			{
				Name:   "migrate",
				Usage:  "database migration",
				Flags:  flags,
				Action: cliaction.MigrateAction,
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
