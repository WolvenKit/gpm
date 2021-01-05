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
    "github.com/WolvenKit/gpm/internal/gpm/game"
    "go.uber.org/zap"
)

type Mod struct {
    Directories ModDirectories
    Creator          string
    Identifier string
    Version string
    DisplayName string
    Description string
    License string
    WebsiteURL string
    Dependencies []string
    Tags []string
    InstallStrategies []game.InstallStrategy
    ExtraData []string
}

// Key locations for the mod's files
type ModDirectories struct {
    InstallDirectory   string
    ArchivePath        string
    TemporaryDirectory string
}

func InitMod(logger *zap.SugaredLogger) *Mod {
    m := new(Mod)
    return m
}

// Updates central mod log
func (m *Mod) UpdateModLog()  {
    // TODO - establish convention / large manifest file thing that stores all mod data from `m *Mod`
}
