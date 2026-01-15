package shell

import (
	"strings"
	"testing"
)

func TestInit_Bash(t *testing.T) {
	script, err := Init("bash")
	if err != nil {
		t.Fatalf("Init(bash) error = %v", err)
	}

	if !strings.Contains(script, "function a()") {
		t.Error("bash script should contain function a()")
	}
	if !strings.Contains(script, "apotheke") {
		t.Error("bash script should reference apotheke")
	}
}

func TestInit_Zsh(t *testing.T) {
	script, err := Init("zsh")
	if err != nil {
		t.Fatalf("Init(zsh) error = %v", err)
	}

	if !strings.Contains(script, "function a()") {
		t.Error("zsh script should contain function a()")
	}
	if !strings.Contains(script, "apotheke") {
		t.Error("zsh script should reference apotheke")
	}
}

func TestInit_Fish(t *testing.T) {
	script, err := Init("fish")
	if err != nil {
		t.Fatalf("Init(fish) error = %v", err)
	}

	if !strings.Contains(script, "function a") {
		t.Error("fish script should contain function a")
	}
	if !strings.Contains(script, "apotheke") {
		t.Error("fish script should reference apotheke")
	}
}

func TestInit_Unsupported(t *testing.T) {
	_, err := Init("powershell")
	if err == nil {
		t.Error("Init(powershell) should return error")
	}
}

func TestInitBash_Content(t *testing.T) {
	script := InitBash()

	required := []string{
		"add|rm|remove|list|ls|edit|show|init|help",
		"command apotheke",
	}

	for _, r := range required {
		if !strings.Contains(script, r) {
			t.Errorf("bash script should contain %q", r)
		}
	}
}

func TestInitZsh_Content(t *testing.T) {
	script := InitZsh()

	required := []string{
		"add|rm|remove|list|ls|edit|show|init|help",
		"command apotheke",
	}

	for _, r := range required {
		if !strings.Contains(script, r) {
			t.Errorf("zsh script should contain %q", r)
		}
	}
}

func TestInitFish_Content(t *testing.T) {
	script := InitFish()

	required := []string{
		"case add rm remove list ls edit show init help",
		"command apotheke",
	}

	for _, r := range required {
		if !strings.Contains(script, r) {
			t.Errorf("fish script should contain %q", r)
		}
	}
}
