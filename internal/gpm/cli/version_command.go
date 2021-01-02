/*
 Copyright (c) 2020 - 2021 the WolvenKit contributors.

 Licensed under the GNU Affero General Public License v3.0 (the "License").
*/

package cli

import (
	"github.com/urfave/cli/v2"
)

// https://pkg.go.dev/github.com/urfave/cli#Command

func versionCommand() *cli.Command {
	command := cli.Command{
		Name:     "Version",
		Aliases:  []string{"v", "--version", "-v"},
		Usage:    "Print the application version",
		Category: NAME,
		Action: func(context *cli.Context) error {
			println(VERSION)
			return nil
		},
	}

	return &command
}
