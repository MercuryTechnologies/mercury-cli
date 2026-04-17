package updatecheck

import (
	"strconv"
	"strings"
)

type semver struct {
	major, minor, patch int
}

// parseSemver accepts strings of the exact form "X.Y.Z" with non-negative integers.
// It rejects pre-release tags, build metadata, and a leading "v" so that we never
// nag users about unreleased builds.
func parseSemver(s string) (semver, bool) {
	s = strings.TrimSpace(s)
	if s == "" {
		return semver{}, false
	}
	parts := strings.Split(s, ".")
	if len(parts) != 3 {
		return semver{}, false
	}
	var out semver
	for i, p := range parts {
		if p == "" {
			return semver{}, false
		}
		n, err := strconv.Atoi(p)
		if err != nil || n < 0 {
			return semver{}, false
		}
		switch i {
		case 0:
			out.major = n
		case 1:
			out.minor = n
		case 2:
			out.patch = n
		}
	}
	return out, true
}

// less reports whether a is strictly older than b.
func (a semver) less(b semver) bool {
	if a.major != b.major {
		return a.major < b.major
	}
	if a.minor != b.minor {
		return a.minor < b.minor
	}
	return a.patch < b.patch
}
