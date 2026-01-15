# Architecture

## Overview

Apotheke is a command bookmark tool built with a layered architecture.

```
┌─────────────────────────────────────────────────────────┐
│                     CLI Layer                           │
│  ┌─────┐ ┌────┐ ┌──────┐ ┌──────┐ ┌────────┐ ┌──────┐  │
│  │ add │ │ rm │ │ list │ │ show │ │ <query>│ │ init │  │
│  └──┬──┘ └─┬──┘ └──┬───┘ └──┬───┘ └───┬────┘ └──┬───┘  │
└─────┼──────┼───────┼────────┼─────────┼─────────┼───────┘
      │      │       │        │         │         │
      v      v       v        v         v         v
┌─────────────────────────────────────────────────────────┐
│                    Core Layer                           │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐              │
│  │  Store   │  │ Resolver │  │ Executor │              │
│  │ (SQLite) │  │ (Fuzzy)  │  │  (Run)   │              │
│  └────┬─────┘  └────┬─────┘  └────┬─────┘              │
└───────┼─────────────┼─────────────┼─────────────────────┘
        │             │             │
        v             │             v
┌───────────────┐     │     ┌───────────────┐
│   Database    │◄────┘     │    Shell      │
│  ~/.local/    │           │ bash/zsh/fish │
│  share/       │           └───────────────┘
│  apotheke/    │
│  apotheke.db  │
└───────────────┘
```

## Components

### CLI Layer (`cmd/apotheke/main.go`)

Entry point using Cobra framework. Handles:
- Command parsing
- Flag processing
- User interaction (prompts, output)

### Store (`internal/db/store.go`)

SQLite database operations:
- `Add()` - Insert new command
- `Remove()` - Delete command
- `Get()` - Get by exact name
- `List()` - List all, optionally filter by tag
- `Search()` - Substring search
- `Update()` - Modify existing
- `IncrementUsage()` - Update frequency and last_used
- `GetAll()` - Get all for fuzzy matching

### Resolver (`internal/resolver/matcher.go`)

Fuzzy matching engine:
- Priority: Exact > Prefix > Fuzzy
- Frecency scoring: `frequency × recency_factor`
- Uses `github.com/sahilm/fuzzy` for fuzzy matching

### Executor (`internal/executor/runner.go`)

Command execution:
- Build full command with args
- Dry-run mode
- Confirmation prompts
- Working directory support
- Shell execution via `$SHELL -c`

### Config (`internal/config/config.go`)

Configuration management:
- XDG Base Directory spec
- Default: `~/.local/share/apotheke/`

### Model (`internal/model/command.go`)

Data structures:
```go
type Command struct {
    ID        int64
    Name      string     // short name
    Cmd       string     // actual command
    Cwd       *string    // working directory
    Tags      string     // comma-separated
    Confirm   bool       // require confirmation
    Frequency int        // usage count
    LastUsed  *time.Time // for ranking
    CreatedAt time.Time
}
```

### Shell (`pkg/shell/init.go`)

Shell integration scripts for bash, zsh, fish.

## Data Flow

### Adding a command

```
User: a add kdp "kubectl delete pod"
  │
  ▼
CLI parses args: name="kdp", cmd="kubectl delete pod"
  │
  ▼
Store.Add() inserts into SQLite
  │
  ▼
Success message
```

### Running a command

```
User: a kd my-pod
  │
  ▼
CLI gets query="kd", args=["my-pod"]
  │
  ▼
Store.GetAll() fetches all commands
  │
  ▼
Resolver.Resolve("kd") finds matches
  │
  ▼
If 1 match: Executor.Execute()
If N matches: Interactive picker
  │
  ▼
Executor builds: "kubectl delete pod my-pod"
  │
  ▼
Check confirm/danger → prompt if needed
  │
  ▼
Execute via shell
  │
  ▼
Store.IncrementUsage()
```

## Database Schema

```sql
CREATE TABLE commands (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    cmd TEXT NOT NULL,
    cwd TEXT,
    tags TEXT DEFAULT '',
    confirm INTEGER DEFAULT 0,
    frequency INTEGER DEFAULT 0,
    last_used DATETIME,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_name ON commands(name);
CREATE INDEX idx_frequency ON commands(frequency DESC);
```

## Frecency Algorithm

```go
func frecencyScore(cmd *Command) int {
    if cmd.LastUsed == nil {
        return cmd.Frequency
    }
    
    hoursSinceUse := time.Since(*cmd.LastUsed).Hours()
    recency := 1.0 / (1.0 + hoursSinceUse/24.0)
    score := float64(cmd.Frequency) * recency * 10
    
    return int(math.Min(score, 10000))
}
```

- Recent usage (within hours) gets high recency factor
- Old usage (days/weeks) gets low recency factor
- Combined with frequency for final score
