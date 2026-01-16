# Concepts

## Matching

When you run `a <query>`, the resolver finds the best match:

| Priority | Type   | Description                        |
| -------- | ------ | ---------------------------------- |
| 1        | Exact  | Query exactly matches command name |
| 2        | Prefix | Query is prefix of command name    |
| 3        | Fuzzy  | Query fuzzy-matches command name   |

> [!IMPORTANT]
> If multiple commands match, an interactive picker is shown.
> 
> Ranking uses **frecency** = frequency × recency. Commands you use often and recently rank higher.

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