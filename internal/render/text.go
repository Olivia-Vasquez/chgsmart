package render

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/Olivia-Vasquez/chgsmart/internal/model" // import the model package to use model.Options and model.Item
)

// TextRender renders the options as a text string.
func TextRender(opt model.Options) string {
	title := opt.Title
	if title == "" {
		title = "Unreleased"
	}

	var b bytes.Buffer
	fmt.Fprintf(&b, "%s\n\n", title)

	if len(opt.Items) == 0 {
		b.WriteString("No changes.\n")
		return b.String()
	}

	// Sort items by hash for deterministic output
	items := make([]model.Item, len(opt.Items))
	copy(items, opt.Items)
	sort.SliceStable(items, func(i, j int) bool {
		return items[i].Hash < items[j].Hash
	})

	for _, item := range items {
		line := fmt.Sprintf("- %s: %s (%s)", item.Type, item.Subject, item.Area)
		b.WriteString(line + "\n")
	}

	return b.String()
}
