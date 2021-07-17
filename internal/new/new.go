package new

import (
	"github.com/neox5/go-picard/internal/new/goproject"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:  "new",
	Usage: "create new stuff",
	Subcommands: []*cli.Command{
		goproject.Command,
	},
}
