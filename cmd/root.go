/*
 Copyright (c) 2020 - 2021 the WolvenKit contributors.

 Licensed under the GNU Affero General Public License v3.0 (the "License").
*/

package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "os"
)

var rootCmd = &cobra.Command{
    Use:   "gpm",
    Short: "A game agnostic mod manager",
    Long: `A Fast(ish) and Flexible, game agnostic, mod manager built with
                love by osulli and WolvenKit Devs in Go.

                Source available at https://github.com/WolvenKit/gpm

                Copyright (c) 2020 - 2021 the WolvenKit contributors.
                Licensed under the GNU Affero General Public License v3.0 (the "License").`,
    Run: func(cmd *cobra.Command, args []string) {
        // Do Stuff Here
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}

/*

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
*/
