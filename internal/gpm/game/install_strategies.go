/*
 Copyright (c) 2020 - 2021 the WolvenKit contributors.

 Licensed under the GNU Affero General Public License v3.0 (the "License").
*/

package game

type InstallStrategies struct {
    Identifier string
    DisplayName string
    InstallationPath string
}

const (
    zip = ".zip"
    rar = ".rar"
)
