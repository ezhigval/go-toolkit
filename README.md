# go-toolkit

![Go](https://img.shields.io/badge/Go-1.25-00ADD8?logo=go&logoColor=white)
[![CI](https://github.com/ezhigval/go-toolkit/actions/workflows/ci.yml/badge.svg)](https://github.com/ezhigval/go-toolkit/actions/workflows/ci.yml)
![License](https://img.shields.io/badge/license-MIT-blue)
![Tier](https://img.shields.io/badge/tier-shared-0e8a16)

Shared infrastructure library for Go backend portfolio projects.

## Packages

| Package | Description |
|---|---|
| `logger` | slog setup with JSON/text output, request context |
| `config` | Generic env config loader |
| `httputil` | JSON responses, RFC 7807-style errors, health handler |
| `pgx` | PostgreSQL connection pool factory |
| `redis` | Redis client factory |
| `middleware` | RequestID, Recover, AccessLog |

## Usage

```go
import (
    "github.com/ezhigval/go-toolkit/config"
    "github.com/ezhigval/go-toolkit/logger"
    tkpgx "github.com/ezhigval/go-toolkit/pgx"
)

type AppConfig struct {
    Port string `env:"PORT" envDefault:"8080"`
    DatabaseURL string `env:"DATABASE_URL,required"`
}

func main() {
    cfg := config.MustLoad[AppConfig]()
    log := logger.New(logger.Config{Level: "info", Format: "json"})

    pool, err := tkpgx.NewPool(context.Background(), tkpgx.Config{
        URL: cfg.DatabaseURL,
    })
    if err != nil {
        log.Error("db connection failed", "error", err)
        os.Exit(1)
    }
    defer pool.Close()
}
```

## Local development

Used as a Go module. During active development, consuming projects use:

```go
// go.mod
replace github.com/ezhigval/go-toolkit => ../../shared/go-toolkit
```

## License

MIT
