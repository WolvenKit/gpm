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
	"github.com/mholt/archiver"
	"go.uber.org/zap"
	"path/filepath"
)

func InstallMod(logger *zap.SugaredLogger, archivePath string, installPath string, identifier string) string {
	archivePath = filepath.FromSlash(archivePath)
	installPath = filepath.FromSlash(installPath)

	logger.Debugf("Unarchiving mods/%s directory from %s to %s", identifier, archivePath, installPath)
	unarchive(archivePath, installPath, identifier)

	mod := filepath.FromSlash(fmt.Sprintf("%s/%s", installPath, identifier))
	return mod
}

func unarchive(archivePath string, installPath string, identifier string) error {
	err := archiver.Extract(archivePath, fmt.Sprintf("mods/%s", identifier), installPath)
	if err != nil {
		return err
	}

	return nil
}
