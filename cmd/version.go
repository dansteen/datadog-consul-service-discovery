package cmd

// this file is lifted from nomad https://github.com/hashicorp/nomad/blob/master/version.go

import (
	"bytes"
	"fmt"
)

// GitCommit is the git commit that was compiled. This will be filled in by the compiler.
var GitCommit string

// GitDescribe is the git description that was compiled
var GitDescribe string

// Version is the main version number that is being run at the moment.
const Version = "0.4.0"

// VersionPrerelease is a pre-release marker for the version. If this is "" (empty string)
// then it means that it is a final release. Otherwise, this is a pre-release
// such as "dev" (in development), "beta", "rc1", etc.
const VersionPrerelease = ""

// GetVersionParts returns the version strings. Printing of the
// version should be used in conjunction with the PrettyVersion method.
func GetVersionParts() (rev, ver, rel string) {
	ver = Version
	rel = VersionPrerelease
	if GitDescribe != "" {
		ver = GitDescribe
		// Trim off a leading 'v', we append it anyways.
		if ver[0] == 'v' {
			ver = ver[1:]
		}
	}
	if GitDescribe == "" && rel == "" && VersionPrerelease != "" {
		rel = "dev"
	}

	return GitCommit, ver, rel
}

// PrettyVersion takes the version parts and formats it in a human readable
// string.
func PrettyVersion(revision, version, versionPrerelease string) string {
	var versionString bytes.Buffer

	fmt.Fprintf(&versionString, "Consuldog v%s", version)
	if versionPrerelease != "" {
		fmt.Fprintf(&versionString, "-%s", versionPrerelease)

		if revision != "" {
			fmt.Fprintf(&versionString, " (%s)", revision)
		}
	}

	return versionString.String()
}
