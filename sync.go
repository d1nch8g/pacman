// 2023 FMNX team.
// Use of this code is governed by GNU General Public License.
// Additional information can be found on official web page: https://fmnx.su/
// Contact email: help@fmnx.su

package pacman

// Default options for pacman sync.
var (
	// Run with sudo priveleges
	Sudo = true
	// Do not reinstall up to date packages.
	Needed = true
	// Do not ask for any confirmation.
	NoConfirm = true
	// Do not show a progress bar when downloading files.
	NoProgressBar = true
	// Do not execute the install scriptlet if one exists.
	NoScriptlet = false
	// Install packages as non-explicitly installed.
	AsDeps = false
)

// Optional flags for pacman sync command.
type SyncOptions struct {
	// Load information from remote repositories.
	Refresh bool
	//
}

// Executes pacman sync command. If you provide options, they will override
// global flags.
func Sync(pkgs []string, opts ...SyncOptions) {

}
