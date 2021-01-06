/*
 Copyright (c) 2020 - 2021 the WolvenKit contributors.

 Licensed under the GNU Affero General Public License v3.0 (the "License").

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package config

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"os"
)

// InitialiseConfig reads in config file and ENV variables if set.
func InitialiseConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Look for env variables with a GPM prefix
	viper.SetEnvPrefix("GPM")

	// Look for config files with .gpm.*
	viper.SetConfigName(".gpm")

	// Look for the config files in these directories
	viper.AddConfigPath(fmt.Sprintf("%s/.gpm", home))
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")

	// Load environment variables that match PREFIX
	viper.AutomaticEnv()

	// Try to load in a config file (if found)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			// TODO - I suspect this error would produce an empty file name as empty string
			_, err = fmt.Fprintln(os.Stderr, "No config file found:", viper.ConfigFileUsed())
			if err != nil {
				panic(err)
			}
		} else {
			// Config file was found but another error was produced
			panic(err)
		}
	}
}
