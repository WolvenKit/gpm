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
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install [IDENTIFIER]",
	Short: "Install the specified mod",
	Args:  cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		InstallMod(args[0])
	},
}

func InstallMod(identifier string) {
	// Todo: Get mod data from gpm data file using the identifier
}
