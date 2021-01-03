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
	"github.com/WolvenKit/gpm/cmd"
	"github.com/WolvenKit/gpm/internal/gpm/mod"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test downloads go to desired directory
func TestDownloadMod(t *testing.T) {
	tmp := createSandbox(false)
	logger := initLogging()

	// Mock user CLI input
	i := new(cmd.Input)
	i.Url = "https://cybermods.net/package/download/osulli/BraindanceProtocol/0.4.0/"
	i.Identifier = "braindance_protocol"
	i.FileType = ".zip"

	_, archivePath := cmd.DownloadMod(logger, tmp, i)

	assert.FileExists(t, archivePath)
}

// Tests mod Configuration is read correctly
func TestReadModConfiguration(t *testing.T) {
	logger := initLogging()

	m := mod.InitMod(logger, "mocks/example_cet_mod")
	m.ReadModConfiguration()
	logger.Debug(m.Creator)

	assert.Equal(t, "WolvenKit", m.Creator)
	assert.Equal(t, "braindance-protocol", m.Identifier)
	assert.Equal(t, "0.0.0", m.Version)
	assert.Equal(t, "Braindance Protocol", m.DisplayName)
	assert.Equal(t, "A collection of LUA scripts to modify your Cyberpunk 2077 experience", m.Description)
	assert.Equal(t, "GNU v3", m.License)
	assert.Equal(t, "https://github.com/WolvenKit/BraindanceProtocol/", m.WebsiteURL)
	assert.Equal(t, []string{""}, m.Dependencies)
	assert.Equal(t, []string{""}, m.Tags)
	assert.Equal(t, []string{"CET"}, m.InstallStrategies)
	assert.Equal(t, []string{""}, m.ExtraData)
}

//// Tests mod Install follows install strategy
//func TestInstallMod(t *testing.T){
//  tmp := createSandbox(false)
//  logger := initLogging()
//
//  // commands.GetConfiguration("cyberpunk", "cybermods", "braindance-protocol", "0.0.0")
//  mod := cmd.InstallMod(logger, archivePath, tmp, identifier)
//  assert.DirExists(t, mod)
//}

// Ensure scenario where mod archive cannot be found is handled
func UnarchiveMissingArchive() {
	// Checks CET not existing handled etc.
}

// Ensure scenario is handled if required directories do not exist
func InstallToInvalidDirectory() {
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
