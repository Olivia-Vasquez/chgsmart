# Changelog

## [Unreleased]

- Initial implementation of core CLI
- Git history parsing and classification
- Markdown changelog output

### Added
- Add pull request template to repository
- Enhance documentation across multiple files: 
    - Update README.md for improved clarity and emphasis on tool functionality. 
    - Refine architecture.md by removing redundant descriptions. 
    - Clarify classification strategy in classification-strategy.md. 
    - Add code-of-conduct.md and contributing.md with appreciation notes. 
    - Expand roadmap.md with detailed milestones and tasks for v1.0. 
    - Introduce development.md for local setup and project structure guidance. 
    - Update scope.md to define goals and boundaries for v1.0. 
    - Revise release-process.md to include detailed steps for releases.
- Add Code of Conduct, Contributing guidelines, and Scope document.
- Add ldflags configuration for versioning in goreleaser
- Add initial configuration and documentation files for changelog generation tool
- Update README.md for improved clarity and add requirements section; add .DS_Store to cmd directory
- Update installation instructions and add help flag; refactor changelog generation options
- Add core functionality for changelog generation and configuration

### Changed
- Enhance error handling in regex compilation and improve command-line flag descriptions.
- Update issue templates
- Remove pyproject.toml configuration file **(BREAKING)**
- Refactor code structure for improved readability and maintainability

### Documentation
- Update creation date in README.md
- Update repository URL in README
---

> Releases are not yet cut. Version format will follow `vX.Y.Z` once release process is implemented.
