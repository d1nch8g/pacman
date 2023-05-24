// 2023 FMNX team.
// Use of this code is governed by GNU General Public License.
// Additional information can be found on official web page: https://fmnx.su/
// Contact email: help@fmnx.su

package pacman

import (
	"io"
	"os"
	"strings"
)

// Optional flags for pacman sync command.
type SyncOptions struct {
	// Run with sudo priveleges. [sudo]
	Sudo bool
	// Do not reinstall up to date packages. [--needed]
	Needed bool
	// Do not ask for any confirmation. [--noconfirm]
	NoConfirm bool
	// Do not show a progress bar when downloading files. [--noprogressbar]
	NoProgressBar bool
	// Do not execute the install scriptlet if one exists. [--noscriptlet]
	NoScriptlet bool
	// Install packages as non-explicitly installed. [--asdeps]
	AsDeps bool
	// Install packages as explictly installed. [--asexplict]
	AsExplict bool
	// Download fresh package databases from the server. [--refresh]
	Refresh bool
	// Upgrade programms that are outdated. [--sysupgrade]
	Upgrade bool
	// Only download, but do not install package. [--downloadonly]
	DownloadOnly bool
	// Clean old packages from cache directory. [--clean]
	Clean bool
	// Clean all packages from cache directory. [-cc]
	CleanAll bool
	// Where command will write output text.
	Output io.Writer
	// Additional parameters, that will be appended to command as arguements.
	AdditionalParameters []string
}

// Those are options that will be set up by default on program execution.
var DefaultOptions = SyncOptions{
	Sudo:          true,
	Needed:        true,
	NoConfirm:     true,
	NoProgressBar: true,
	Output:        os.Stdout,
}

// Executes pacman sync command. This command will read sync options and form
// command based on first elements from the array.
func Sync(pkgs []string, opts ...SyncOptions) {
	if opts == nil {
		opts = []SyncOptions{DefaultOptions}
	}
	o := opts[0]
	command := ""
	if o.Sudo {
		command += "sudo "
	}
	command += "pacman -S "
	if o.Needed {
		command += "--needed "
	}
	if o.NoConfirm {
		command += "--noconfirm "
	}
	if o.NoProgressBar {
		command += "--noprogressbar "
	}
	if o.NoScriptlet {
		command += "--noscriptlet "
	}
	if o.AsDeps {
		command += "--asdeps "
	}
	if o.AsExplict {
		command += "--asexplict "
	}
	if o.Refresh {
		command += "--refresh "
	}
	if o.Upgrade {
		command += "--sysupgrade "
	}
	if o.DownloadOnly {
		command += "--downloadonly"
	}
	if o.DownloadOnly {
		command += "--downloadonly"
	}
	if o.Clean {
		command += "--clean"
	}
	if o.CleanAll {
		command += "-cc"
	}
	command += strings.Join(o.AdditionalParameters, " ") + " "
	command += strings.Join(pkgs, " ")

}
