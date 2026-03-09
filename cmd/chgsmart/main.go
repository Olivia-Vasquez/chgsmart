package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Olivia-Vasquez/chgsmart/internal/area"
	"github.com/Olivia-Vasquez/chgsmart/internal/classify"
	"github.com/Olivia-Vasquez/chgsmart/internal/config"
	"github.com/Olivia-Vasquez/chgsmart/internal/gitlog"
	"github.com/Olivia-Vasquez/chgsmart/internal/render"
	"github.com/Olivia-Vasquez/chgsmart/internal/semver"
)

func main() {
	// Define command-line flags with consistent naming and helpful descriptions
	var (
		fromRef       = flag.String("from", "", "git ref to start from (required)")
		toRef         = flag.String("to", "HEAD", "git ref to end at (default HEAD)")
		outPath       = flag.String("out", "", "output path (default stdout)")
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
		fromRef 	 string
		toRef 		 string
		groupBy       string
		maxCommits    int
		includeMerges bool
		ignore		  []string
	}

	func parseConfig(jsonData []byte) Config{
		var cfg Config
		if err := json.Unmarshal(jsonData, &cfg); err != nil {
			fmt.Fprintln(os.Stderr, "error parsing config:", err)
			os.Exit(1)
		} else
		return cfg
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

	       cfg, err := config.Load(*configPath)
	       if err != nil {
		       fmt.Fprintln(os.Stderr, "error loading config:", err)
		       os.Exit(1)
	       }

	       // CLI flags override config values if set
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
	kept := make([]render.Item, 0, len(commits))
	for _, c := range commits {
		if config.MatchesAny(ignore, c.Subject) {
			continue
		}

		ct := classify.TypeOf(c.Subject)
		breaking := classify.IsBreaking(c.Subject, c.Body)
		ar := area.MatchArea(areas, c.Subject, c.Body)

		kept = append(kept, render.Item{
			Hash:     c.Hash,
			Subject:  c.Subject,
			Type:     string(ct),
			Area:     ar,
			Breaking: breaking,
		})
	}

	       md := render.Markdown(render.Options{
		       Title:   "Unreleased",
		       GroupBy: cfg.GroupBy,
		       Items:   kept,
	       })

	       // Semver bump suggestion (best-effort)
	       bump := semver.SuggestBump(kept)
	       var bumpLine string
	       if v, ok := semver.ParseVersionFromRef(cfg.FromRef); ok {
		       next := semver.ApplyBump(v, bump)
		       bumpLine = fmt.Sprintf("\nSuggested version bump: %s (%s -> %s)\n", bump, v.StringWithV(), next.StringWithV())
	       } else {
		       bumpLine = fmt.Sprintf("\nSuggested version bump: %s\n", bump)
	       }

	output := md + bumpLine

	if *outPath == "" {
		fmt.Print(output)
		return
	}

	if err := os.MkdirAll(filepath.Dir(*outPath), 0o755); err != nil && filepath.Dir(*outPath) != "." {
		fmt.Fprintln(os.Stderr, "error creating output dir:", err)
		os.Exit(1)
	}
	if err := os.WriteFile(*outPath, []byte(output), 0o644); err != nil {
		fmt.Fprintln(os.Stderr, "error writing output:", err)
		os.Exit(1)
	}
}
