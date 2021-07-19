package golib

import (
	"errors"

	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:      "golib",
	Usage:     "generates new GO library",
	ArgsUsage: "LIBNAME [DIR]",
	Action:    golib,
	Category:  "new",
}

type library struct {
	Name string
}

var (
	l library
)

func golib(c *cli.Context) error {
	if !c.Args().Present() {
		cli.ShowSubcommandHelp(c)
		return errors.New("error: missing argument LIBNAME")
	}

	l.Name = c.Args().First() // parse project name

	return nil
}
