package main

import (
	_ "github.com/joho/godotenv/autoload"

	"log"
	"os"

	"github.com/mitchellh/cli"

	"github.com/dinngodev/furucombo-reward-scripts/commands"
)

const version = "0.1.0"

func main() {

	c := cli.NewCLI("rewarder", version)
	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"once": func() (cli.Command, error) {
			return &commands.OnceCommand{}, nil
		},
		"staking": func() (cli.Command, error) {
			return &commands.StakingCommand{}, nil
		},
		"stakingv2": func() (cli.Command, error) {
			return &commands.StakingV2Command{}, nil
		},
		"bonus": func() (cli.Command, error) {
			return &commands.BonusCommand{}, nil
		},
		"usage": func() (cli.Command, error) {
			return &commands.UsageCommand{}, nil
		},
		"polygon-usage": func() (cli.Command, error) {
			return &commands.PolygonUsageCommand{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
