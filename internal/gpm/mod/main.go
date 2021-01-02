/*
 Copyright (c) 2020 - 2021 the WolvenKit contributors.

 Licensed under the GNU Affero General Public License v3.0 (the "License").
*/

package mod

import 	"github.com/WolvenKit/gpm/internal/gpm/game"

type Mod struct {
    Creator string
    Identifier string
    Version string
    DisplayName string
    Description string
    License string
    WebsiteURL string
    Dependencies []string
    Tags []string
    InstallStrategies []game.InstallStrategies
    ExtraData []string
}

func InitMod() Mod{
    m := Mod{}

    return m
}
