/*
 Copyright (c) 2020 - 2021 the WolvenKit contributors.

 Licensed under the GNU Affero General Public License v3.0 (the "License").
*/

package main

import (
    "github.com/urfave/cli/v2"
    "time"
)

func main() {
    app := &cli.App{
        Name:                   "gpm",
        Version:                "0.0.0",
        HelpName:               "",
        Usage:                  "",
        UsageText:              "",
        ArgsUsage:              "",
        Description:            "",
        Commands:               nil,
        Flags:                  nil,
        EnableBashCompletion:   false,
        HideHelp:               false,
        HideHelpCommand:        false,
        HideVersion:            false,
        BashComplete:           nil,
        Before:                 nil,
        After:                  nil,
        Action:                 nil,
        CommandNotFound:        nil,
        OnUsageError:           nil,
        Compiled:               time.Time{},
        Authors:                []*cli.Author{
            {
                Name: "osulli",
            },
        Copyright:              "",
        Reader:                 nil,
        Writer:                 nil,
        ErrWriter:              nil,
        ExitErrHandler:         nil,
        Metadata:               nil,
        ExtraInfo:              nil,
        CustomAppHelpTemplate:  "",
        UseShortOptionHandling: false,
    }

}
