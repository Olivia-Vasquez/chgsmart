# Roadmap

## Vision
chgsmart aims to make changelog generation effortless and reliable for any project, even with inconsistent commit messages. The goal is to provide clear, categorized release notes with minimal manual effort.

---

### Now
- **Stabilize core CLI**
  - Ensure reliable git parsing, classification, and output.
  - Rationale: Foundation for all features; must be robust.

### Next
- **Configurable classification rules**
  - Allow users to define custom categories and heuristics.
  - Rationale: Adapt to diverse workflows.
- **Additional output formats**
  - Support plain text, JSON, and custom templates.
  - Rationale: Integrate with more tools and pipelines.

### Later
- **GitHub Action integration**
  - Automate changelog generation in CI/CD.
  - Rationale: Streamline release process.
- **PR-based release notes**
  - Summarize changes from pull requests.
  - Rationale: Improve changelog accuracy for teams.
- **Area detection from commit scope**
  - Infer areas automatically from commit messages.
  - Rationale: Reduce config burden.
