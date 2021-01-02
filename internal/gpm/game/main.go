/*
 Copyright (c) 2020 - 2021 the WolvenKit contributors.

 Licensed under the GNU Affero General Public License v3.0 (the "License").
*/

package game

type Game struct {
    GameName string
}

type InstallStrategies struct {
    Name string
    ModRegistries []ModRegistries
}

const (
    zip = ".zip"
    rar = ".rar"
)

type ModRegistries struct {
    Name string
    WebsiteURL string
    BaseURL string
    FileTypes []string
}
