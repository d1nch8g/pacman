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

const (
	pacman  = `pacman`
	sudo    = `sudo`
	makepkg = `makepkg`
)

func init() {
	_, err := exec.LookPath(pacman)
	if err != nil {
		fmt.Println("unable to find pacman in system")
		os.Exit(1)
	}
	_, err = exec.LookPath(sudo)
	if err != nil {
		fmt.Println("unable to find sudo in system")
		os.Exit(1)
	}
	_, err = exec.LookPath(makepkg)
	if err != nil {
		fmt.Println("unable to find makepkg in system")
		os.Exit(1)
	}
}

func pacmanCmd(sudo bool, args ...string) *exec.Cmd {
	if sudo {
		args := append([]string{pacman}, args...)
		return exec.Command("sudo", args...)
	}
	return exec.Command(pacman, args...)
}

func formOptions[Options any](arr []Options, dv *Options) *Options {
	if len(arr) != 1 {
		return dv
	}
	return &arr[1]
}
