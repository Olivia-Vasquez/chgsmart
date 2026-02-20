package area

import (
	"regexp"
	"sort"
)

type Matcher struct {
	Area  string
	Regex *regexp.Regexp
}

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
				panic(err)
			}
			matchers = append(matchers, Matcher{Area: area, Regex: r})
		}
	}
	return matchers
}

func MatchArea(matchers []Matcher, subject, body string) string {
	for _, m := range matchers {
		if m.Regex.MatchString(subject) || m.Regex.MatchString(body) {
			return m.Area
		}
	}
	return ""
}