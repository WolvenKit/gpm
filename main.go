/*
 Copyright (c) 2020 - 2021 the WolvenKit contributors.

 Licensed under the GNU Affero General Public License v3.0 (the "License").
*/

package main

import (
	"github.com/urfave/cli/v2"
	"os"
	"time"
)

const VERSION = "0.0.0"

var AUTHORS = []*cli.Author{
	{Name: "osulli", Email: "https://github.com/osulli"},
}

var COMMANDS = []*cli.Command{
	versionCommand(VERSION),
}

func main() {
	app := &cli.App{
		Name:                   "gpm",
		Version:                VERSION,
		HelpName:               "help",
		Usage:                  "A game agnostic mod manager",
		Authors:                AUTHORS,
		Commands:               COMMANDS,
		Compiled:               time.Time{},
		Copyright:              "Copyright (c) 2020 - 2021 the WolvenKit contributors.\n   Licensed under the GNU Affero General Public License v3.0 (the \"License\").",
		UseShortOptionHandling: true,
		EnableBashCompletion:   true,
		HideHelp:               false,
		Action:                 nil,
		CommandNotFound:        nil,
		OnUsageError:           nil,
	}
	app.Run(os.Args)
}

func versionCommand(version string) *cli.Command {
	command := cli.Command{
		Name:     "Version",
		Aliases:  []string{"v", "--version", "-v", "version"},
		Usage:    "Print the application version",
		Category: "help",
		Action: func(context *cli.Context) error {
			println(VERSION)
			return nil
		},
	}

	return &command
}
