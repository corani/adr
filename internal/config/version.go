package config

import (
	"runtime/debug"
)

type Version struct {
	Version string
	Commit  string
	BuiltAt string
}

func GetVersion() Version {
	version := Version{
		Version: "dev",
		Commit:  "none",
		BuiltAt: "unknown",
	}

	if info, ok := debug.ReadBuildInfo(); ok && info != nil {
		if info.Main.Version != "" && info.Main.Version != "(devel)" {
			version.Version = info.Main.Version
		}

		for _, s := range info.Settings {
			switch s.Key {
			case "vcs.revision":
				version.Commit = s.Value
			case "vcs.time":
				version.BuiltAt = s.Value
			}
		}
	}

	return version
}
