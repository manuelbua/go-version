package version

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const Unversioned = "(unversioned)"

// populated via -ldflags
var (
	shortVersion = Unversioned
	commitHash   = ""
	commitStamp  = ""
	buildUser    = ""
	buildHost    = ""
	buildStamp   = ""
	buildDirty   = ""
)

type VersionInformation struct {
	Short       string
	CommitHash  string
	CommitStamp string
	BuildUser   string
	BuildHost   string
	BuildStamp  string
	BuildDirty  string
	Long        string
}

var VersionInfo = VersionInformation{
	Short:       Unversioned,
	CommitHash:  "",
	CommitStamp: "",
	BuildUser:   "",
	BuildHost:   "",
	BuildStamp:  "",
	BuildDirty:  "",
	Long:        "",
}

func init() {
	stamp, _ := strconv.Atoi(commitStamp)
	commitDate := time.Unix(int64(stamp), 0).UTC().Format("2006/01/02-15:04:05")

	stamp, _ = strconv.Atoi(buildStamp)
	buildDate := time.Unix(int64(stamp), 0).UTC().Format("Mon Jan 02 15:04:05 MST 2006")

	VersionInfo = VersionInformation{
		Short:       fmt.Sprintf("%s%s", shortVersion, buildDirty),
		CommitHash:  commitHash,
		CommitStamp: commitStamp,
		BuildUser:   buildUser,
		BuildHost:   buildHost,
		BuildStamp:  buildStamp,
		BuildDirty:  buildDirty,
		Long: fmt.Sprintf(`%s-%s%s (%s) built on %s@%s (%s-%s-%s) at %s`,
			shortVersion, commitHash, buildDirty, commitDate,
			buildUser, buildHost,
			runtime.Version(), runtime.GOOS, runtime.GOARCH,
			buildDate,
		),
	}
}

func GetVersion() string {
	if len(strings.TrimSpace(VersionInfo.Short)) == 0 {
		return Unversioned
	}
	return VersionInfo.Short
}

func GetVersionLong() string {
	if GetVersion() == Unversioned {
		return Unversioned
	}
	return VersionInfo.Long
}
