package render

import (
	"sort"
	"strings"
)

type Item struct {
	Hash     string
	Subject  string
	Type     string // "Added", "Fixed", ...
	Area     string
	Breaking bool
}

type Options struct {
	Title   string // "Unreleased"
	GroupBy string // "type" or "area"
	Items   []Item
}

func Markdown(opt Options) string {
	var b strings.Builder
	title := opt.Title
	if title == "" {
		title = "Unreleased"
	}

	b.WriteString("## " + title + "\n\n")

	if len(opt.Items) == 0 {
		b.WriteString("_No changes._\n")
		return b.String()
	}

	// Sort items by hash for deterministic output
	items := make([]Item, len(opt.Items))
	copy(items, opt.Items)
	sort.SliceStable(items, func(i, j int) bool {
		return items[i].Hash < items[j].Hash
	})

	if opt.GroupBy == "area" {
		byArea := groupByArea(items)
		areas := sortedKeys(byArea)

		for _, area := range areas {
			header := area
			if header == "" {
				header = "other"
			}
			b.WriteString("### " + titleCase(header) + "\n")
			byType := groupByType(byArea[area])
			writeTypes(&b, byType)
			b.WriteString("\n")
		}
		return b.String()
	}

	byType := groupByType(items)
	writeTypes(&b, byType)
	return b.String()
}

func writeTypes(b *strings.Builder, byType map[string][]Item) {
	order := []string{"Added", "Fixed", "Changed", "Docs", "Tests", "Chore"}
	for _, t := range order {
		items := byType[t]
		if len(items) == 0 {
			continue
		}
		b.WriteString("### " + sectionName(t) + "\n")
		for _, it := range items {
			line := "- "
			if it.Area != "" {
				line += titleCase(it.Area) + ": "
			}
			line += cleanSubject(it.Subject)
			if it.Breaking {
				line += " **(BREAKING)**"
			}
			b.WriteString(line + "\n")
		}
		b.WriteString("\n")
	}

	// any unexpected types (just in case)
	extraTypes := make([]string, 0)
	for t := range byType {
		found := false
		for _, k := range order {
			if k == t {
				found = true
				break
			}
		}
		if !found {
			extraTypes = append(extraTypes, t)
		}
	}
	sort.Strings(extraTypes)
	for _, t := range extraTypes {
		items := byType[t]
		if len(items) == 0 {
			continue
		}
		b.WriteString("### " + t + "\n")
		for _, it := range items {
			b.WriteString("- " + cleanSubject(it.Subject) + "\n")
		}
		b.WriteString("\n")
	}
}

func groupByType(items []Item) map[string][]Item {
	m := map[string][]Item{}
	for _, it := range items {
		m[it.Type] = append(m[it.Type], it)
	}
	return m
}

func groupByArea(items []Item) map[string][]Item {
	m := map[string][]Item{}
	for _, it := range items {
		m[it.Area] = append(m[it.Area], it)
	}
	return m
}

func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func sectionName(t string) string {
	switch t {
	case "Added":
		return "Added"
	case "Fixed":
		return "Fixed"
	case "Changed":
		return "Changed"
	case "Docs":
		return "Documentation"
	case "Tests":
		return "Tests"
	case "Chore":
		return "Chore"
	default:
		return t
	}
}

func cleanSubject(s string) string {
	return strings.TrimSpace(strings.ReplaceAll(s, "\n", " "))
}

func titleCase(s string) string {
	if s == "" {
		return s
	}
	parts := strings.FieldsFunc(s, func(r rune) bool { return r == '-' || r == '_' || r == ' ' })
	for i := range parts {
		parts[i] = strings.ToUpper(parts[i][:1]) + strings.ToLower(parts[i][1:])
	}
	return strings.Join(parts, " ")
}
