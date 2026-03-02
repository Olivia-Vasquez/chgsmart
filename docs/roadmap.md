# Roadmap

This document contains a breakdown of milestones and tasks currently planned for the first release (v1.0) of chgsmart.

v1.0 has five major milestones, outlined below.

## Milestone 1
### Repository Skeleton
**Goal Completion: March 11, 2026**

Establish a clean, professional repository structure that clearly communicates project scope, development workflow, and contribution expectations.

* Define CLI scope and non-goals in documentation
* Standardize repository structure (`cmd/`, `internal/`, `docs/`, etc.)
* Add `CONTRIBUTING.md`
* Add Code of Conduct
* Add issue templates (bug report, feature request)
* Add pull request template with review checklist
* Create `docs/development.md` with local setup and workflow instructions
* Ensure README clearly links to development documentation

---

## Milestone 2
### CLI Polish and UX Stability
**Goal Completion: March 25, 2026**

Refine the CLI surface area and ensure predictable, production-ready behavior.

* Finalize command surface and flag conventions
* Standardize `--help` output formatting
* Implement consistent exit codes and structured error handling
* Improve user-facing error messages
* Ensure deterministic output ordering
* Add optional configuration support (e.g., `.chgsmart.yaml` or environment variables)
* Add optional output formats (e.g., `text`, `json`)
* Address identified UX and edge-case issues

---

## Milestone 3
### Quality, Testing, and Continuous Integration
**Goal Completion: April 8, 2026**

Introduce automated quality gates and establish a reliable CI pipeline.

* Add unit tests for core logic
* Add golden tests for CLI output validation
* Ensure test coverage for critical functionality
* Add linting and formatting enforcement
* Configure GitHub Actions to:
  * Run tests on pull requests
  * Enforce linting
  * Validate build success across supported platforms
* Document testing strategy in `docs/development.md`

---

## Milestone 4
### Release Engineering and Versioning
**Goal Completion: April 22, 2026**

Establish a formal release process aligned with semantic versioning.

* Adopt and document semantic versioning strategy
* Define and document changelog conventions
* Create `CHANGELOG.md`
* Implement automated GitHub release workflow triggered by tags
* Configure cross-platform binary builds
* Attach compiled artifacts to GitHub releases
* Add release verification checklist
* Create `docs/release-process.md` documenting:
  * Version bump procedure
  * Tag creation
  * Changelog updates
  * Artifact verification

---

## Milestone 5
### Distribution and Installation Experience
**Goal Completion: May 6, 2026**

Ensure chgsmart can be easily installed and adopted by new users.

* Provide clear installation instructions (macOS, Linux, Windows)
* Add checksum generation and verification instructions
* Evaluate Homebrew distribution strategy
* (If appropriate) Create and publish a Homebrew tap
* Validate installation flow from a clean environment
* Perform final v1.0 release readiness review


---

## Timeline Overview

| Milestone | Focus Area | Target Completion |
|-----------|------------|------------------|
| 1 | Repository Skeleton | March 11, 2026 |
| 2 | CLI Polish | March 25, 2026 |
| 3 | Quality & CI | April 8, 2026 |
| 4 | Release Engineering | April 22, 2026 |
| 5 | Distribution | May 6, 2026 |

---

roadmap.md | Created: Feb. 25, 2026 | Last Updated: Mar. 02, 2026