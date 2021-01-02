/*
 Copyright (c) 2020 - 2021 the WolvenKit contributors.

 Licensed under the GNU Affero General Public License v3.0 (the "License").
*/

package game

type Game struct {
    Identifier string
    DisplayName string
    Versions []string
    InstallStrategies []InstallStrategies
    ModRegistries []ModRegistries
}
