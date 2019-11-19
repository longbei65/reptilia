package version

import (
	"fmt"
	"runtime"
)

var (
	// AppName is the name of server.
	AppName string
	// GitCommit is the git commit that was compiled.
	GitCommit string
	// Version is a tag of project release.
	Version string
	// BuildDate is the ISO 8601 day that was built.
	BuildDate string
)

// PrintCliVersion print server info.
func PrintCliVersion() {
	fmt.Printf("%s\nversion: %s\ngit commit: %s\ngo version: %s\nbuild on: %s\n",
		AppName,
		Version,
		GitCommit,
		runtime.Version(),
		BuildDate,
	)
}
