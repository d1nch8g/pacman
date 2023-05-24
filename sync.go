// 2023 FMNX team.
// Use of this code is governed by GNU General Public License.
// Additional information can be found on official web page: https://fmnx.su/
// Contact email: help@fmnx.su

package pacman

// Default options for pacman sync. [--sync]
var (
	// Run with sudo priveleges. [sudo]
	Sudo = true
	// Do not reinstall up to date packages. [--needed]
	Needed = true
	// Do not ask for any confirmation. [--noconfirm]
	NoConfirm = true
	// Do not show a progress bar when downloading files. [--noprogressbar]
	NoProgressBar = true
	// Do not execute the install scriptlet if one exists. [--noscriptlet]
	NoScriptlet = false
	// Install packages as non-explicitly installed. [--asdeps]
	AsDeps = false
	// Install packages as explictly installed. [--asexplict]
	AsExplict = false
)

// Optional flags for pacman sync command.
type SyncOptions struct {
	// Sync information from remote repositories. []
	Refresh bool
	// Update programms that are outdated. [--sysupgrade]
	Update bool
	// Only download, but do not install package. [-w]
	
}

// Executes pacman sync command. If you provide options, they will override
// global flags.
func Sync(pkgs []string, opts ...SyncOptions) {

}
