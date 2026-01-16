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

[Tiếng Việt](./README.VN.md) | English

> [!WARNING]
> **Currently, Apotheke is under development.**

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

## Shell setup

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
> [!IMPORTANT]
> - Please read from the [docs](./docs/commands.md)

## Tutorial
> [!IMPORTANT]
> - Please read from the [docs](./docs/tutorial.md)

## More examples
> [!IMPORTANT]
> - Please read from the [docs](./docs/examples.md)

## Uninstall

```bash
rm ~/.local/bin/apotheke
rm -rf ~/.local/share/apotheke
```

Remove the `eval` line from your shell config file.

<details>
  <summary><h2>FAQ</h2></summary>
  
### Q: Why did you build this?  
A: To store 'codex resume' and 'cursor-agent --resume=' commands that I always forget after turn off the terminal.

### Q: How is Apotheke different from shell aliases?  
A: Shell aliases are static and require editing config files. Apotheke offers:
- Fuzzy matching (a kd → kubectl delete pod)
- Frecency ranking (frequently used commands rank higher)
- Tags and organization
- Safety confirmations for dangerous commands
- Argument appending (a kdp my-pod → kubectl delete pod my-pod)

### Q: How is it different from shell history?  
A: History searches all commands. Apotheke only stores commands you explicitly bookmark with meaningful names.

### Q: Does it work on Windows?  
A: Yes, but shell integration requires Git Bash, WSL, or PowerShell with custom setup.

### Q: Can I sync across machines?  
A: Not built-in yet. Maybe in the future. Or you can manually copy the database file.

### Q: What does "Apotheke" mean?  
A: Well, I just asked ChatGPT what is "storage" in Acient Greek and it said 'Apotheke'.
</details>

## License

[AGPL-3.0 license](./LICENSE)


## Contributing

Please follow my [Contributing](.github/CONTRIBUTING.md) when you make a pull request.
