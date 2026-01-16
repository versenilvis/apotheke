package executor

import (
	"testing"

	"github.com/versenilvis/apotheke/internal/model"
)

func TestExecutor_BuildCommand(t *testing.T) {
	e := New()

	tests := []struct {
		name     string
		cmd      *model.Command
		args     []string
		expected string
	}{
		{
			"no args",
			&model.Command{Cmd: "kubectl delete pod"},
			nil,
			"kubectl delete pod",
		},
		{
			"single arg",
			&model.Command{Cmd: "kubectl delete pod"},
			[]string{"my-pod"},
			"kubectl delete pod my-pod",
		},
		{
			"multiple args",
			&model.Command{Cmd: "kubectl delete pod"},
			[]string{"my-pod", "-n", "prod"},
			"kubectl delete pod my-pod -n prod",
		},
		{
			"empty args slice",
			&model.Command{Cmd: "docker ps"},
			[]string{},
			"docker ps",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := e.BuildCommand(tt.cmd, tt.args)
			if got != tt.expected {
				t.Errorf("BuildCommand() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestExecutor_Options(t *testing.T) {
	e := New(WithDryRun(true), WithNoConfirm(true))

	if !e.dryRun {
		t.Error("dryRun should be true")
	}
	if !e.noConfirm {
		t.Error("noConfirm should be true")
	}
}

func TestExecutor_Default(t *testing.T) {
	e := New()

	if e.dryRun {
		t.Error("dryRun should be false by default")
	}
	if e.noConfirm {
		t.Error("noConfirm should be false by default")
	}
}
