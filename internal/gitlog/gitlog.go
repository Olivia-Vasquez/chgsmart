package gitlog

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/Olivia-Vasquez/chgsmart/internal/model"
)

func ReadCommits(fromRef, toRef string, includeMerges bool, maxCommits int) ([]model.Commit, error) {
	if strings.TrimSpace(fromRef) == "" {
		return nil, fmt.Errorf("--from is required")
	}
	if strings.TrimSpace(toRef) == "" {
		toRef = "HEAD"
	}

	rangeArg := fmt.Sprintf("%s..%s", fromRef, toRef)

	args := []string{"log", "--pretty=format:%H%x1f%s%x1f%b%x1e"}
	if !includeMerges {
		args = append(args, "--no-merges")
	}
	if maxCommits > 0 {
		args = append(args, fmt.Sprintf("-%d", maxCommits))
	}
	args = append(args, rangeArg)

	cmd := exec.Command("git", args...)
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("git log failed: %w", err)
	}

	raw := string(out)
	records := strings.Split(raw, "\x1e") // record separator
	commits := make([]model.Commit, 0, len(records))

	for _, rec := range records {
		rec = strings.TrimSpace(rec)
		if rec == "" {
			continue
		}
		fields := strings.Split(rec, "\x1f")
		if len(fields) < 2 {
			continue
		}
		hash := strings.TrimSpace(fields[0])
		subject := strings.TrimSpace(fields[1])
		body := ""
		if len(fields) >= 3 {
			body = strings.TrimSpace(fields[2])
		}
		commits = append(commits, model.Commit{
			Hash:    hash,
			Subject: subject,
			Body:    body,
		})
	}

	// Normalize bodies (git can embed newlines weirdly)
	for i := range commits {
		commits[i].Body = strings.TrimSpace(bytes.NewBufferString(commits[i].Body).String())
	}

	return commits, nil
}