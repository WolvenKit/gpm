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
	"github.com/WolvenKit/gpm/internal/gpm/mod"
	"github.com/WolvenKit/gpm/internal/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var downloadCmd = &cobra.Command{
	Use:   "download [IDENTIFIER] [URL]",
	Short: "Download the specified mod",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		DownloadMod(args[0], args[1])
	},
}

func DownloadMod(identifier string, url string) {
	// TODO - toggle the 'game', not just CP77
	cpp := viper.GetString("cyberpunk_path")

	logger := log.GetLogger()
	m := mod.InitMod(logger)

	logger.Debug(viper.GetViper())

	i := new(mod.DownloadInput)
	i.Identifier = identifier
	i.Url = url

	err := m.Download(logger, cpp, i)
	if err != nil {
		logger.Fatal(err)
	}
}
