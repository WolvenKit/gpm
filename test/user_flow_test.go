/*
 Copyright (c) 2020 - 2021 the WolvenKit contributors.

 Licensed under the GNU Affero General Public License v3.0 (the "License").
*/

package test

import (
    "github.com/WolvenKit/gpm/cmd"
    "github.com/stretchr/testify/assert"
    "testing"
)


func TestDownloadMod(t *testing.T) {
    tmp := createSandbox(false)
    logger := initLogging()

    url := "https://cybermods.net/package/download/osulli/BraindanceProtocol/0.4.0/"
    identifier := "braindance_protocol"
    fileType := ".zip"

    _, archivePath := cmd.DownloadMod(logger, url, tmp, identifier, fileType)

    assert.FileExists(t, archivePath)
}

//func TestInstallMod(t *testing.T){
//    tmp := createSandbox()
//
//    commands.GetConfiguration("cyberpunk", "cybermods", "braindance-protocol", "0.0.0")
//    mod := cmd.InstallMod(archivePath, tmp, identifier)
//    assert.DirExists(t, mod)
//}

func InstallNoDirectories(){
    // Checks CET not existing handled etc.
}

func InstallMod(t *testing.T) {
}

func UninstallMod(t *testing.T) {
}

func EnableMod(t *testing.T) {
}

func DisableMod(t *testing.T) {
}
