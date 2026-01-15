# Testing

## Test Results

| Test | Result |
|------|--------|
| Unit tests (32 tests) | ✅ PASS |
| Add/List/Show/Remove | ✅ |
| Tag filtering | ✅ |
| Fuzzy search (`list -q`) | ✅ |
| Fuzzy run | ✅ |
| Dry-run | ✅ |
| Duplicate handling | ✅ Error message |
| Non-existent command | ✅ Error message |
| Shell init (bash/zsh/fish) | ✅ |
| Unsupported shell | ✅ Error message |
| Special chars (quotes, pipes, semicolon, $variables) | ✅ |
| Long commands | ✅ Truncated in list |
| --cwd flag | ✅ |
| --tags flag | ✅ |
| --confirm flag | ✅ |

## Running Tests

### Unit Tests

```bash
go test ./...
```

Verbose output:

```bash
go test ./... -v
```

### Coverage

```bash
go test ./... -cover
```

## Test Files

| Package | Test File | What's Tested |
|---------|-----------|---------------|
| `internal/model` | `command_test.go` | HasTag, IsDangerous, splitTags |
| `internal/config` | `config_test.go` | DefaultConfig, XDG paths, ensureDir |
| `internal/db` | `store_test.go` | Add, Get, Remove, List, Search, Update, IncrementUsage |
| `internal/resolver` | `matcher_test.go` | Exact/Prefix/Fuzzy match, frecency scoring |
| `internal/executor` | `runner_test.go` | BuildCommand, options |
| `pkg/shell` | `init_test.go` | Init scripts for bash/zsh/fish |

## Unit Test Examples

### Model Tests

```go
func TestCommand_HasTag(t *testing.T) {
    tests := []struct {
        name     string
        tags     string
        tag      string
        expected bool
    }{
        {"empty tags", "", "foo", false},
        {"single tag match", "foo", "foo", true},
        {"multiple tags match", "foo,bar,baz", "bar", true},
    }
    // ...
}
```

### Store Tests

```go
func TestStore_AddAndGet(t *testing.T) {
    store, cleanup := setupTestStore(t)
    defer cleanup()

    cmd := &model.Command{
        Name: "test",
        Cmd:  "echo hello",
    }

    err := store.Add(cmd)
    // assert no error

    got, err := store.Get("test")
    // assert matches input
}
```

### Resolver Tests

```go
func TestResolver_Resolve_ExactMatch(t *testing.T) {
    r := New()
    commands := []*model.Command{
        {Name: "kdp", Cmd: "kubectl delete pod"},
        {Name: "kd", Cmd: "kubectl delete"},
    }

    matches := r.Resolve("kd", commands)
    // assert exact match "kd" is returned
}
```

## Integration Tests

Manual testing script:

```bash
#!/bin/bash
set -e

echo "=== Integration Tests ==="

# Setup
./apotheke add test1 "echo hello"
./apotheke add test2 "echo world" --confirm
./apotheke add test3 "echo foo" --tags test,danger

# Test list
./apotheke list

# Test fuzzy search
./apotheke list -q t1

# Test show
./apotheke show test3

# Test dry-run
./apotheke --dry-run test1

# Test execute
./apotheke test1

# Cleanup
./apotheke rm test1
./apotheke rm test2
./apotheke rm test3

echo "=== All tests passed ==="
```

## Edge Cases Tested

### Duplicate Commands

```bash
./apotheke add dup "echo 1"
./apotheke add dup "echo 2"
# Error: Command 'dup' already exists.
```

### Non-existent Commands

```bash
./apotheke rm nonexistent
# Error: Command 'nonexistent' not found

./apotheke show nonexistent
# Error: Command 'nonexistent' not found

./apotheke xyz123
# Error: No command found matching 'xyz123'
```

### Special Characters

```bash
# Quotes
./apotheke add sq 'echo "hello world"'
# ✅ Works

# Pipes
./apotheke add pipe 'echo hello | grep hello'
# ✅ Works

# Semicolons
./apotheke add semi 'echo a; echo b'
# ✅ Works

# Variables
./apotheke add var 'echo $HOME'
# ✅ Works, expands at runtime
```

### Shell Init

```bash
./apotheke init bash   # ✅
./apotheke init zsh    # ✅
./apotheke init fish   # ✅
./apotheke init powershell
# Error: unsupported shell: powershell
```

## Adding New Tests

### 1. Create test file

```bash
touch internal/mypackage/mypackage_test.go
```

### 2. Write test

```go
package mypackage

import "testing"

func TestMyFeature(t *testing.T) {
    // table-driven test
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {"case 1", "input", "expected"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := MyFeature(tt.input)
            if got != tt.expected {
                t.Errorf("got %q, want %q", got, tt.expected)
            }
        })
    }
}
```

### 3. Run

```bash
go test ./internal/mypackage/...
```
