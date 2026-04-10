package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Olivia-Vasquez/chgsmart/internal/area"
	"github.com/Olivia-Vasquez/chgsmart/internal/classify"
	"github.com/Olivia-Vasquez/chgsmart/internal/config"
	"github.com/Olivia-Vasquez/chgsmart/internal/gitlog"
	"github.com/Olivia-Vasquez/chgsmart/internal/model"
	"github.com/Olivia-Vasquez/chgsmart/internal/output"
)

func main() {
	// Define command-line flags with consistent naming and helpful descriptions
	var (
		fromRef       = flag.String("from", "", "git ref to start from (required)")
		toRef         = flag.String("to", "HEAD", "git ref to end at (default HEAD)")
		outPath       = flag.String("out", "", "output path (default stdout)")
		format        = flag.String("format", "", "output format: md|markdown|json|text (default based on output file extension or markdown)")
		configPath    = flag.String("config", "", "path to .chgsmart.yml in project root (default auto-detect)")
		groupBy       = flag.String("group-by", "", "group by: type|area (default from config or type)")
		maxCommits    = flag.Int("max-commits", 0, "max commits to read (0 = unlimited)")
		includeMerges = flag.Bool("include-merges", false, "include merge commits")
		help          = flag.Bool("help", false, "show help")
		versionFlag   = flag.Bool("version", false, "show version")
	)
	flag.Parse()

	// Create config structure with defaults, then override with config file and command-line flags
	type Config struct {
		fromRef       string
		toRef         string
		groupBy       string
		maxCommits    int
		includeMerges bool
		ignore        []string
	}

	// Get version info from git at build time using -ldflags
	var (
		version = "dev"
		commit  = "none"
		date    = "unknown"
	)

	if *versionFlag {
		fmt.Printf("chgsmart version %s (commit %s, date %s)\n", version, commit, date)
		return
	}

	// Create comprehensive help message that includes usage, options, and examples
	if *help {
		fmt.Printf("chgsmart - Generate changelog from git history\n\n")
		fmt.Printf("Usage:\n")
		fmt.Printf("  chgsmart --from <ref> [options]\n\n")
		fmt.Printf("Options:\n")
		flag.PrintDefaults()
		fmt.Printf("\nExamples:\n")
		fmt.Printf("  # Generate changelog from last tag to HEAD\n")
		fmt.Printf("  chgsmart --from $(git describe --tags --abrev=0)\n\n")
		fmt.Printf("  # Generate changelog for a specific range and output to file\n")
		fmt.Printf("  chgsmart --from v1.0.0 --to v1.1.0 --out CHANGELOG.md\n\n")
		return
	}

	// Load config from file
	cfg, err := config.Load(*configPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error loading config:", err)
		os.Exit(1)
	}

	// Override config with command-line flags if set, with validation for required fields
	// TODO: Add --format and --outPath to config and override logic here as well
	if *fromRef != "" {
		cfg.FromRef = *fromRef
	}
	if cfg.FromRef == "" {
		fmt.Fprintln(os.Stderr, "error: --from is required (or from_ref in config)")
		flag.Usage()
		os.Exit(2)
	}
	if *toRef != "" {
		cfg.ToRef = *toRef
	}
	if *groupBy != "" {
		cfg.GroupBy = *groupBy
	}
	if cfg.GroupBy == "" {
		cfg.GroupBy = "type"
	}
	if *maxCommits != 0 {
		cfg.MaxCommits = *maxCommits
	}
	if *includeMerges {
		cfg.IncludeMerges = *includeMerges
	}

	// Read git commits, applying filters based on config/flags
	commits, err := gitlog.ReadCommits(cfg.FromRef, cfg.ToRef, cfg.IncludeMerges, cfg.MaxCommits)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading git history:", err)
		os.Exit(1)
	}

	ignore := config.MustCompileRegexList(cfg.Ignore)
	areas := area.MustCompileAreaMatchers(cfg.Areas)

	// Classify and filter commits, preparing them for rendering
	kept := make([]model.Item, 0, len(commits))
	for _, c := range commits {
		if config.MatchesAny(ignore, c.Subject) {
			continue
		}

		ct := classify.TypeOf(c.Subject)
		breaking := classify.IsBreaking(c.Subject, c.Body)
		ar := area.MatchArea(areas, c.Subject, c.Body)

		kept = append(kept, model.Item{
			Hash:     c.Hash,
			Subject:  c.Subject,
			Type:     string(ct),
			Area:     ar,
			Breaking: breaking,
		})
	}

	// Push kept to output for rendering and file writing
	opt := model.Options{
		Title:   fmt.Sprintf("%s..%s", cfg.FromRef, cfg.ToRef),
		GroupBy: cfg.GroupBy,
		Items:   kept,
	}

	if err := output.Render(opt, *format, *outPath); err != nil {
		fmt.Fprintln(os.Stderr, "error rendering output:", err)
		os.Exit(1)
	}

	// A few test lines for new linter workflow on git
	fmt.Printf("Generated changelog with %d items\n", len(kept))
	fmt.Printf("Config: from=%s to=%s group_by=%s max_commits=%d include_merges=%v\n",
		cfg.FromRef, cfg.ToRef, cfg.GroupBy, cfg.MaxCommits, cfg.IncludeMerges)

}
