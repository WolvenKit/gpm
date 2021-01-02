/*
 Copyright (c) 2020 - 2021 the WolvenKit contributors.

 Licensed under the GNU Affero General Public License v3.0 (the "License").
*/

package cli

import "github.com/urfave/cli/v2"

func downloadCommand() *cli.Command {
    command := cli.Command{
        Name:     "Download",
        Aliases:  []string{"d", "--download", "-d"},
        Usage:    "Download the specified mod",
        Category: NAME,
        Action: func(context *cli.Context) error {
            println(VERSION)
            return nil
        },
    }

    return &command
}

func downloadMod(){
    // Downloads mod from the Mod Registry
}
