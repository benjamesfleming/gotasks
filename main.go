package main

import (
	"log"
	"os"
	"sort"

	. "git.benfleming.nz/benfleming/gotasks/app/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	cliApp := &cli.App{
		HelpName:    "gotasks",
		Name:        "gotasks",
		Usage:       "get stuff done",
		Description: "A Task Tracker & Todo List",
		Version:     "v0.2.1",
		Flags:       []cli.Flag{},
		Commands: []*cli.Command{
			InstallCommand,
			StartCommand,
		},
	}

	sort.Sort(cli.FlagsByName(cliApp.Flags))
	sort.Sort(cli.CommandsByName(cliApp.Commands))

	err := cliApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
