// 2023 FMNX team.
// Use of this code is governed by GNU General Public License.
// Additional information can be found on official web page: https://fmnx.su/
// Contact email: help@fmnx.su

package pacman

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func init() {
	cmd := exec.Command("fakeroot")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Unable to set fakeroot environment")
		os.Exit(1)
	}
	fmt.Println("[fakeroot] - done")
}

func TestInstallNano(t *testing.T) {

}
