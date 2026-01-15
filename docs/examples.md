# Examples

Real-world examples of using Apotheke.

## Long Multi-Command Sequences

No need to create a `.sh` file! Chain commands with `;` or `&&`:

```bash
# Sequential execution (run all regardless of success)
a add setup "npm install; npm run build; npm run start"

# Conditional execution (stop on first error)
a add deploy "git pull && npm install && npm run build && pm2 restart app"

# Complex CI-like pipeline
a add release "npm run lint && npm run test && npm run build && npm publish"

# Database backup and migrate
a add db-update "pg_dump mydb > backup.sql && rails db:migrate && rails db:seed"

# Clean and rebuild
a add rebuild "rm -rf dist && npm run build && npm run test"
```

**Tip:** Use quotes to preserve the entire command chain:
```bash
a add full-deploy "cd ~/project && git pull origin main && docker-compose down && docker-compose up -d --build"
```

## Docker Commands

```bash
# Add common Docker commands
a add dps "docker ps -a"
a add dimg "docker images"
a add dlog "docker logs -f"
a add dexec "docker exec -it"
a add dprune "docker system prune -af" --tags docker,danger --confirm

# Run
a dps                     # docker ps -a
a dlog my-container       # docker logs -f my-container
a dexec my-container bash # docker exec -it my-container bash
```

## Kubernetes Commands

```bash
# Add kubectl commands
a add k "kubectl"
a add kgp "kubectl get pods"
a add kgs "kubectl get svc"
a add kdp "kubectl delete pod" --tags k8s,danger
a add klog "kubectl logs -f"
a add kexec "kubectl exec -it"
a add kctx "kubectl config use-context"

# Run with arguments
a kgp -n staging              # kubectl get pods -n staging
a kdp my-pod -n production    # kubectl delete pod my-pod -n production
a klog my-pod -c my-container # kubectl logs -f my-pod -c my-container
a kexec my-pod -- bash        # kubectl exec -it my-pod -- bash
```

## Git Workflows

```bash
# Add git commands
a add gs "git status"
a add gd "git diff"
a add gc "git commit -m"
a add gp "git push origin"
a add gl "git log --oneline -10"
a add gco "git checkout"
a add gb "git branch"
a add gpf "git push --force-with-lease" --tags git,danger --confirm

# Run
a gs                # git status
a gc "fix bug"      # git commit -m "fix bug"
a gp main           # git push origin main
a gco feature/new   # git checkout feature/new
```

## Project-Specific Commands

```bash
# Development
a add dev "npm run dev" --cwd ~/projects/myapp
a add build "npm run build" --cwd ~/projects/myapp
a add test "npm run test" --cwd ~/projects/myapp

# Deployment
a add deploy-staging "make deploy ENV=staging" --confirm
a add deploy-prod "make deploy ENV=production" --tags danger --confirm

# Database
a add db-migrate "rails db:migrate"
a add db-reset "rails db:drop db:create db:migrate db:seed" --tags danger --confirm
```

## SSH and Remote

```bash
# Add SSH shortcuts
a add ssh-prod "ssh user@production.server.com"
a add ssh-staging "ssh user@staging.server.com"
a add ssh-db "ssh -L 5432:localhost:5432 user@db.server.com"

# SCP files
a add scp-logs "scp user@server:/var/log/app.log ."
```

## System Administration

```bash
# System commands
a add disk "df -h"
a add mem "free -h"
a add ports "netstat -tulpn"
a add proc "ps aux | grep"

# With danger tag for destructive commands
a add rm-cache "rm -rf ~/.cache/*" --tags system,danger --confirm
a add restart-nginx "sudo systemctl restart nginx" --confirm
```

## Using Tags

```bash
# Add commands with tags
a add dps "docker ps" --tags docker
a add kgp "kubectl get pods" --tags k8s
a add deploy "make deploy" --tags deploy,danger

# Filter by tag
a list --tag docker    # Show only docker commands
a list --tag k8s       # Show only k8s commands
a list --tag danger    # Show all dangerous commands
```

## Using Confirmation

```bash
# Commands that always ask before running
a add drop-db "psql -c 'DROP DATABASE mydb'" --confirm
a add rm-all "rm -rf *" --tags danger --confirm

# Skip confirmation with -y flag
a -y drop-db           # Execute without asking
```

## Dry Run

```bash
# Preview command without executing
a --dry-run kgp -n production
# Output: → kubectl get pods -n production
# (does not execute)

# Useful for checking argument expansion
a --dry-run dexec my-container bash
# Output: → docker exec -it my-container bash
```

## Search and Filter

```bash
# Search by name or command content
a list -q docker       # Find commands containing "docker"
a list -q kubectl      # Find commands containing "kubectl"
a list -q deploy       # Find commands containing "deploy"

# Fuzzy matching when running
a kd my-pod            # Matches "kdp" → kubectl delete pod my-pod
a dp                   # Matches "dps" → docker ps -a
```

## Working Directory

```bash
# Commands that run in specific directory
a add front-dev "npm run dev" --cwd ~/projects/frontend
a add back-dev "go run main.go" --cwd ~/projects/backend
a add docs-build "mkdocs build" --cwd ~/projects/docs

# When you run, it automatically cd to that directory first
a front-dev            # cd ~/projects/frontend && npm run dev
```
