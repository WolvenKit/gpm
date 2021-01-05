/*
 Copyright (c) 2020 - 2021 the WolvenKit contributors.

 Licensed under the GNU Affero General Public License v3.0 (the "License").

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package test

import (
	"go.uber.org/zap"
	"io/ioutil"
	"log"
	"path/filepath"
)

func initLogging() *zap.SugaredLogger {
	logger, _ := zap.NewDevelopment()
	defer func() {
		err := logger.Sync()
		if err != nil {
			panic(err)
		}
	}() // flushes buffer, if any
	sugar := logger.Sugar()
	return sugar
}

func createSandbox() string {
	dir, err := ioutil.TempDir("", "test-")
	if err != nil {
		log.Fatal(err)
	}
	dir = filepath.FromSlash(dir)

	return dir
}
