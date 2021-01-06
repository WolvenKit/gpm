/*
 Copyright (c) 2020 - 2021 the WolvenKit contributors.

 Licensed under the GNU Affero General Public License v3.0 (the "License").

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package command

import (
	"fmt"
	"github.com/WolvenKit/gpm/internal/gpm/config"
	"github.com/spf13/cobra"
	"os"

	"github.com/spf13/viper"
)

var (
	cfgFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gpm",
	Short: "A game agnostic mod manager",
	Long: `A Fast(ish) and Flexible, game agnostic, mod manager built with
                love by osulli and WolvenKit Devs in Go.

                Source available at https://github.com/WolvenKit/gpm

                Copyright (c) 2020 - 2021 the WolvenKit contributors.
                Licensed under the GNU Affero General Public License v3.0 (the "License").`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by gpm/main.go.main()
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Define flags and configuration settings.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gpm/.gpm.toml)")

	// Define sub-commands
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(downloadCmd)
	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(describeCmd)
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(completionCmd)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		config.InitialiseConfig()
	}
}
