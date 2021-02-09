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
		cli.Command{
			Name:      "dispense",
			Usage:     "dispense a single programming fortune cookies",
			UsageText: "wisdom dispense",
			Action:    DispenseWisdom,
		},
		cli.Command{
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
	fmt.Println("TODO: dispense some wisdom")
	return nil
}

func DispenseChoice(c *cli.Context) error {
	fmt.Println("TODO: dispense some choice")
	return nil
}
