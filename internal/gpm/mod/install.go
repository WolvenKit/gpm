/*
 Copyright (c) 2020 - 2021 the WolvenKit contributors.

 Licensed under the GNU Affero General Public License v3.0 (the "License").

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package mod

import (
	"fmt"
	"github.com/WolvenKit/gpm/internal/gpm/game"
	"github.com/mholt/archiver"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"path/filepath"
)

func (m *Mod) Install(logger *zap.SugaredLogger, directories ModDirectories) error {
	// Extract the manifest.toml from the archive
	extractManifest(logger, m)

	// Process the extracted manifest.toml before installation
	viper.AddConfigPath(m.Directories.TemporaryDirectory)
	m.ReadModConfiguration(logger)
	logger.Debug(m)

	// Get the InstallPath from the InstallStrategy
	processInstallStrategy(logger, m)

	// Extract and Install the entire mod according to the manifest
	unarchiveMod(logger, m)

	return nil
}

// Extract the manifest.toml into the ArchivePath as manifest.toml
func extractManifest(logger *zap.SugaredLogger, m *Mod) error {
	d := filepath.FromSlash(fmt.Sprintf("%s/%s", m.Directories.ArchivePath, "manifest.toml"))
	err := archiver.Extract(m.Directories.ArchivePath, "manifest.toml", d)
	if err != nil {
		return err
	}

	return nil
}

// Extract the mods/ directory into the InstallDirectory
func unarchiveMod(logger *zap.SugaredLogger, m *Mod) error {
	p := filepath.FromSlash(fmt.Sprintf("mods/%s", m.Identifier))

	err := archiver.Extract(m.Directories.ArchivePath, p, m.Directories.InstallDirectory)
	if err != nil {
		return err
	}
	return nil
}

// Process the InstallStrategy defined in manifest.toml
func processInstallStrategy(logger *zap.SugaredLogger, m *Mod) error {
	/*
			   TODO - process multiple strategies once InstallStrategy doc completed
			   How will we handle multiple install strategies as proposed?
		        Review this entire function
	*/

	// Mock a strategy -- would normally be loaded from a central configuration
	is := new(game.InstallStrategy)
	is.Identifier = "CET"
	is.DisplayName = "Cyberpunk Engine Tweaks"
	is.InstallationPath = "bin/x64/plugins/cyber_engine_tweaks/mods/"

	// Set the mod's strategies to mocked strategy instead of what is in the mod's manifest.toml
	m.InstallStrategies = []game.InstallStrategy{*is}

	// Now process the strategy
	p := filepath.FromSlash(fmt.Sprintf("%s/%s/%s", viper.GetString("cyberpunk_path"), m.InstallStrategies[0].InstallationPath, m.Identifier))
	logger.Debugf("Install Directory set to: %s", p)
	m.Directories.InstallDirectory = p

	return nil
}
