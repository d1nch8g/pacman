// 2023 FMNX team.
// Use of this code is governed by GNU General Public License.
// Additional information can be found on official web page: https://fmnx.su/
// Contact email: help@fmnx.su

package pacman

import (
	"bytes"
	"errors"
	"strings"
)

// Query parameters for pacman packages.
type QueryParameters struct {
	// List packages explicitly installed. [--explicit]
	Explicit bool
	// List packages installed as dependencies. [--deps]
	Deps bool
	// Query only for packages installed from official repositories. [--native]
	Native bool
	// Query for packages installed from other sources. [--foreign]
	Foreign bool
	// Unrequired packages (not a dependency for other one). [--unrequired]
	Unrequired bool
	// Additional queue parameters.
	AdditionalParams []string
}

var QueryDefault = QueryParameters{}

type PackageInfo struct {
	Name    string
	Version string
}

func Query(p *QueryParameters) ([]PackageInfo, error) {
	if p == nil {
		p = &QueryDefault
	}

	args := []string{"-Q"}
	if p.Explicit {
		args = append(args, "--explicit")
	}
	if p.Deps {
		args = append(args, "--deps")
	}
	if p.Native {
		args = append(args, "--native")
	}
	if p.Foreign {
		args = append(args, "--foreign")
	}
	if p.Unrequired {
		args = append(args, "--unrequired")
	}

	cmd := pacmanCmd(false, args...)

	var b bytes.Buffer
	cmd.Stdout = &b
	cmd.Stderr = &b

	err := cmd.Run()
	if err != nil {
		if b.String() == "" {
			return nil, nil
		}
		return nil, errors.New("unable to query packages: " + b.String())
	}
	return parsePackages(b.String()), nil
}

func parsePackages(q string) []PackageInfo {
	var rez []PackageInfo
	for _, v := range strings.Split(q, "\n") {
		if v == "" {
			break
		}
		splt := strings.Split(v, " ")
		rez = append(rez, PackageInfo{
			Name:    splt[0],
			Version: splt[1],
		})
	}
	return rez
}
