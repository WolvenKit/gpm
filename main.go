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

func main() {
	app := &cli.App{
		Name:        "gpm",
		Version:     "0.0.0",
		HelpName:    "help",
		Description: "A game agnostic mod manager",
		Authors: []*cli.Author{
			{Name: "osulli"},
		},
		Commands:               nil,
		EnableBashCompletion:   true,
		HideHelp:               false,
		Action:                 nil,
		CommandNotFound:        nil,
		OnUsageError:           nil,
		Compiled:               time.Time{},
		Copyright:              "Copyright (c) 2020 - 2021 the WolvenKit contributors.\n   Licensed under the GNU Affero General Public License v3.0 (the \"License\").",
		UseShortOptionHandling: true,
	}
	app.Run(os.Args)
}
