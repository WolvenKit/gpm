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
    "github.com/spf13/viper"
    "go.uber.org/zap"
)

// Read the mod's manifest.toml
func (m *Mod) ReadModConfiguration(logger *zap.SugaredLogger)  {
    logger.Debugf("Processing manifest.toml in %s", m.Directories.InstallDirectory)

    viper.AddConfigPath(m.Directories.InstallDirectory)
    viper.SetConfigName("manifest")
    viper.SetConfigType("toml")

    m.Creator = viper.GetString("creator")
    m.Identifier = viper.GetString("identifier")
    m.Version = viper.GetString("version")
    m.DisplayName = viper.GetString("display_name")
    m.Description = viper.GetString("description")
    m.License = viper.GetString("license")
    m.WebsiteURL = viper.GetString("website_url")
    m.Dependencies = viper.GetStringSlice("dependencies")
    m.Tags = viper.GetStringSlice("tags")
    // TODO - not sure why this doesn't work
    //m.InstallStrategies = viper.Get("install_strategies")
    m.ExtraData = viper.GetStringSlice("extra_data")
}
