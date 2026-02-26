# Classification Strategy

## What is Classification?

Classification in chgsmart means assigning each commit to a category (e.g., Added, Fixed, Changed) based on its message. This helps generate clear, grouped changelogs.

## Current Approach

- **Conventional Commit Parsing**: Recognizes prefixes like `feat`, `fix`, `docs`, etc.
- **Heuristics**: Uses keywords in commit messages to infer type (e.g., "add", "bug", "refactor").
- **Fallback**: Defaults to "Changed" if no clear match.

## Categories

- Added
- Fixed
- Changed
- Docs
- Tests
- Chore

## Examples

| Commit Message                              | Classified As |
|---------------------------------------------|---------------|
| "feat: add login endpoint"                  | Added         |
| "fix: patch race condition"                 | Fixed         |
| "refactor: cleanup user handler"            | Changed       |
| "docs: update README"                       | Docs          |
| "test: add integration tests"               | Tests         |
| "chore: update dependencies"                | Chore         |
| "add support for YAML config"               | Added         |
| "bug: resolve crash on startup"             | Fixed         |
| "optimize query performance"                | Changed       |
| "readme tweaks"                             | Docs          |

## Known Limitations

- Relies on message content; unclear messages may be misclassified.
- Only basic conventional commit prefixes supported.
- Area classification requires config and is optional.

## Future Improvements (Planned)

- Custom classification rules via config
- Support for additional commit types
- Improved handling of ambiguous messages
- Area detection from commit scope
