package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/pytyagi/wisdom/lib/api"
	"github.com/pytyagi/wisdom/lib/wisdom"
	"github.com/urfave/cli"
)

const (
	name              = "wisdom"
	usage             = "dispense some programming wisdom fortune cookies"
	dispenseName      = "dispense"
	dispenseUsage     = "dispense a single programming wisdom fortune cookie"
	dispenseUsageText = "wisdom dispense"
	serveName         = "serve"
	serveUsage        = "run an API server to dispense programming wisdom"
	serveUsageText    = "wisdom serve -e ENV -p PORT --host HOST --api-path API_PATH"
	envShortArg       = "e"
	envArg            = "env"
	envEnvVar         = "ENV"
	envUsage          = "environment (dev | test | stage | prod)"
	hostArg           = "host"
	hostEnvVar        = "HOST"
	hostUsage         = "host to listen on "
	portShortArg      = "p"
	portArg           = "port"
	portEnvVar        = "PORT"
	portUsage         = "port to listen on "
	apiPathArg        = "api-path"
	apiPathEnvVar     = "API_PATH"
	apiPathUsage      = "url path prefix for mounting API router"
	webDirArg         = "web-dir"
	webDirEnvVar      = "WEB_DIR"
	webDirUsage       = "specify path to local web assets (e.g. swagger UI)"
	quotesFile        = "quotes.json"
)

func main() {

	app := cli.NewApp()
	app.Name = "wisdom"
	app.Usage = "dispense some programming wisdom fortune cookies"
	app.UsageText = "wisdom dispense"
	app.Version = wisdom.Version

	app.Commands = []cli.Command{
		{
			Name:      "dispense",
			Usage:     "dispense a single programming fortune cookies",
			UsageText: "wisdom dispense",
			Action:    DispenseWisdom,
		},
		{
			Name:      "serve",
			Usage:     "Run an API Server",
			UsageText: "wisdom serve -e ENV -p PORT --host HOST --api-path API_PATH",
			Action:    RunServer,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "env,e",
					Value:  api.DefaultEnv,
					Usage:  "environment(dev|test|stage|prod)",
					EnvVar: "ENV",
				},
				cli.StringFlag{
					Name:   "host",
					Value:  api.DefaultHost,
					Usage:  "host to listen on",
					EnvVar: "HOST",
				},
				cli.IntFlag{
					Name:   "port,p",
					Value:  api.DefaultPort,
					Usage:  "port to listen on",
					EnvVar: "PORT",
				},
				cli.StringFlag{
					Name:   "api-path",
					Value:  api.DefaultAPIPath,
					Usage:  "url path prefix for mounting api router",
					EnvVar: "API_PATH",
				},
			},
		},
	}
	err := app.Run(os.Args)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

// DispenseWisdom is an action
func DispenseWisdom(c *cli.Context) error {

	rand.Seed(time.Now().UnixNano())
	d, err := wisdom.FromFile("quotes.json")
	if err != nil {
		return errors.Wrap(err, "wisdom.FromFile failed")
	}
	fmt.Println(d.Random())
	return nil
}

func RunServer(c *cli.Context) error {

	cfg := api.NewConfig(
		c.String(envArg),
		c.String(hostArg),
		c.Int(portArg),
		c.String(apiPathArg),
	)

	log.Println("Version ", wisdom.Version)
	log.Println("ENV ", cfg.Env)
	log.Println("HOST ", cfg.Host)
	log.Println("PORT ", cfg.Port)
	log.Println("API PATH", cfg.APIPath)

	fmt.Println("RUN SERVER")
	return nil
}
