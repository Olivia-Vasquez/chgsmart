# chgsmart
Generate human-friendly changelogs from git history, even when commit messages aren’t perfect.

**chgsmart** is a Go-based CLI tool for generating categorized changelogs from git history. It works even with imperfect commit messages, using heuristics and conventional commit parsing to classify changes.

## Status
Alpha – Core functionality is present; release process and advanced features are planned.

## Features

- Parses git history to generate readable changelogs
- Handles imperfect or inconsistent commit messages
- Groups related changes for clarity
- Outputs Markdown changelogs
- Configurable via YAML (areas, ignore patterns)

## Requirements
- Go 1.20 or newer

## Installation

Install with Go:

```bash
go install github.com/Olivia-Vasquez/chgsmart/cmd/chgsmart@latest
```
Ensure your Go bin directory is in your PATH:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

## Quickstart
Run in your project directory

```bash
chgsmart --from v0.1.0 --to HEAD --out CHANGELOG.md
```


Common options:

- `--from <ref>`: Start git ref (required)
- `--to <ref>`: End git ref (default HEAD)
- `--out <filepath>`:Output path (default stdout)
- `--group-by <category>`: Group by category (`type`, `area`) 
- `--include-merges`: Include merge commit
- `--max-commits <N>`: Limit number of commits
- `--versionFlag`: Version details

Example:

```bash
chgsmart --from v0.1.0 --to HEAD --group-by area --out CHANGELOG.md
``` 

## Example Output
```
## Unreleased

### Added
- feat: add login endpoint

### Fixed
- fix: patch race condition

### Changed
- refactor: cleanup user handler
```

## Documentation

- docs/architecture.md
- docs/classification-strategy.md
- docs/code-of-conduct.md
- docs/development.md
- docs/release-process.md
- docs/roadmap.md
- docs/scope.md

## Customization

You can customize changelog templates by editing the `templates/` directory. See the [docs](docs/customization.md) for details.

## Contributing

Contributions are welcome! Please open issues or pull requests for bug fixes, features, or documentation improvements.

## License

This project is licensed under the MIT License.

## Contact

For questions or feedback, open an issue or email oliviavasquez@fakemail.com.

---

README.md | Created ... | Last Updated: Mar. 02, 2026