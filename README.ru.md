# go-toolkit

![Go](https://img.shields.io/badge/Go-1.25-00ADD8?logo=go&logoColor=white)
[![CI](https://github.com/ezhigval/go-toolkit/actions/workflows/ci.yml/badge.svg)](https://github.com/ezhigval/go-toolkit/actions/workflows/ci.yml)
![License](https://img.shields.io/badge/license-MIT-blue)
![Tier](https://img.shields.io/badge/tier-shared-0e8a16)

**[English](README.md)** · Русский

Общая инфраструктурная библиотека для Go backend-проектов портфолио.

## Пакеты

| Пакет | Описание |
|---|---|
| `logger` | Настройка slog: JSON/text, контекст запроса |
| `config` | Универсальная загрузка конфига из env |
| `httputil` | JSON-ответы, ошибки в стиле RFC 7807, health handler |
| `pgx` | Фабрика пула подключений PostgreSQL |
| `redis` | Фабрика клиента Redis |
| `middleware` | RequestID, Recover, AccessLog |

## Использование

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

## Локальная разработка

Подключается как Go-модуль. Пока библиотека в активной разработке, в зависимых проектах:

```go
// go.mod
replace github.com/ezhigval/go-toolkit => ../../shared/go-toolkit
```

## Лицензия

MIT
