package version

import "runtime/debug"

func String() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		return info.Main.Version
	}
	return "failed to read version with 'debug.ReadBuildInfo'"
}
