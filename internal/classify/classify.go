package classify

import (
	"regexp"
	"strings"

	"chgsmart/internal/model"
)

var (
	conventional = regexp.MustCompile(`^(?i)(feat|fix|refactor|perf|docs|test|chore)(\([^)]+\))?:\s+`)
)

func TypeOf(subject string) model.CommitType {
	s := strings.TrimSpace(subject)

	// 1) Conventional commits
	if conventional.MatchString(s) {
		prefix := strings.ToLower(conventional.FindStringSubmatch(s)[1])
		switch prefix {
		case "feat":
			return model.TypeAdded
		case "fix":
			return model.TypeFixed
		case "refactor":
			return model.TypeChanged
		case "perf":
			return model.TypeChanged
		case "docs":
			return model.TypeDocs
		case "test":
			return model.TypeTests
		case "chore":
			return model.TypeChore
		}
	}

	// 2) Heuristics
	l := strings.ToLower(s)
	switch {
	case containsAny(l, "fix", "bug", "patch", "hotfix"):
		return model.TypeFixed
	case containsAny(l, "add", "implement", "support", "introduce", "create"):
		return model.TypeAdded
	case containsAny(l, "refactor", "cleanup", "restructure", "rework"):
		return model.TypeChanged
	case containsAny(l, "perf", "optimiz", "speed", "faster"):
		return model.TypeChanged
	case containsAny(l, "doc", "readme", "changelog"):
		return model.TypeDocs
	case containsAny(l, "test", "spec", "snapshot"):
		return model.TypeTests
	default:
		return model.TypeChanged
	}
}

func IsBreaking(subject, body string) bool {
	s := strings.ToLower(subject)
	b := strings.ToLower(body)

	if strings.Contains(s, "breaking") ||
		strings.Contains(s, "break ") ||
		strings.Contains(s, "remove ") ||
		strings.Contains(s, "drop support") ||
		strings.Contains(s, "incompatible") {
		return true
	}
	if strings.Contains(b, "breaking change") {
		return true
	}
	// conventional commits "feat!: ..." style
	if strings.Contains(subject, "!:") {
		return true
	}
	return false
}

func containsAny(s string, subs ...string) bool {
	for _, sub := range subs {
		if strings.Contains(s, sub) {
			return true
		}
	}
	return false
}