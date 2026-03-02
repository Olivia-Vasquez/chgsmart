# Development Guide

## Prerequisites
- Go 1.26+
- Git 2.x
- macOS or Linux (Windows is supported but untested)

## Local Setup

### 1. Clone the repository

```bash
git clone https://github.com/Olivia-Vasquez/chgsmart.git
cd chgsmart
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Build

```bash 
go build ./cmd/chgsmart
```

**chgsmart** also supports `go install`:

```bash
go install ./cmd/chgsmart
```

## Running Locally

To run **chgsmart** locally after building:

```bash
./chgsmart --help
```

Or if installed via `go install`:

```bash
chgsmart --help
```

## Project Structure 

```
cmd/chgsmart/        – CLI entry point  
internal/area/       – Area classification logic  
internal/model/      – Core domain models  
internal/git/        – Git interaction layer
```

## Testing

**chgsmart** does not currently contain automated testing such as unit or integration tests. See `roadmap.md` for details on how the development team plans to include automated testing as part of our long-term plans.

## Linting

**chgsmart** does not currently have linting. See `roadmap.md` for details on how the development team plans to include linting as part of our long-term plans.

## Releasing

1. Update version
2. Update CHANGELOG
3. Create git tag
4. Push tag
5. GitHub Actions builds artifacts

See `release-process.md` for a more detailed release guide.

---

development.md | Created: Mar. 01, 2026 | Last Updated: Mar. 02, 2026