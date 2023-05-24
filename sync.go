// 2023 FMNX team.
// Use of this code is governed by GNU General Public License.
// Additional information can be found on official web page: https://fmnx.su/
// Contact email: help@fmnx.su

package pacman

import (
	"bytes"
	"errors"
	"io"
	"os"
	"os/exec"
	"strings"
)

// Optional parameters for pacman sync command.
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
	Stdout io.Writer
	// Where command will write output text.
	Stderr io.Writer
	// Input from user is command will ask for something.
	Input io.Reader
	// Additional parameters, that will be appended to command as arguements.
	AdditionalParams []string
}

// Those are options that will be set up by default on program execution.
var SyncDefault = SyncOptions{
	Sudo:          true,
	Needed:        true,
	NoConfirm:     true,
	NoProgressBar: true,
	Stdout:        os.Stdout,
	Stderr:        os.Stderr,
}

// Executes pacman sync command. This command will read sync options and form
// command based on first elements from the array.
func Sync(pkgs string, opts ...SyncOptions) error {
	if opts == nil {
		opts = []SyncOptions{SyncDefault}
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
		command += "--asexplicit "
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

	command += strings.Join(o.AdditionalParams, " ") + " " + pkgs

	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = o.Stdout
	cmd.Stderr = o.Stderr
	cmd.Stdin = o.Input
	return cmd.Run()
}

// Sync command for package string list.
func SyncList(pkgs []string, opts ...SyncOptions) error {
	return Sync(strings.Join(pkgs, " "), opts...)
}

// Options to apply when searching for some package.
type SearchOptions struct {
	// Run with sudo priveleges. [sudo]
	Sudo bool
	// Download fresh package databases from the server. [--refresh]
	Refresh bool
	// Input from user is command will ask for something.
	Input io.Reader
}

// Structure to recieve from search result
type SearchResult struct {
	Repo    string
	Name    string
	Version string
	Desc    string
}

var SearchDefault = SearchOptions{
	Sudo:    true,
	Refresh: true,
	Input:   os.Stdin,
}

// Search for packages.
func Search(re string, opts ...SearchOptions) ([]SearchResult, error) {
	if opts == nil {
		opts = append(opts, SearchDefault)
	}
	o := opts[0]

	command := ""
	if o.Sudo {
		command += "sudo "
	}
	command += "pacman -Ss "
	if o.Refresh {
		command += "--refresh "
	}
	command += re

	cmd := exec.Command("bash", "-c", command)
	var b bytes.Buffer
	cmd.Stdout = &b
	cmd.Stderr = &b
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		if b.String() == `` {
			return nil, nil
		}
		return nil, errors.New("unable to search: " + b.String())
	}
	return serializeOutput(b.String()), nil
}

func serializeOutput(output string) []SearchResult {
	var rez []SearchResult
	lines := strings.Split(output, "\n")
	for i, line := range lines {
		if i%2 == 1 {
			continue
		}
		spl := strings.Split(line, " ")
		repoName := strings.Split(spl[0], "/")
		rez = append(rez, SearchResult{
			Repo:    repoName[0],
			Name:    repoName[1],
			Version: spl[1],
			Desc:    strings.Trim(lines[i+1], " "),
		})
	}
	return rez
}
