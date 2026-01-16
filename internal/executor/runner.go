package executor

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/versenilvis/apotheke/internal/model"
)

type Executor struct {
	dryRun    bool
	noConfirm bool
}

type Option func(*Executor)

func WithDryRun(dryRun bool) Option {
	return func(e *Executor) {
		e.dryRun = dryRun
	}
}

func WithNoConfirm(noConfirm bool) Option {
	return func(e *Executor) {
		e.noConfirm = noConfirm
	}
}

func New(opts ...Option) *Executor {
	e := &Executor{}
	for _, opt := range opts {
		opt(e)
	}
	return e
}

func (e *Executor) BuildCommand(cmd *model.Command, args []string) string {
	if len(args) == 0 {
		return cmd.Cmd
	}
	return cmd.Cmd + " " + strings.Join(args, " ")
}

func (e *Executor) Execute(cmd *model.Command, args []string) error {
	fullCmd := e.BuildCommand(cmd, args)

	cyan := color.New(color.FgCyan, color.Bold)
	cyan.Fprintf(os.Stderr, "→ %s\n", fullCmd)

	if e.dryRun {
		yellow := color.New(color.FgYellow)
		yellow.Fprintln(os.Stderr, "(dry-run mode, not executing)")
		return nil
	}

	needsConfirm := cmd.Confirm || cmd.IsDangerous()
	if needsConfirm && !e.noConfirm {
		if !confirm("Execute?") {
			yellow := color.New(color.FgYellow)
			yellow.Fprintln(os.Stderr, "Cancelled.")
			return nil
		}
	}

	return runCommand(fullCmd, cmd.Cwd)
}

func (e *Executor) PrintCommand(cmd *model.Command, args []string) {
	fullCmd := e.BuildCommand(cmd, args)
	fmt.Println(fullCmd)
}

func runCommand(cmdStr string, cwd *string) error {
	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "/bin/sh"
	}

	cmd := exec.Command(shell, "-c", cmdStr)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if cwd != nil && *cwd != "" {
		cmd.Dir = *cwd
	}

	return cmd.Run()
}

func confirm(prompt string) bool {
	gray := color.New(color.FgHiBlack)
	gray.Fprintf(os.Stderr, "%s [y/N]: ", prompt)

	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		return false
	}

	response = strings.TrimSpace(strings.ToLower(response))
	return response == "y" || response == "yes"
}
