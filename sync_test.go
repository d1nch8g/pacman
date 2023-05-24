// 2023 FMNX team.
// Use of this code is governed by GNU General Public License.
// Additional information can be found on official web page: https://fmnx.su/
// Contact email: help@fmnx.su

package pacman

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInstallNano(t *testing.T) {
	assert := assert.New(t)

	err := Sync("nano")
	assert.NoError(err)

	err = Sync("git", SyncOptions{
		Sudo:                 true,
		Needed:               true,
		NoConfirm:            true,
		NoProgressBar:        true,
		NoScriptlet:          true,
		AsExplict:            true,
		Refresh:              true,
		Upgrade:              true,
		Output:               os.Stdout,
		AdditionalParameters: []string{},
	})
	assert.NoError(err)
}
