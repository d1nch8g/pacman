// 2023 FMNX team.
// Use of this code is governed by GNU General Public License.
// Additional information can be found on official web page: https://fmnx.su/
// Contact email: help@fmnx.su

package pacman

import (
	"bytes"
	"errors"
	"os/exec"
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
	command := "pacman -Q "
	if p.Explicit {
		command += "--explicit "
	}
	if p.Deps {
		command += "--deps "
	}
	if p.Native {
		command += "--native "
	}
	if p.Foreign {
		command += "--foreign "
	}
	if p.Unrequired {
		command += "--unrequired "
	}

	command += strings.Join(p.AdditionalParams, " ")

	var b bytes.Buffer
	cmd := exec.Command("bash", "-c", command)
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
