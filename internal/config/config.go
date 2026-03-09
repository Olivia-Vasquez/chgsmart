package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"gopkg.in/yaml.v3"
)

type Config struct {
	GroupBy       string              `yaml:"group_by"`
	Ignore        []string            `yaml:"ignore"`
	Areas         map[string][]string `yaml:"areas"`
	FromRef       string              `yaml:"from_ref"`
	ToRef         string              `yaml:"to_ref"`
	MaxCommits    int                 `yaml:"max_commits"`
	IncludeMerges bool                `yaml:"include_merges"`
}

func Default() Config {
	return Config{
		GroupBy:       "type",
		Ignore:        []string{`^Merge `, `^WIP`},
		Areas:         map[string][]string{},
		FromRef:       "",
		ToRef:         "HEAD",
		MaxCommits:    0,
		IncludeMerges: false,
	}
}

func Load(path string) (Config, error) {
	cfg := Default()

	// Auto-detect .chgsmart.yml in project root if path not provided
	if path == "" {
		if _, err := os.Stat(".chgsmart.yml"); err == nil {
			path = ".chgsmart.yml"
		}
	}

	if path == "" {
		return cfg, nil // no config is fine
	}

	b, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return Config{}, err
	}
	if err := yaml.Unmarshal(b, &cfg); err != nil {
		return Config{}, err
	}

	if cfg.GroupBy != "" && cfg.GroupBy != "type" && cfg.GroupBy != "area" {
		return Config{}, errors.New("config group_by must be 'type' or 'area'")
	}
	if cfg.Ignore == nil {
		cfg.Ignore = []string{}
	}
	if cfg.Areas == nil {
		cfg.Areas = map[string][]string{}
	}
	return cfg, nil
}

func MustCompileRegexList(patterns []string) []*regexp.Regexp {
	rs := make([]*regexp.Regexp, 0, len(patterns))
	for _, p := range patterns {
		r, err := regexp.Compile(p)
		if err != nil {
			text := "error compiling regex pattern '%s': %v"
			panic(fmt.Errorf(text, p, err))
		}
		rs = append(rs, r)
	}
	return rs
}

func MatchesAny(rs []*regexp.Regexp, s string) bool {
	for _, r := range rs {
		if r.MatchString(s) {
			return true
		}
	}
	return false
}
