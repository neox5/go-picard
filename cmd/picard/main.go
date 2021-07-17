package main

import (
	"log"
	"os"

	"github.com/neox5/go-picard/internal/new"
	"github.com/neox5/go-picard/internal/say"
	"github.com/urfave/cli/v2"
)

func main() {
	// TODO: App versioning

	picard := &cli.App{}
	picard.EnableBashCompletion = true
	picard.UseShortOptionHandling = true

	// global flags
	picard.Flags = []cli.Flag{}

	// root commands
	picard.Commands = []*cli.Command{
		new.Command,
		say.Command,
	}

	err := picard.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
