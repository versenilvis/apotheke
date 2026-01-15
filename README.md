<div align="center">
  <h1>Apotheke</h1>
  <p><b>Apotheke (/a.poˈteː.kə/) is a smart alias command tool.</b></p>
</div>

<!-- <div align="center">
  <a href="https://github.com/versenilvis/apotheke/stargazers">
    <img src="https://img.shields.io/github/stars/versenilvis/apotheke?style=for-the-badge&logo=github&color=E3B341&logoColor=D9E0EE&labelColor=000000" alt="GitHub stars">
  </a>
</div> -->
<div align="center">
  
  [![Stars](https://img.shields.io/badge/Stars-000?style=for-the-badge&logo=github&logoColor=white&labelColor=000000)](https://github.com/versenilvis/apotheke/stargazers)
  [![Twitter](https://img.shields.io/badge/Follow_me-000?style=for-the-badge&logo=x&logoColor=white&labelColor=000000)](https://x.com/VerseNilVis)

</div>


<div align="center">

  [![License: AGPL-3.0 license](https://img.shields.io/badge/License-AGPL_v3-blue?style=for-the-badge&logo=github&logoColor=white)](./LICENSE.md)
  [![Status](https://img.shields.io/badge/status-beta-yellow?style=for-the-badge&logo=github&logoColor=white)]()
  [![Documentation](https://img.shields.io/badge/docs-available-brightgreen?style=for-the-badge&logo=github&logoColor=white)](./docs/README.md)
  [![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen?style=for-the-badge&logo=github&logoColor=white)](./CONTRIBUTING.md)

</div>

## Preview

<img width="2531" height="742" alt="image" src="https://github.com/user-attachments/assets/cba2cc36-aa43-468d-994e-fc86cfb77c4f" />

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

> [!WARNING]
> **Currently, Apotheke is in development.**

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

**Must** (quote the command when it has special characters):

```bash
a add apo "curl -sS https://raw.githubusercontent.com/versenilvis/apotheke/main/install.sh | sh"
```

---

### rm

Remove a command bookmark.  

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

Examples:

```bash
a list                # show all
a list --tag docker   # show only docker-tagged commands
a list -q kubectl     # search for "kubectl"
```
> [!TIP]
> Or just type 'a' to list all commands.

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

## Database

| Item     | Location                                 |
| -------- | ---------------------------------------- |
| Database | `~/.local/share/apotheke/apotheke.db`    |
| Override | Set `XDG_DATA_HOME` environment variable |

## FAQ

**Q: How is Apotheke different from shell aliases?**
A: Shell aliases are static and require editing config files. Apotheke offers:
- Fuzzy matching (a kd → kubectl delete pod)
- Frecency ranking (frequently used commands rank higher)
- Tags and organization
- Safety confirmations for dangerous commands
- Argument appending (a kdp my-pod → kubectl delete pod my-pod)

**Q: How is it different from shell history?**
A: History searches all commands. Apotheke only stores commands you explicitly bookmark with meaningful names.

**Q: Does it work on Windows?**
A: Yes, but shell integration requires Git Bash, WSL, or PowerShell with custom setup.

**Q: Can I sync across machines?**
A: Not built-in yet. Maybe in the future. Or you can manually copy the database file.

## License

[AGPL-3.0 license](./LICENSE)


## Contributing

Please follow our [Contributing](.github/CONTRIBUTING.md) when you make a pull request.
