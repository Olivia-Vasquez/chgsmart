package output

import (
	"fmt"
	"os"
	"strings"

	"github.com/Olivia-Vasquez/chgsmart/internal/model" // import the model package to use model.Options and model.Item
	"github.com/Olivia-Vasquez/chgsmart/internal/render"
	"github.com/Olivia-Vasquez/chgsmart/internal/semver" // import the semver package to use semver.SuggestBump
)

// Render renders the options to specified output path. The file extension is determined by format parameter or inferred from the output path. If no output path is provided, it prints to stdout.
func Render(opt model.Options, format string, outPath string) error {

	// Normalize "md" alias to "markdown"
	if format == "md" {
		format = "markdown"
	}

	// Get file extension and determine format if not explicitly set
	if format == "" {
		format = "markdown" // default format
		if len(outPath) > 0 {
			switch {
			case hasExtension(outPath, ".json"):
				format = "json"
			case hasExtension(outPath, ".txt", ".text"):
				format = "text"
			case hasExtension(outPath, ".md", ".markdown"):
				format = "markdown"
			}
		}
	}

	var out string
	switch format {
	case "json":
		if len(outPath) > 0 && !hasExtension(outPath, ".json") {
			outPath = outPath + ".json"
		}
		bump := semver.SuggestBump(opt.Items)
		out = render.JSONRender(opt, string(bump))
	case "text":
		if len(outPath) > 0 && !hasExtension(outPath, ".txt", ".text") {
			outPath = outPath + ".txt"
		}
		bump := semver.SuggestBump(opt.Items)
		out = render.TextRender(opt)
		out += fmt.Sprintf("\nSuggested version bump: %s\n", bump)
	case "markdown":
		if len(outPath) > 0 && !hasExtension(outPath, ".md", ".markdown") {
			outPath = outPath + ".md"
		}
		bump := semver.SuggestBump(opt.Items)
		out = render.MarkdownRender(opt)
		out += fmt.Sprintf("\nSuggested version bump: %s\n", bump)
	default:
		return fmt.Errorf("unsupported output format: %s", format)
	}
	// Save final output to file
	if len(outPath) > 0 {
		if err := os.WriteFile(outPath, []byte(out), 0644); err != nil {
			return fmt.Errorf("failed to write output to file: %w", err)
		}
		fmt.Printf("Output written to %s\n", outPath)
	} else {
		fmt.Println(out)
	}

	return nil
}

func hasExtension(path string, exts ...string) bool {
	lowerPath := strings.ToLower(path)
	for _, ext := range exts {
		if strings.HasSuffix(lowerPath, strings.ToLower(ext)) {
			return true
		}
	}
	return false
}
