/*
 Copyright (c) 2020 - 2021 the WolvenKit contributors.

 Licensed under the GNU Affero General Public License v3.0 (the "License").
*/

package commands

import "github.com/urfave/cli/v2"

func downloadCommand() *cli.Command {
    command := cli.Command{
        Name:     "Download",
        Aliases:  []string{"d", "--download", "-d"},
        Usage:    "Download the specified mod",
        Category: "Download",
        Action: func(context *cli.Context) error {
            //DownloadMod()
            return nil
        },
    }

    return &command
}

func DownloadMod(game string, registry string, identifier string, version string ){
    // Downloads mod from the Mod Registry
}
