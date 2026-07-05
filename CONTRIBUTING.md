# Contributing

Personal library — external PRs welcome for bug fixes.

## Development

```bash
go test ./...
golangci-lint run ./...
```

## Versioning

Semver tags. Breaking API changes bump minor until v1.0.0 stabilizes.

## Commit style

Plain language preferred. Examples: `add redis ping timeout`, `fix nil logger in middleware`.
