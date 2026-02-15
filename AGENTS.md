# AGENTS.md

## Project

Go library for parsing user agent strings. Zero dependencies.

- Module: `github.com/infobits-io/useragent`
- Go version: 1.26
- Main branch: `master`

## Commands

```bash
make lint     # Run golangci-lint
make test     # Run tests with -v -race
make build    # Build the package
```

## Code Style

- Follow existing patterns in the codebase
- All exported functions and methods must have doc comments ending with a period
- Max line length is 200 characters
- Use `gofumpt` formatting conventions
- Tests must call `t.Parallel()` at both the top-level and subtest level

## Linting

Uses golangci-lint v2 with a strict linter set configured in `.golangci.yml`. Always run `make lint` before committing. The linter auto-fixes formatting issues.

## Testing

Run `make test` to execute tests with race detection. All new functionality must include tests.

## Commits

Use conventional commit style (e.g. `feat:`, `fix:`, `chore:`, `docs:`).
