package main

import (
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
	var (
		fromRef       = flag.String("from", "", "git ref to start from (required)")
		toRef         = flag.String("to", "HEAD", "git ref to end at (default HEAD)")
		outPath       = flag.String("out", "", "output path (default stdout)")
		configPath    = flag.String("config", "", "path to .chgsmart.yml (default auto-detect)")
		groupBy       = flag.String("group-by", "", "group by: type|area (default from config or type)")
		maxCommits    = flag.Int("max-commits", 0, "max commits to read (0 = unlimited)")
		includeMerges = flag.Bool("include-merges", false, "include merge commits")
	)
	flag.Parse()

	if *fromRef == "" {
		fmt.Fprintln(os.Stderr, "error: --from is required")
		flag.Usage()
		os.Exit(2)
	}

	cfg, err := config.Load(*configPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error loading config:", err)
		os.Exit(1)
	}
	if *groupBy != "" {
		cfg.GroupBy = *groupBy
	}
	if cfg.GroupBy == "" {
		cfg.GroupBy = "type"
	}

	commits, err := gitlog.ReadCommits(*fromRef, *toRef, *includeMerges, *maxCommits)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading git history:", err)
		os.Exit(1)
	}

	ignore := config.MustCompileRegexList(cfg.Ignore)
	areas := area.MustCompileAreaMatchers(cfg.Areas)

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
	if v, ok := semver.ParseVersionFromRef(*fromRef); ok {
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