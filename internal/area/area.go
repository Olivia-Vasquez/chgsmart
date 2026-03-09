package area

import (
	"fmt"
	"regexp"
	"sort"
)

type Matcher struct {
	Area  string
	Regex *regexp.Regexp
}

// TODO: this panics on invalid regexes, which is not ideal. We should probably return an error instead
func MustCompileAreaMatchers(areas map[string][]string) []Matcher {
	matchers := make([]Matcher, 0)

	// stable ordering (nice for deterministic results)
	keys := make([]string, 0, len(areas))
	for k := range areas {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, area := range keys {
		for _, pat := range areas[area] {
			r, err := regexp.Compile(pat)
			if err != nil {

				// Throw error instead of panicking, since this is a library function and we want to allow the caller to handle errors gracefully
				text := "error compiling area regex for area '%s': %v"
				panic(fmt.Errorf(text, area, err))
			}
			matchers = append(matchers, Matcher{Area: area, Regex: r})
		}
	}
	return matchers
}

// TODO: this needs to be more efficient if there are many matchers, but for typical use cases this should be fine
func MatchArea(matchers []Matcher, subject, body string) string {
	for _, m := range matchers {
		if m.Regex.MatchString(subject) || m.Regex.MatchString(body) {
			return m.Area
		}
	}
	return ""
}