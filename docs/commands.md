# Commands

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

<details>
  <summary>Examples:</summary>

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

**Must** (quote the command when it has special characters):

```bash
a add apo "curl -sS https://raw.githubusercontent.com/versenilvis/apotheke/main/install.sh | sh"
```
</details>

---

### rm

Remove a command bookmark.  

```
a rm <name>
```
<details>
  <summary>Examples:</summary>

```bash
a rm kdp
```

**Not working** (fuzzy match):

```bash
a rm k
```
</details>

> [!NOTE]
> For safety, apotheke only deletes the bookmark if you specify the name exactly.

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

<details>
  <summary>Examples:</summary>

```bash
a list                # show all
a list --tag docker   # show only docker-tagged commands
a list -q kubectl     # search for "kubectl"
```
</details>

> [!TIP]
> Just type 'a' to list all commands.

---

### show

Show details of a specific command.

```
a show <name>
```

<details>
  <summary>Examples:</summary>

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
</details>

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

<details>
  <summary>Examples:</summary>

```bash
a kdp my-pod              # runs: kubectl delete pod my-pod
a kd my-pod               # fuzzy matches "kdp" -> kubectl delete pod my-pod
a --dry-run kdp my-pod    # shows command without running
a -y kdp my-pod           # skip confirmation prompt
```
</details>

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

<details>
  <summary>Examples:</summary>

```bash
eval "$(apotheke init zsh)"
```

</details>

---