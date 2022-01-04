package config

type Version struct {
	Version string
	Commit  string
	BuiltAt string
	BuiltBy string
}

func SetVersion(v, c, d, b string) {
	version = Version{
		Version: v,
		Commit:  c,
		BuiltAt: d,
		BuiltBy: b,
	}
}

func GetVersion() Version {
	return version
}

//nolint:gochecknoglobals
var version Version
