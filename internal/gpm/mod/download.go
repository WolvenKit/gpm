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
	"go.uber.org/zap"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// TODO: Deprecate FileType (Does archiver need a file extension?) (Can we get the file name with http.Get ?)
// gpm download https://mods.net/001 braindance-protocol zip
type DownloadInput struct {
	Url        string
	Identifier string
	FileType   string
}

func (m *Mod) Download(logger *zap.SugaredLogger, downloadDir string, input *DownloadInput) error {
	logger.Debugf("Downloading %s%s from %s", input.Identifier, input.FileType, input.Url)
	response, err := http.Get(input.Url)

	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		logger.Errorf("Received response code whilst downloading a mod: %d", response.StatusCode)
	}

	p := filepath.FromSlash(fmt.Sprintf("%s/%s%s", downloadDir, input.Identifier, input.FileType))
	logger.Debugf("Creating archive at %s", p)
	file, err := os.Create(p)
	if err != nil {
		logger.Errorf(err.Error())
		return err
	}
	defer file.Close()

	logger.Debugf("Saving archive data to %s", file.Name())
	_, err = io.Copy(file, response.Body)
	if err != nil {
		logger.Errorf(err.Error())
		return err
	}
	logger.Debugf("Archive data saved at %s", file.Name())

	m.Directories.ArchivePath = file.Name()

	return nil
}
