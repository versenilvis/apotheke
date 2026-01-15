# Tutorial

This tutorial will guide you through using apotheke to bookmark and run commands.

## Installation

```bash
# Build from source
git clone https://github.com/verse/apotheke
cd apotheke
make install

# Add to your shell
echo 'eval "$(apotheke init zsh)"' >> ~/.zshrc
source ~/.zshrc
```

Now you can use `a` as a shortcut for `apotheke`.

## Basic Usage

### 1. Add Your First Command

```bash
a add hello "echo Hello, World!"
```

Output:
```
✓ Added 'hello' → echo Hello, World!
```

### 2. Run the Command

```bash
a hello
```

Output:
```
→ echo Hello, World!
Hello, World!
```

### 3. List All Commands

```bash
a list
# or just
a
```

Output:
```
hello  → echo Hello, World!
```

### 4. Remove a Command

```bash
a rm hello
```

## Real-World Examples

### Docker Commands

```bash
# Add common docker commands
a add dps "docker ps -a"
a add dimg "docker images"
a add drm "docker rm -f"
a add dprune "docker system prune -af" --confirm

# Use them
a dps                    # list containers
a drm my-container       # remove container (args appended)
a dprune                 # asks confirmation before running
```

### Kubernetes Commands

```bash
# Add kubectl commands
a add kgp "kubectl get pods"
a add kgs "kubectl get svc"
a add kdp "kubectl delete pod" --confirm --tags k8s,danger
a add klogs "kubectl logs -f"

# Use them
a kgp -n prod            # kubectl get pods -n prod
a kdp my-pod -n staging  # kubectl delete pod my-pod -n staging
a klogs my-pod           # kubectl logs -f my-pod
```

### Project-Specific Commands

```bash
# Add commands with working directory
a add build "npm run build" --cwd ~/projects/myapp
a add dev "npm run dev" --cwd ~/projects/myapp
a add test "go test ./..." --cwd ~/projects/api

# Run from anywhere
a build     # cd ~/projects/myapp && npm run build
a dev       # cd ~/projects/myapp && npm run dev
```

### Git Workflows

```bash
# Add git commands
a add glog "git log --oneline -20"
a add gpush "git push origin HEAD"
a add gclean "git clean -fd" --confirm

# Use them
a glog
a gpush
```

## Fuzzy Matching

You don't need to type the full command name:

```bash
# If you have: kdp, kgp, kgs
a kd     # matches kdp
a kg     # if multiple matches, shows picker:
         # 1) kgp → kubectl get pods
         # 2) kgs → kubectl get svc
         # Select [1-2]:
```

## Tags and Organization

### Adding Tags

```bash
a add deploy "make deploy" --tags prod,danger
a add staging "make staging" --tags staging
```

### Filtering by Tag

```bash
a list --tag prod      # show only prod-tagged commands
a list --tag danger    # show all dangerous commands
```

### Searching

```bash
a list -q docker       # fuzzy search for "docker"
a list -q kube         # fuzzy search for "kube"
```

## Safety Features

### Confirmation Mode

For dangerous commands:

```bash
a add drop-db "psql -c 'DROP DATABASE prod'" --confirm
a drop-db
# Output:
# → psql -c 'DROP DATABASE prod'
# Execute? [y/N]:
```

### Danger Tag

Commands with `danger` tag always require confirmation:

```bash
a add nuke "rm -rf /" --tags danger
a nuke
# Always asks for confirmation
```

### Dry Run

Preview without executing:

```bash
a --dry-run kdp my-pod
# Output:
# → kubectl delete pod my-pod
# (dry-run mode, not executing)
```

### Skip Confirmation

Override confirmation:

```bash
a -y drop-db
# Skips confirmation prompt
```

## Tips

### Short Names

Use short, memorable names:
- `dps` instead of `docker-ps`
- `kgp` instead of `kubectl-get-pods`

### Consistent Naming

Use prefixes for groups:
- `d*` for docker: `dps`, `drm`, `dimg`
- `k*` for kubectl: `kgp`, `kdp`, `klogs`
- `g*` for git: `glog`, `gpush`, `gclean`

### Arguments

Arguments after the command name are appended:

```bash
a add greet "echo Hello"
a greet World           # → echo Hello World
a greet "John Doe"      # → echo Hello John Doe
```

## Troubleshooting

### Command Not Found

```bash
a xyz
# Error: No command found matching 'xyz'
```

Solution: Check `a list` for available commands.

### Duplicate Name

```bash
a add test "echo 1"
a add test "echo 2"
# Error: Command 'test' already exists. Use 'apotheke rm test' to remove it first.
```

Solution: Remove first, then add again.

### Shell Not Working

If `a` doesn't work:

```bash
# Check if init was added
grep apotheke ~/.zshrc

# Re-add if missing
echo 'eval "$(apotheke init zsh)"' >> ~/.zshrc
source ~/.zshrc
```
