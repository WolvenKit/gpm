/*
 Copyright (c) 2020 - 2021 the WolvenKit contributors.

 Licensed under the GNU Affero General Public License v3.0 (the "License").
*/

package game

type installStrategies struct {
	Identifier       string
	DisplayName      string
	InstallationPath string
}

const (
	zip = ".zip"
	rar = ".rar"
)

var InstallStrategies = []installStrategies{
	{
		Identifier:       "CET",
		DisplayName:      "Cyberpunk Engine Tweaks",
		InstallationPath: "bin/x64/plugins/cyber_engine_tweaks/mods/",
	},
}
