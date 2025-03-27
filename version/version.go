package version

import (
	"fmt"
	"runtime"
)

var appName = "ha-server"

func GetAppName() string {
	return appName
}

var (
	gitMajor     string // major version, always numeric
	gitMinor     string // minor version, numeric possibly followed by "+"
	gitVersion   string // semantic version, derived by build scripts
	gitCommit    string // sha1 from git, output of $(git rev-parse HEAD)
	gitTreeState string // state of git tree, either "clean" or "dirty"
	buildDate    string // build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
)

// Info exposes information about the version used for the current running code.
type Info struct {
	Major        string `json:"major,omitzero"`
	Minor        string `json:"minor,omitzero"`
	GitVersion   string `json:"gitVersion,omitzero"`
	GitCommit    string `json:"gitCommit,omitzero"`
	GitTreeState string `json:"gitTreeState,omitzero"`
	BuildDate    string `json:"buildDate,omitzero"`
	GoVersion    string `json:"goVersion,omitzero"`
	Compiler     string `json:"compiler,omitzero"`
	Platform     string `json:"platform,omitzero"`
}

// Get returns an Info object with all the information about the current running code.
func Get() Info {
	return Info{
		Major:        gitMajor,
		Minor:        gitMinor,
		GitVersion:   gitVersion,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

// String returns info as a human-friendly version string.
func (info Info) String() string {
	return info.GitVersion
}
