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

<details>
<summary><b>add</b> - Add a new command bookmark</summary>

```
a add <name> <command> [flags]
```

| Flag | Description |
|------|-------------|
| `--cwd <path>` | Working directory |
| `--tags <tags>` | Comma-separated tags (`danger` auto-enables confirmation) |
| `--confirm` | Require confirmation |

```bash
a add kdp "kubectl delete pod"
a add build "npm run build" --cwd ~/project
a add prune "docker system prune -af" --tags docker,danger
```

> [!NOTE]
> Quote commands with special characters: `-`, `|`, `>`, `$`

</details>

<details>
<summary><b>rm</b> - Remove a bookmark</summary>

```bash
a rm <name>
```

> [!NOTE]
> Must use exact name. Fuzzy match not supported for safety.

</details>

<details>
<summary><b>list</b> - List all commands</summary>

```bash
a list                # all
a list --tag docker   # filter by tag
a list -q kubectl     # search
a                     # shortcut
```

</details>

<details>
<summary><b>show</b> - Show command details</summary>

```bash
a show kdp
```

</details>

<details>
<summary><b>run</b> - Execute command (default)</summary>

```bash
a kdp my-pod           # run with args
a kd my-pod            # fuzzy match
a --dry-run kdp        # preview only
a -y kdp               # skip confirmation
```

</details>

<details>
<summary><b>init</b> - Shell setup</summary>

```bash
eval "$(apotheke init zsh)"   # zsh
eval "$(apotheke init bash)"  # bash
apotheke init fish | source   # fish
```

</details>

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
