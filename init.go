// 2023 FMNX team.
// Use of this code is governed by GNU General Public License.
// Additional information can be found on official web page: https://fmnx.su/
// Contact email: help@fmnx.su

package pacman

import (
	"fmt"
	"os"
	"os/exec"
)

func init() {
	_, err := exec.LookPath("pacman")
	if err != nil {
		fmt.Println("unable to find pacman in system")
		os.Exit(1)
	}
	_, err = exec.LookPath("bash")
	if err != nil {
		fmt.Println("unable to find bash in system")
		os.Exit(1)
	}
}
