/*
 Copyright (c) 2020 - 2021 the WolvenKit contributors.

 Licensed under the GNU Affero General Public License v3.0 (the "License").
*/

package commands

import (
    "path/filepath"
    "github.com/mholt/archiver"
)

func InstallMod(archivePath string, installPath string) {
    archivePath = filepath.FromSlash(archivePath)
    installPath = filepath.FromSlash(installPath)

    unarchive(archivePath, installPath)
}


func unarchive(archivePath string, installPath string) error{
    err := archiver.Unarchive(archivePath, installPath)
    if err != nil {
        return err
    }

    return nil
}
