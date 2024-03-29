package goproject

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/neox5/go-picard/internal/info"
	"github.com/neox5/go-picard/internal/util/fileutil"
	"github.com/neox5/go-picard/pgk/scheduler"
	"github.com/urfave/cli/v2"
)

// goproject command definition
var Command = &cli.Command{
	Name:      "goproject",
	Usage:     "generates new GO project",
	ArgsUsage: "PROJECTNAME [DIR]",
	Action:    goproject,
	Category:  "new",
}

type project struct {
	Author  *info.Author
	Name    string
	Created string
	Year    string
}

var (
	p project
)

func goproject(c *cli.Context) error {
	if !c.Args().Present() {
		cli.ShowSubcommandHelp(c)
		return errors.New("error: missing argument PROJECTNAME")
	}

	p.Author = info.Neox5()
	p.Name = c.Args().First() // parse project name
	p.Created = time.Now().Format("Mon Jan 1 2006")
	p.Year = time.Now().Format("2006")

	s := scheduler.Scheduler{
		createFiles,
		initGoProject,
		goModTidy,
	}

	return s.Run()
}

func createFiles() error {
	files := []fileutil.FileTemplate{
		{Name: p.Name + "/cmd/" + p.Name + "/main.go", Tmpl: mainGoTmpl, Data: p.Name},
		{Name: p.Name + "/doc/.gitkeep", Tmpl: "", Data: nil},
		{Name: p.Name + "/internal/.gitkeep", Tmpl: "", Data: nil},
		{Name: p.Name + "/pkg/.gitkeep", Tmpl: "", Data: nil},
		{Name: p.Name + "/makefile", Tmpl: makefileTmpl, Data: nil},
		{Name: p.Name + "/LICENSE", Tmpl: licenseTmpl, Data: p},
		{Name: p.Name + "/.gitignore", Tmpl: gitignoreTmpl, Data: nil},
	}

	fmt.Println("GO project folder created!")
	return fileutil.CreateMultiple(files)
}

func initGoProject() error {
	return runCommand(p.Name, "go", "mod", "init", "github.com/neox5/"+p.Name)
}

func goModTidy() error {
	return runCommand(p.Name, "go", "mod", "tidy")
}

func runCommand(dir, app string, args ...string) error {
	cmd := exec.Command(app, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	path := dir
	if !filepath.IsAbs(dir) {
		path = "./" + path
	}
	fmt.Printf("%s: %s %s\n", path, app, strings.Join(args, " "))
	return cmd.Run()
}
