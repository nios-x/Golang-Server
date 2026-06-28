# articles-go

A minimal Go web application for managing articles. This repository contains a small HTTP server, configuration, and SQLite storage for development.

## Prerequisites

- Go 1.18 or newer

## Quick Start

Run the application directly from the module root:

```bash
go run ./articles-app
```

Build a binary:

```bash
go build -o bin/articles ./articles-app
# On Windows: .\bin\articles.exe
```

## Configuration

The app reads configuration from config/local.yaml. Edit that file to adjust database paths and settings for local development.

## Project Structure

- `articles-app/` — application entrypoint (`main.go`)
- `internal/config/` — configuration loader and types
- `internal/http/handlers/` — HTTP handlers (users, articles, etc.)
- `internal/storage/` — storage interfaces
- `internal/storage/sqlite/` — SQLite implementation
- `internal/types/` — shared types
- `util/response/` — helper for HTTP responses

## Notes

- Designed for local development; production deployment may require configuration changes (database, logging, secrets).

## License

MIT
