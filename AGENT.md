# AGENTS.md

This file helps AI coding agents understand the apotheke codebase.

## Overview

Apotheke is a command bookmark tool written in Go. It allows users to save frequently used commands with short names and execute them with fuzzy matching.

## Project Structure

```
apotheke/
├── cmd/root.go               # CLI entry point, cobra commands
├── internal/
│   ├── config/config.go      # XDG config paths
│   ├── db/store.go           # SQLite CRUD operations
│   ├── executor/runner.go    # Command execution with safety
│   ├── model/command.go      # Command struct and helpers
│   └── resolver/matcher.go   # Fuzzy matching with frecency
├── pkg/shell/init.go         # Shell init scripts (bash/zsh/fish)
├── docs/
│   ├── commands.md           # Command reference
│   ├── concepts.md           # Matching, safety, data
│   ├── examples.md           # Real-world examples
│   ├── tutorial.md           # User tutorial
│   └── testing.md            # Test documentation
├── go.mod
├── Makefile
└── README.md
```

## Conventions

- No comments in source files
- Tests use table-driven testing
- Test files have `_test.go` suffix
- SQLite for storage at `~/.local/share/apotheke/apotheke.db`
- XDG Base Directory spec for config paths
- License: AGPL-3.0

## Key Components

### Store (internal/db/store.go)

SQLite database operations. Methods: `Add`, `Remove`, `Get`, `List`, `Search`, `Update`, `IncrementUsage`, `GetAll`.

### Resolver (internal/resolver/matcher.go)

Fuzzy matching with priority: exact > prefix > fuzzy. Uses frecency scoring (frequency + recency).

### Executor (internal/executor/runner.go)

Runs commands with dry-run and confirmation support. Dangerous commands (tag "danger" or --confirm flag) require user confirmation.

### Shell (pkg/shell/init.go)

Generates shell init scripts that create the `a` function alias.

## Commands

- `apotheke add <name> <command>` - add bookmark
- `apotheke rm <name>` - remove bookmark
- `apotheke list` - list all (supports fuzzy search with `-q`)
- `apotheke show <name>` - show details
- `apotheke <query> [args...]` - execute matched command
- `apotheke init <shell>` - print shell init script

## Building

```bash
go build -o apotheke ./cmd
go test ./...
```

## Dependencies

- github.com/spf13/cobra - CLI framework
- github.com/mattn/go-sqlite3 - SQLite driver
- github.com/sahilm/fuzzy - Fuzzy matching
- github.com/fatih/color - Colored output

## Common Tasks for AI Agents

### Adding a new CLI flag

1. Add flag variable in `cmd/root.go`
2. Register in `init()` function
3. Use in run function
4. Update README.md

### Adding a new store method

1. Add method to `internal/db/store.go`
2. Add test to `internal/db/store_test.go`
3. Use in CLI

### Adding shell support

1. Add function in `pkg/shell/init.go`
2. Add case in `Init()` switch
3. Add test in `pkg/shell/init_test.go`

## Pitfalls to Avoid

- Do NOT add comments to source files
- Do NOT use `store.Search()` for fuzzy matching - use `resolver.Resolve()` instead
- Do NOT auto-execute when multiple commands match - show picker
- Do NOT skip confirmation for commands with `danger` tag or `--confirm` flag
- Arguments after query in `a <query> [args...]` are APPENDED to the saved command

## Testing

```bash
# Run all tests
go test ./...

# Run specific package
go test ./internal/db/...

# Verbose
go test ./... -v
```

## Documentation

- [Commands](docs/commands.md) - Command reference
- [Concepts](docs/concepts.md) - Matching, safety, data
- [Examples](docs/examples.md) - Real-world examples
- [Tutorial](docs/tutorial.md) - User tutorial
- [Testing](docs/testing.md) - Test documentation
