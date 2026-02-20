package semver

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/Olivia-Vasquez/chgsmart/internal/render"
)

type Version struct {
	Major int
	Minor int
	Patch int
}

func (v Version) StringWithV() string {
	return "v" + strconv.Itoa(v.Major) + "." + strconv.Itoa(v.Minor) + "." + strconv.Itoa(v.Patch)
}

type Bump string

const (
	Major Bump = "major"
	Minor Bump = "minor"
	Patch Bump = "patch"
)

func SuggestBump(items []render.Item) Bump {
	hasAdded := false
	for _, it := range items {
		if it.Breaking {
			return Major
		}
		if it.Type == "Added" {
			hasAdded = true
		}
	}
	if hasAdded {
		return Minor
	}
	return Patch
}

func ApplyBump(v Version, bump Bump) Version {
	switch bump {
	case Major:
		return Version{Major: v.Major + 1, Minor: 0, Patch: 0}
	case Minor:
		return Version{Major: v.Major, Minor: v.Minor + 1, Patch: 0}
	default:
		return Version{Major: v.Major, Minor: v.Minor, Patch: v.Patch + 1}
	}
}

var verRe = regexp.MustCompile(`(?i)\bv?(\d+)\.(\d+)\.(\d+)\b`)

func ParseVersionFromRef(ref string) (Version, bool) {
	m := verRe.FindStringSubmatch(ref)
	if len(m) != 4 {
		return Version{}, false
	}
	maj, _ := strconv.Atoi(m[1])
	min, _ := strconv.Atoi(m[2])
	pat, _ := strconv.Atoi(m[3])
	return Version{Major: maj, Minor: min, Patch: pat}, true
}

// convenience wrapper: accept tags like "v1.2.3" or "1.2.3"
func ParseVersionFromRefMaybe(ref string) (Version, bool) {
	return ParseVersionFromRef(strings.TrimSpace(ref))
}