# Design Notes

Why this exists: I got tired of copying the same 200 lines into every new service.

## Package boundaries

- **logger** — no external deps beyond slog
- **httputil** — HTTP only, no DB imports
- **pgx/redis** — thin wrappers, apps own queries

## Non-goals (for now)

- Full framework (no DI container)
- Kafka client — tracked in issue #2

## Consumers

Used by url-shortener, weather-proxy, file-pipeline, and upcoming middle-tier services.
