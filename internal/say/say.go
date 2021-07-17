package say

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:  "say",
	Usage: "prints what picard should say",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "out",
			Aliases: []string{"o"},
			Value:   false,
			Usage:   "adds 'Picard out...' at the end",
		},
	},
	Action: say,
}

func say(c *cli.Context) error {
	fmt.Println(c.Args().Get(0))
	if c.Bool("out") {
		fmt.Println("Picard out...")
	}
	return nil
}
