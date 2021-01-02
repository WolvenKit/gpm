/*
 Copyright (c) 2020 - 2021 the WolvenKit contributors.

 Licensed under the GNU Affero General Public License v3.0 (the "License").
*/

package test

import (
    "fmt"
    "github.com/WolvenKit/gpm/internal/gpm/commands"
    "github.com/stretchr/testify/assert"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "testing"
)

func testSandbox() string{
    dir, err := ioutil.TempDir("", "test-")
    if err != nil {
        log.Fatal(err)
    }
    dir = filepath.FromSlash(fmt.Sprintf(dir))

    defer os.RemoveAll(dir)

    return dir
}


func TestDownloadMod(t *testing.T) {
    tmp := testSandbox()

    // TODO - read from config and build URL using /gpm/game package
    //commands.GetConfiguration("cyberpunk", "cybermods", "braindance-protocol", "0.0.0")
    url := "https://cybermods.net/package/download/osulli/BraindanceProtocol/0.4.0/"
    identifier := "braindance_protocol"
    fileType := ".zip"


    _, archivePath := commands.DownloadMod(url, tmp, identifier, fileType)
    mod := commands.InstallMod(archivePath, tmp, identifier)

    assert.DirExists(t, mod)
}

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
