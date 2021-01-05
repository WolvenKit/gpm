/*
 Copyright (c) 2020 - 2021 the WolvenKit contributors.

 Licensed under the GNU Affero General Public License v3.0 (the "License").

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package test

import (
	"fmt"
	"github.com/WolvenKit/gpm/internal/gpm/game"
	"github.com/WolvenKit/gpm/internal/gpm/mod"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// User flow
/*
   User downloads mod to $downloadPath
   User unpacks mod archive into $downloadPath/tmp/
   GPM reads mods manifest in $downloadPath/tmp/$archiveName/manifest.toml
   User install mod from $downloadPath/tmp/$archiveName/**
   User enables mod
   User disables mod
   User uninstalls mod
*/

// Test downloads go to desired directory
func TestDownloadMod(t *testing.T) {
	tmp := createSandbox()
	defer os.RemoveAll(tmp)

	logger := initLogging()

	m := mod.InitMod(logger)

	// Mock user CLI input
	i := new(mod.DownloadInput)
	i.Url = "https://cybermods.net/package/download/osulli/BraindanceProtocol/0.4.0/"
	i.Identifier = "braindance_protocol"
	i.FileType = ""

	// Download Mod
	err := m.Download(logger, tmp, i)
	if err != nil {
		panic(err)
	}

	// Assert the mod archive was downloaded into the correct path
	assert.FileExists(t, m.Directories.ArchivePath)
}

// Tests mod Configuration is read correctly
func TestReadModConfiguration(t *testing.T) {
	logger := initLogging()

	m := mod.InitMod(logger)
	// Simulate an already installed mod with a set InstallDirectory
	m.Directories.InstallDirectory = "mocks/example_cet_mod"
	m.ReadModConfiguration(logger, m.Directories.InstallDirectory)

	assert.Equal(t, mod.ModDirectories{
		InstallDirectory:   "mocks/example_cet_mod",
		ArchivePath:        "",
		TemporaryDirectory: "",
	}, m.Directories)
	assert.Equal(t, "WolvenKit", m.Creator)
	assert.Equal(t, "braindance_protocol", m.Identifier)
	assert.Equal(t, "0.0.0", m.Version)
	assert.Equal(t, "Braindance Protocol", m.DisplayName)
	assert.Equal(t, "A collection of LUA scripts to modify your Cyberpunk 2077 experience", m.Description)
	assert.Equal(t, "GNU v3", m.License)
	assert.Equal(t, "https://github.com/WolvenKit/BraindanceProtocol/", m.WebsiteURL)
	assert.Equal(t, []string(nil), m.Dependencies)
	assert.Equal(t, []string(nil), m.Tags)
	assert.Equal(t, []game.InstallStrategy([]game.InstallStrategy(nil)), m.InstallStrategies)
	assert.Equal(t, []string(nil), m.ExtraData)
}

// Tests mod Install follows install strategy
func TestInstallMod(t *testing.T) {
	tmp := createSandbox()
	err := os.MkdirAll(fmt.Sprintf("%s/Games/Cyberpunk 2077/bin/x64/plugins/cyber_engine_tweaks/mods/", tmp), 0777)
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(tmp)

	logger := initLogging()

	g := game.InitGame(logger)
	g.Directories.GameRoot = fmt.Sprintf("%s/Games/Cyberpunk 2077", tmp)
	m := mod.InitMod(logger)

	// Simulate a downloaded mod archive, with the manifest extracted
	m.Directories.TemporaryDirectory = tmp
	m.Directories.ArchivePath = "mocks/example_cet_mod.rar"

	m.ReadModConfiguration(logger, m.Directories.TemporaryDirectory)

	err = m.Install(logger, g)
	if err != nil {
		panic(err)
	}
	assert.DirExists(t, m.Directories.InstallDirectory)

	assert.FileExists(t, fmt.Sprintf("%s/mods/%s/init.lua", m.Directories.InstallDirectory, m.Identifier))
}

// Ensure scenario where mod manifest has missing keys is handled
func TestModManifestMissingKeys(t *testing.T) {
	// Checks CET not existing handled etc.
}

// Ensure scenario where mod archive cannot be found is handled
func TestUnarchiveMissingArchive(t *testing.T) {
	// Checks CET not existing handled etc.
}

// Ensure scenario is handled if required directories do not exist
func TestInstallToInvalidDirectory(t *testing.T) {
	// Checks CET not existing handled etc.
}

func TestUninstallMod(t *testing.T) {
}

func TestEnableMod(t *testing.T) {
}

func TestDisableMod(t *testing.T) {
}
