package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/versenilvis/apotheke/internal/config"
	"github.com/versenilvis/apotheke/internal/db"
	"github.com/versenilvis/apotheke/internal/executor"
	"github.com/versenilvis/apotheke/internal/model"
	"github.com/versenilvis/apotheke/internal/resolver"
	"github.com/versenilvis/apotheke/pkg/shell"
)

var version = "0.1.1"

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "apotheke [query] [args...]",
	Short: "Command bookmark tool - like zoxide for commands",
	Long: `Apotheke is a CLI tool for bookmarking and quickly executing commands.
Like zoxide helps you navigate directories, apotheke helps you run commands.

Examples:
  apotheke add kdp "kubectl delete pod"
  apotheke kdp my-pod
  apotheke list`,
	Version: version,
	Args:    cobra.ArbitraryArgs,
	Run:     runRoot,
}

var (
	dryRun    bool
	noConfirm bool
)

func init() {
	rootCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Show command without executing")
	rootCmd.Flags().BoolVarP(&noConfirm, "yes", "y", false, "Skip confirmation prompts")

	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(rmCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(showCmd)
	rootCmd.AddCommand(resolveCmd)
	rootCmd.AddCommand(initCmd)
}

func runRoot(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		listCmd.Run(cmd, args)
		return
	}

	store, err := getStore()
	if err != nil {
		errorExit("Failed to open database: %v", err)
	}
	defer store.Close()

	query := args[0]
	cmdArgs := args[1:]

	commands, err := store.GetAll()
	if err != nil {
		errorExit("Failed to list commands: %v", err)
	}

	if len(commands) == 0 {
		errorExit("No commands found. Add one with: apotheke add <name> <command>")
	}

	res := resolver.New()
	matches := res.Resolve(query, commands)

	if len(matches) == 0 {
		errorExit("No command found matching '%s'", query)
	}

	var selected *model.Command

	if len(matches) == 1 {
		selected = matches[0].Command
	} else {
		selected = pickCommand(matches)
		if selected == nil {
			return
		}
	}

	exec := executor.New(
		executor.WithDryRun(dryRun),
		executor.WithNoConfirm(noConfirm),
	)

	if err := exec.Execute(selected, cmdArgs); err != nil {
		os.Exit(1)
	}

	store.IncrementUsage(selected.Name)
}

var addCmd = &cobra.Command{
	Use:   "add <name> <command>",
	Short: "Add a new command bookmark",
	Long: `Add a new command bookmark with the given name.

Examples:
  apotheke add kdp "kubectl delete pod"
  apotheke add dps "docker ps -a" --confirm
  apotheke add deploy "make deploy" --cwd ~/project --tags deploy,danger`,
	Args: cobra.MinimumNArgs(2),
	Run:  runAdd,
}

var (
	addCwd     string
	addTags    string
	addConfirm bool
)

func init() {
	addCmd.Flags().StringVar(&addCwd, "cwd", "", "Working directory for the command")
	addCmd.Flags().StringVar(&addTags, "tags", "", "Comma-separated tags")
	addCmd.Flags().BoolVar(&addConfirm, "confirm", false, "Require confirmation before running")
}

func runAdd(cmd *cobra.Command, args []string) {
	store, err := getStore()
	if err != nil {
		errorExit("Failed to open database: %v", err)
	}
	defer store.Close()

	name := args[0]
	cmdStr := strings.Join(args[1:], " ")

	command := &model.Command{
		Name:    name,
		Cmd:     cmdStr,
		Tags:    addTags,
		Confirm: addConfirm,
	}

	if addCwd != "" {
		command.Cwd = &addCwd
	}

	if err := store.Add(command); err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint") {
			errorExit("Command '%s' already exists. Use 'apotheke rm %s' to remove it first.", name, name)
		}
		errorExit("Failed to add command: %v", err)
	}

	green := color.New(color.FgGreen)
	green.Printf("✓ Added '%s' → %s\n", name, cmdStr)
}

var rmCmd = &cobra.Command{
	Use:     "rm <name>",
	Aliases: []string{"remove", "delete"},
	Short:   "Remove a command bookmark",
	Args:    cobra.ExactArgs(1),
	Run:     runRm,
}

func runRm(cmd *cobra.Command, args []string) {
	store, err := getStore()
	if err != nil {
		errorExit("Failed to open database: %v", err)
	}
	defer store.Close()

	name := args[0]

	existing, err := store.Get(name)
	if err != nil || existing == nil {
		errorExit("Command '%s' not found", name)
	}

	if err := store.Remove(name); err != nil {
		errorExit("Failed to remove command: %v", err)
	}

	yellow := color.New(color.FgYellow)
	yellow.Printf("✓ Removed '%s'\n", name)
}

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List all command bookmarks",
	Run:     runList,
}

var (
	listTag   string
	listQuery string
)

func init() {
	listCmd.Flags().StringVar(&listTag, "tag", "", "Filter by tag")
	listCmd.Flags().StringVarP(&listQuery, "query", "q", "", "Search query")
}

func runList(cmd *cobra.Command, args []string) {
	store, err := getStore()
	if err != nil {
		errorExit("Failed to open database: %v", err)
	}
	defer store.Close()

	var commands []*model.Command

	if listQuery != "" {
		allCommands, err := store.GetAll()
		if err != nil {
			errorExit("Failed to list commands: %v", err)
		}
		res := resolver.New()
		matches := res.Resolve(listQuery, allCommands)
		for _, m := range matches {
			commands = append(commands, m.Command)
		}
	} else {
		commands, err = store.List(listTag)
		if err != nil {
			errorExit("Failed to list commands: %v", err)
		}
	}

	if len(commands) == 0 {
		fmt.Println("No commands found.")
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	cyan := color.New(color.FgCyan)
	gray := color.New(color.FgHiBlack)

	for _, c := range commands {
		name := cyan.Sprint(c.Name)
		cmdStr := c.Cmd
		if len(cmdStr) > 60 {
			cmdStr = cmdStr[:57] + "..."
		}

		extra := ""
		if c.Tags != "" {
			extra = gray.Sprintf(" [%s]", c.Tags)
		}
		if c.Confirm {
			extra += gray.Sprint(" (confirm)")
		}

		fmt.Fprintf(w, "%s\t→ %s%s\n", name, cmdStr, extra)
	}
	w.Flush()
}

var showCmd = &cobra.Command{
	Use:   "show <name>",
	Short: "Show details of a command bookmark",
	Args:  cobra.ExactArgs(1),
	Run:   runShow,
}

func runShow(cmd *cobra.Command, args []string) {
	store, err := getStore()
	if err != nil {
		errorExit("Failed to open database: %v", err)
	}
	defer store.Close()

	name := args[0]
	command, err := store.Get(name)
	if err != nil {
		errorExit("Command '%s' not found", name)
	}

	cyan := color.New(color.FgCyan, color.Bold)
	gray := color.New(color.FgHiBlack)

	cyan.Printf("%s\n", command.Name)
	fmt.Printf("  Command:   %s\n", command.Cmd)
	if command.Cwd != nil {
		fmt.Printf("  Directory: %s\n", *command.Cwd)
	}
	if command.Tags != "" {
		fmt.Printf("  Tags:      %s\n", command.Tags)
	}
	fmt.Printf("  Confirm:   %v\n", command.Confirm)
	fmt.Printf("  Frequency: %d\n", command.Frequency)
	if command.LastUsed != nil {
		fmt.Printf("  Last used: %s\n", command.LastUsed.Format("2006-01-02 15:04:05"))
	}
	gray.Printf("  Created:   %s\n", command.CreatedAt.Format("2006-01-02 15:04:05"))
}

var resolveCmd = &cobra.Command{
	Use:    "resolve <query> [args...]",
	Short:  "Resolve a query to a command (for shell integration)",
	Hidden: true,
	Args:   cobra.MinimumNArgs(1),
	Run:    runResolve,
}

func runResolve(cmd *cobra.Command, args []string) {
	store, err := getStore()
	if err != nil {
		os.Exit(1)
	}
	defer store.Close()

	query := args[0]
	cmdArgs := args[1:]

	commands, err := store.GetAll()
	if err != nil || len(commands) == 0 {
		os.Exit(1)
	}

	res := resolver.New()
	matches := res.Resolve(query, commands)

	if len(matches) == 0 {
		errorExit("No command found matching '%s'", query)
	}

	var selected *model.Command

	if len(matches) == 1 {
		selected = matches[0].Command
	} else {
		selected = pickCommand(matches)
		if selected == nil {
			os.Exit(1)
		}
	}

	needsConfirm := selected.Confirm || selected.IsDangerous()
	if needsConfirm {
		cyan := color.New(color.FgCyan, color.Bold)
		exec := executor.New()
		fullCmd := exec.BuildCommand(selected, cmdArgs)
		cyan.Fprintf(os.Stderr, "→ %s\n", fullCmd)

		if !confirm("Execute?") {
			yellow := color.New(color.FgYellow)
			yellow.Fprintln(os.Stderr, "Cancelled.")
			os.Exit(1)
		}
	}

	exec := executor.New()
	exec.PrintCommand(selected, cmdArgs)

	store.IncrementUsage(selected.Name)
}

var initCmd = &cobra.Command{
	Use:   "init <shell>",
	Short: "Print shell initialization script",
	Long: `Print the initialization script for your shell.
Add to your shell config file:

Bash:  eval "$(apotheke init bash)"
Zsh:   eval "$(apotheke init zsh)"
Fish:  apotheke init fish | source`,
	Args:      cobra.ExactArgs(1),
	ValidArgs: []string{"bash", "zsh", "fish"},
	Run:       runInit,
}

func runInit(cmd *cobra.Command, args []string) {
	script, err := shell.Init(args[0])
	if err != nil {
		errorExit("%v", err)
	}
	fmt.Print(script)
}

func getStore() (*db.Store, error) {
	cfg, err := config.DefaultConfig()
	if err != nil {
		return nil, err
	}
	return db.New(cfg.DBPath)
}

func errorExit(format string, args ...interface{}) {
	red := color.New(color.FgRed)
	red.Fprintf(os.Stderr, "Error: "+format+"\n", args...)
	os.Exit(1)
}

func pickCommand(matches []resolver.Match) *model.Command {
	yellow := color.New(color.FgYellow)
	cyan := color.New(color.FgCyan)

	yellow.Fprintln(os.Stderr, "Multiple matches found:")
	for i, m := range matches {
		if i >= 10 {
			gray := color.New(color.FgHiBlack)
			gray.Fprintf(os.Stderr, "  ... and %d more\n", len(matches)-10)
			break
		}
		fmt.Fprintf(os.Stderr, "  %s) %s → %s\n",
			cyan.Sprint(i+1),
			cyan.Sprint(m.Command.Name),
			m.Command.Cmd,
		)
	}

	fmt.Fprint(os.Stderr, "Select [1-", len(matches), "]: ")
	var choice int
	_, err := fmt.Scanln(&choice)
	if err != nil || choice < 1 || choice > len(matches) {
		return nil
	}

	return matches[choice-1].Command
}

func confirm(prompt string) bool {
	gray := color.New(color.FgHiBlack)
	gray.Fprintf(os.Stderr, "%s [y/N]: ", prompt)

	var response string
	fmt.Scanln(&response)
	response = strings.TrimSpace(strings.ToLower(response))
	return response == "y" || response == "yes"
}
