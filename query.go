// 2023 FMNX team.
// Use of this code is governed by GNU General Public License.
// Additional information can be found on official web page: https://fmnx.su/
// Contact email: help@fmnx.su

package pacman

// Query parameters for pacman packages.
type QueryParameters struct {
	// List packages explicitly installed. [--explicit]
	Explicit bool
	// List packages installed as dependencies. [--deps]
	Deps bool
	// Query only for programs installed from official repositories. [--]
}

type PackageInfo struct {
	Name    string
	Version string
}

func Query(p QueryParameters) {

}
