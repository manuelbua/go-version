package version

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

const Unversioned = "(unversioned)"

// populated via -ldflags
var (
	branch         = ""
	tag            = ""
	countTagToHead = ""
	commitHash     = ""
	commitStamp    = ""
	buildUser      = ""
	buildHost      = ""
	buildStamp     = ""
	buildDirty     = ""
)

type VersionInformation struct {
	Branch         string
	Tag            string
	CountTagToHead int
	CommitHash     string
	CommitStamp    string
	BuildUser      string
	BuildHost      string
	BuildStamp     string
	BuildDirty     bool
	Short          string
	Long           string
}

var VersionInfo = VersionInformation {
	Branch:         "",
	Tag:            "",
	CountTagToHead: 0,
	CommitHash:     "",
	CommitStamp:    "",
	BuildUser:      "",
	BuildHost:      "",
	BuildStamp:     "",
	BuildDirty:     false,
	Short:          Unversioned,
	Long:           "",
}

func init() {
	stamp, _ := strconv.Atoi(commitStamp)
	commitDate := time.Unix(int64(stamp), 0).UTC().Format("2006/01/02 15:04:05 MST 2006")

	stamp, _ = strconv.Atoi(buildStamp)
	buildDate := time.Unix(int64(stamp), 0).UTC().Format("Mon Jan 02 15:04:05 MST 2006")

    countTag, _ := strconv.Atoi(countTagToHead)
    isDirty, _ := strconv.ParseBool(buildDirty)

	VersionInfo = VersionInformation{
        Branch:      branch,
        Tag:         tag,
        CountTagToHead: countTag,
		CommitHash:  commitHash,
		CommitStamp: commitStamp,
		BuildUser:   buildUser,
		BuildHost:   buildHost,
		BuildStamp:  buildStamp,
		BuildDirty:  isDirty,
        Short:       "",
		Long:        "",
	}

    baseVersion := tag
    if VersionInfo.CountTagToHead > 0 {
        baseVersion += fmt.Sprintf("+%d", countTag)
    }

    if VersionInfo.BuildDirty {
        baseVersion += "(dirty)"
    }

    VersionInfo.Short = baseVersion
    VersionInfo.Long = fmt.Sprintf("%s-%s (%s)\nBuilt by %s@%s (%s-%s-%s) on %s",
        baseVersion, commitHash, commitDate,
        buildUser, buildHost,
        runtime.Version(), runtime.GOOS, runtime.GOARCH,
        buildDate,
    )
}

func GetVersion() string {
	return VersionInfo.Short
}

func GetVersionLong() string {
	if GetVersion() == Unversioned {
		return Unversioned
	}
	return VersionInfo.Long
}
