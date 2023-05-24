// 2023 FMNX team.
// Use of this code is governed by GNU General Public License.
// Additional information can be found on official web page: https://fmnx.su/
// Contact email: help@fmnx.su

package pacman

import (
	"io"
	"os"
	"os/exec"
	"strings"
)

// Optional flags for pacman remove command.
type RemoveOptions struct {
	// Run with sudo priveleges. [sudo]
	Sudo bool
	// Do not ask for any confirmation. [--noconfirm]
	NoConfirm bool
	// Remove with all unnecessary packages. [--recursive]
	Recursive bool
	// Remove with all explicitly installed deps. [-ss]
	ForceRecursive bool
	// Remove configuration files aswell. [--nosave]
	WithConfigs bool
	// Where command will write output text.
	Stdout io.Writer
	// Where command will write output text.
	Stderr io.Writer
	// Input from user is command will ask for something.
	Input io.Reader
	// Additional parameters, that will be appended to command as arguements.
	AdditionalParams []string
}

var RemoveDefault = RemoveOptions{
	Sudo:        true,
	NoConfirm:   true,
	Recursive:   true,
	WithConfigs: true,
	Stdout:      os.Stdout,
	Stderr:      os.Stderr,
	Input:       os.Stdin,
}

func Remove(pkgs string, opts ...RemoveOptions) error {
	if opts == nil {
		opts = []RemoveOptions{RemoveDefault}
	}
	o := opts[0]
	command := ""
	if o.Sudo {
		command += "sudo "
	}
	command += "pacman -R "

	if o.NoConfirm {
		command += "--noconfirm "
	}
	if o.Recursive {
		command += "--recursive "
	}
	if o.ForceRecursive {
		command += "-ss"
	}
	if o.WithConfigs {
		command += "--nosave"
	}

	command += strings.Join(o.AdditionalParams, " ") + " " + pkgs

	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = o.Stdout
	cmd.Stderr = o.Stderr
	cmd.Stdin = o.Input
	return cmd.Run()
}

func RemoveList(pkgs []string, opts ...RemoveOptions) error {
	return Remove(strings.Join(pkgs, " "), opts...)
}
