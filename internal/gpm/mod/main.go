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
    "github.com/spf13/viper"
    "go.uber.org/zap"
)

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
    InstallStrategies []string
    ExtraData []string
}

func InitMod(logger *zap.SugaredLogger, directory string) *Mod {
    viper.AddConfigPath(directory)
    viper.SetConfigName("manifest")
    viper.SetConfigType("toml")

    err := viper.ReadInConfig()
    if err != nil {
        panic(fmt.Errorf("Fatal error config file: %s \n", err))
    }

    m := new(Mod)
    return m
}

func (m *Mod) ReadModConfiguration()  {
    m.Creator = viper.GetString("creator")
    m.Identifier = viper.GetString("identifier")
    m.Version = viper.GetString("version")
    m.DisplayName = viper.GetString("display_name")
    m.Description = viper.GetString("description")
    m.License = viper.GetString("license")
    m.WebsiteURL = viper.GetString("website_url")
    m.Dependencies = viper.GetStringSlice("dependencies")
    m.Tags = viper.GetStringSlice("tags")
    m.InstallStrategies = viper.GetStringSlice("install_strategies")
    m.ExtraData = viper.GetStringSlice("extra_data")
}

