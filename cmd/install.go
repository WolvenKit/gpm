/*
 Copyright (c) 2020 - 2021 the WolvenKit contributors.

 Licensed under the GNU Affero General Public License v3.0 (the "License").
*/

package cmd

import (
    "fmt"
    "github.com/mholt/archiver"
    "path/filepath"
)

func InstallMod(archivePath string, installPath string, identifier string) string {
    archivePath = filepath.FromSlash(archivePath)
    installPath = filepath.FromSlash(installPath)

    unarchive(archivePath, installPath, identifier)

    mod := filepath.FromSlash(fmt.Sprintf("%s/%s", installPath, identifier))
    return mod
}


func unarchive(archivePath string, installPath string, identifier string) error{
    err := archiver.Extract(archivePath, fmt.Sprintf("mods/%s", identifier), installPath)
    if err != nil {
        return err
    }

    return nil
}
