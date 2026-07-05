# Changelog

Format based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/).

## [Unreleased]

### Added
- Kafka helper package (planned)

## [0.1.0] - 2026-07-06

### Added
- `logger` — slog JSON/text handlers, request context
- `config` — generic env loader
- `httputil` — JSON responses, problem details, health handler
- `pgx` — pool factory with ping
- `redis` — client factory
- `middleware` — RequestID, Recover, AccessLog

[Unreleased]: https://github.com/ezhigval/go-toolkit/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/ezhigval/go-toolkit/releases/tag/v0.1.0
