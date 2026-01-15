# Apotheke

Apotheke (/a.poˈteː.kə/) is a command bookmark tool like zoxide but for commands.

## Install

**One-liner (recommended):**

```bash
curl -sS https://raw.githubusercontent.com/versenilvis/apotheke/main/install.sh | sh
```

**Or with Go:**

```bash
go install github.com/versenilvis/apotheke/cmd/apotheke@latest
```

**Or build from source:**

```bash
git clone https://github.com/versenilvis/apotheke
cd apotheke
make install
```

## Shell Setup

Add to your shell config to enable the `a` shortcut:

Bash:

```bash
echo 'eval "$(apotheke init bash)"' >> ~/.bashrc
```

Zsh:

```bash
echo 'eval "$(apotheke init zsh)"' >> ~/.zshrc
```

Fish:

```fish
echo 'apotheke init fish | source' >> ~/.config/fish/config.fish
```

## Commands

### add

Add a new command bookmark.

```
a add <name> <command> [flags]
```

| Flag            | Description                                                                           |
| --------------- | ------------------------------------------------------------------------------------- |
| `--cwd <path>`  | Set working directory. Command will `cd` to this path before executing.               |
| `--tags <tags>` | Comma-separated tags for organizing commands. Tag `danger` auto-enables confirmation. |
| `--confirm`     | Always ask for confirmation before running this command.                              |

Examples:

**Recommended** (quote the command):

```bash
a add ax "codex resume 019bc1e9-fb36-7f12-957f-061a532a9265"
a add kdp "kubectl delete pod"
a add deploy "make deploy" --confirm
a add build "npm run build" --cwd ~/project
a add prune "docker system prune -af" --tags docker,danger
```

**Also works** (no quotes, all args after name become the command):

```bash
a add ax codex resume 019bc1e9-fb36-7f12-957f-061a532a9265
a add kdp kubectl delete pod
```

**Do not** (command name with space between characters):

```bash
a add ax shell codex resume 019bc1e9-fb36-7f12-957f-061a532a9265
a add kdp del kubectl delete pod
```

---

### rm

Remove a command bookmark.  
For safety, apotheke only deletes the bookmark if you specify the name exactly.

```
a rm <name>
```

Example:

```bash
a rm kdp
```

**Not working** (fuzzy match):

```bash
a rm k
```

---

### list

List all saved commands.

```
a list [flags]
```

| Flag          | Description                         |
| ------------- | ----------------------------------- |
| `--tag <tag>` | Filter commands by tag.             |
| `-q <query>`  | Search commands by name or content. |

Examples:

```bash
a list                # show all
a list --tag docker   # show only docker-tagged commands
a list -q kubectl     # search for "kubectl"
```

---

### show

Show details of a specific command.

```
a show <name>
```

Example:

```bash
a show kdp
# Output:
# kdp
#   Command:   kubectl delete pod
#   Tags:      k8s,danger
#   Confirm:   true
#   Frequency: 5
#   Last used: 2026-01-15 10:30:00
```

---

### run (default)

Execute a saved command. This is the default action when you type `a <query>`.

```
a <query> [args...]
```

| Flag        | Description                                                        |
| ----------- | ------------------------------------------------------------------ |
| `--dry-run` | Show the command that would be executed, but don't run it.         |
| `-y`        | Skip confirmation prompt (for commands that require confirmation). |

Arguments after `<query>` are appended to the saved command.

Examples:

```bash
a kdp my-pod              # runs: kubectl delete pod my-pod
a kd my-pod               # fuzzy matches "kdp" -> kubectl delete pod my-pod
a --dry-run kdp my-pod    # shows command without running
a -y kdp my-pod           # skip confirmation prompt
```

---

### init

Print shell initialization script. Use with `eval` to enable the `a` function.

```
apotheke init <shell>
```

| Shell  | Description       |
| ------ | ----------------- |
| `bash` | Bash shell script |
| `zsh`  | Zsh shell script  |
| `fish` | Fish shell script |

Example:

```bash
eval "$(apotheke init zsh)"
```

---

### help

Show help for any command.

```
a help [command]
a <command> --help
```

Examples:

```bash
a help           # general help
a help add       # help for add command
a add --help     # same as above
```

## Matching

When you run `a <query>`, the resolver finds the best match:

| Priority | Type   | Description                        |
| -------- | ------ | ---------------------------------- |
| 1        | Exact  | Query exactly matches command name |
| 2        | Prefix | Query is prefix of command name    |
| 3        | Fuzzy  | Query fuzzy-matches command name   |

If multiple commands match, an interactive picker is shown.

Ranking uses **frecency** = frequency × recency. Commands you use often and recently rank higher.

## Safety

Commands are dangerous (unlike `cd`). Safety features:

| Feature            | Description                                               |
| ------------------ | --------------------------------------------------------- |
| `--confirm` flag   | Command always asks "Execute? [y/N]" before running       |
| `danger` tag       | Commands tagged with "danger" always require confirmation |
| `--dry-run`        | Preview command without executing                         |
| Interactive picker | Multiple matches require explicit selection               |

## Data

| Item     | Location                                 |
| -------- | ---------------------------------------- |
| Database | `~/.local/share/apotheke/apotheke.db`    |
| Override | Set `XDG_DATA_HOME` environment variable |

## License

AGPL-3.0 license
