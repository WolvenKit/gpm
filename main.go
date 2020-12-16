/*
 Copyright (C) 2020 the WolvenKit contributors.

 Licensed under the GNU Affero General Public License v3.0 (the "License").
*/

package main

import (
	"github.com/gruntwork-io/gruntwork-cli/entrypoint"

	"github.com/WolvenKit/gpm/internal/gpm/cli"
)

func main() {
	app := cli.InitCLI()

	// Use the entrypoint package, which takes care of exit codes, stack traces, and panics
	entrypoint.RunApp(app)
}
