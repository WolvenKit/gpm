/*
 Copyright (c) 2020 - 2021 the WolvenKit contributors.

 Licensed under the GNU Affero General Public License v3.0 (the "License").
*/

package cli

import (
	"github.com/gruntwork-io/gruntwork-cli/entrypoint"
	"github.com/urfave/cli/v2"
)

var NAME = "gpm"
var VERSION = "poc"

func InitCLI() *cli.App {
	app := entrypoint.NewApp()

	app.Name = NAME
	app.Authors = []*cli.Author{
		{
			Name: "osulli",
		},
	}

	app.Version = VERSION

	app.EnableBashCompletion = true

	app.Commands = []*cli.Command{versionCommand()}

	return app
}
