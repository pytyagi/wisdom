package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pytyagi/wisdom/lib/wisdom"
	"github.com/urfave/cli"
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
			Name:      "choice",
			Usage:     "dispense a single choice fortune cookies",
			UsageText: "wisdom dispense",
			Action:    DispenseWisdom,
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
	q := wisdom.NewQuote("Go is growing on me", "Piyush Tyagi")
	fmt.Println(q)
	return nil
}
