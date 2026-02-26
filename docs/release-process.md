# Release Process

## Local Development

- Build: `go build ./cmd/chgsmart`
- Run: `go run ./cmd/chgsmart --from <ref> --to <ref>`
- Install: `go install github.com/Olivia-Vasquez/chgsmart/cmd/chgsmart@latest`

## Versioning

- Semantic versioning (major.minor.patch) is planned.
- Version bump suggestion based on commit types (see `internal/semver`).

## Cutting a Release

1. Update version (Planned)
2. Run tests (Planned)
3. Generate changelog:
   - `chgsmart --from <last-tag> --to HEAD --out CHANGELOG.md`
4. Tag git:
   - `git tag vX.Y.Z`
   - `git push --tags`
5. Build binaries (Planned)

## GitHub Actions (Planned)

- Automated changelog generation
- Release tagging
- Build and upload binaries

