# Architecture

# Data Flow

1. **CLI Invocation**: User runs `chgsmart` with flags (e.g., `--from`, `--to`, `--out`).
2. **Git Log Retrieval**: The tool fetches commit history between specified refs using `git log`.
3. **Parsing**: Commits are parsed into structured objects.
4. **Classification**: Each commit is classified by type (Added, Fixed, etc.) and optionally by area.
5. **Formatting**: Commits are grouped and formatted into Markdown or other output formats.
6. **Output**: Changelog is written to file or stdout.

## Module/Package Map

- `cmd/chgsmart/main.go`: CLI entrypoint, argument parsing, orchestration.
- `internal/gitlog`: Fetches and parses git history.
- `internal/classify`: Classifies commit types using heuristics and conventional commit rules.
- `internal/area`: Matches commit areas (if configured).
- `internal/config`: Loads configuration and regexes.
- `internal/render`: Formats changelog output (Markdown, etc).
- `internal/semver`: Suggests version bumps and parses versions.
- `internal/model`: Shared data structures.

## Extensibility Points

- **Categories**: Add new types in `internal/classify` and `internal/model`.
- **Areas**: Extend area matchers in `internal/area`.
- **Output Formats**: Add new formatters in `internal/render`.
- **Config**: Expand YAML config options in `internal/config`.

## ASCII Diagram

```
[CLI] --flags--> [main.go]
   |
   v
[gitlog] --commits--> [classify] --types--> [render] --output--> [file/stdout]
   |                        |
   |                        v
   |                  [area] (optional)
   |
   v
[config] --regexes--> [main.go]
```
architecture.md | Created: Feb. 25, 2026 | Last Updated: Mar. 02, 2026