/*
 Copyright (c) 2020 - 2021 the WolvenKit contributors.

 Licensed under the GNU Affero General Public License v3.0 (the "License").

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package game

type InstallStrategy struct {
	Identifier       string
	DisplayName      string
	InstallationPath string
}

const (
	zip = ".zip"
	rar = ".rar"
)

var InstallStrategies = []InstallStrategy{
	{
		Identifier:       "CET",
		DisplayName:      "Cyberpunk Engine Tweaks",
		InstallationPath: "bin/x64/plugins/cyber_engine_tweaks/mods/",
	},
}
