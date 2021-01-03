/*
 Copyright (c) 2020 - 2021 the WolvenKit contributors.

 Licensed under the GNU Affero General Public License v3.0 (the "License").

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download the specified mod",
	Run: func(cmd *cobra.Command, args []string) {
		//DownloadMod("","","","")
	},
}

// Downloads mod from the Mod Registry
func DownloadMod(logger *zap.SugaredLogger, downloadDir string, input *Input) (error, string) {
	logger.Debugf("Downloading %s%s from %s", input.Identifier, input.FileType, input.Url)
	response, err := http.Get(input.Url)
	if err != nil {
		return err, ""
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		logger.Errorf("Received response code %s", response.StatusCode)
	}

	p := filepath.FromSlash(fmt.Sprintf("%s/%s%s", downloadDir, input.Identifier, input.FileType))
	logger.Debugf("Saving archive to %s", p)
	file, err := os.Create(p)
	if err != nil {
		return err, ""
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err, ""
	}
	logger.Debugf("Archive saved at %s", file.Name())

	return nil, p
}
