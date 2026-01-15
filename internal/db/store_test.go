package db

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/verse/apotheke/internal/model"
)

func setupTestStore(t *testing.T) (*Store, func()) {
	tmpDir := t.TempDir()
	dbPath := filepath.Join(tmpDir, "test.db")

	store, err := New(dbPath)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	cleanup := func() {
		store.Close()
		os.Remove(dbPath)
	}

	return store, cleanup
}

func TestStore_AddAndGet(t *testing.T) {
	store, cleanup := setupTestStore(t)
	defer cleanup()

	cmd := &model.Command{
		Name:    "test",
		Cmd:     "echo hello",
		Tags:    "test,example",
		Confirm: true,
	}

	err := store.Add(cmd)
	if err != nil {
		t.Fatalf("Add() error = %v", err)
	}

	got, err := store.Get("test")
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}

	if got.Name != cmd.Name {
		t.Errorf("Name = %q, want %q", got.Name, cmd.Name)
	}
	if got.Cmd != cmd.Cmd {
		t.Errorf("Cmd = %q, want %q", got.Cmd, cmd.Cmd)
	}
	if got.Tags != cmd.Tags {
		t.Errorf("Tags = %q, want %q", got.Tags, cmd.Tags)
	}
	if got.Confirm != cmd.Confirm {
		t.Errorf("Confirm = %v, want %v", got.Confirm, cmd.Confirm)
	}
}

func TestStore_AddDuplicate(t *testing.T) {
	store, cleanup := setupTestStore(t)
	defer cleanup()

	cmd := &model.Command{Name: "test", Cmd: "echo hello"}

	err := store.Add(cmd)
	if err != nil {
		t.Fatalf("Add() error = %v", err)
	}

	err = store.Add(cmd)
	if err == nil {
		t.Error("Add() should return error for duplicate")
	}
}

func TestStore_Remove(t *testing.T) {
	store, cleanup := setupTestStore(t)
	defer cleanup()

	cmd := &model.Command{Name: "test", Cmd: "echo hello"}
	store.Add(cmd)

	err := store.Remove("test")
	if err != nil {
		t.Fatalf("Remove() error = %v", err)
	}

	_, err = store.Get("test")
	if err == nil {
		t.Error("Get() should return error after Remove()")
	}
}

func TestStore_List(t *testing.T) {
	store, cleanup := setupTestStore(t)
	defer cleanup()

	store.Add(&model.Command{Name: "cmd1", Cmd: "echo 1", Tags: "a"})
	store.Add(&model.Command{Name: "cmd2", Cmd: "echo 2", Tags: "b"})
	store.Add(&model.Command{Name: "cmd3", Cmd: "echo 3", Tags: "a,b"})

	all, err := store.List("")
	if err != nil {
		t.Fatalf("List() error = %v", err)
	}
	if len(all) != 3 {
		t.Errorf("List() returned %d items, want 3", len(all))
	}

	withTagA, err := store.List("a")
	if err != nil {
		t.Fatalf("List(a) error = %v", err)
	}
	if len(withTagA) != 2 {
		t.Errorf("List(a) returned %d items, want 2", len(withTagA))
	}
}

func TestStore_Search(t *testing.T) {
	store, cleanup := setupTestStore(t)
	defer cleanup()

	store.Add(&model.Command{Name: "kubectl-delete", Cmd: "kubectl delete"})
	store.Add(&model.Command{Name: "kubectl-get", Cmd: "kubectl get"})
	store.Add(&model.Command{Name: "docker-ps", Cmd: "docker ps"})

	results, err := store.Search("kubectl")
	if err != nil {
		t.Fatalf("Search() error = %v", err)
	}
	if len(results) != 2 {
		t.Errorf("Search(kubectl) returned %d items, want 2", len(results))
	}
}

func TestStore_IncrementUsage(t *testing.T) {
	store, cleanup := setupTestStore(t)
	defer cleanup()

	store.Add(&model.Command{Name: "test", Cmd: "echo hello"})

	for i := 0; i < 5; i++ {
		err := store.IncrementUsage("test")
		if err != nil {
			t.Fatalf("IncrementUsage() error = %v", err)
		}
	}

	cmd, _ := store.Get("test")
	if cmd.Frequency != 5 {
		t.Errorf("Frequency = %d, want 5", cmd.Frequency)
	}
	if cmd.LastUsed == nil {
		t.Error("LastUsed should not be nil")
	}
}

func TestStore_Update(t *testing.T) {
	store, cleanup := setupTestStore(t)
	defer cleanup()

	store.Add(&model.Command{Name: "test", Cmd: "echo hello"})

	cwd := "/tmp"
	updated := &model.Command{
		Name:    "test",
		Cmd:     "echo world",
		Cwd:     &cwd,
		Tags:    "updated",
		Confirm: true,
	}

	err := store.Update(updated)
	if err != nil {
		t.Fatalf("Update() error = %v", err)
	}

	got, _ := store.Get("test")
	if got.Cmd != "echo world" {
		t.Errorf("Cmd = %q, want %q", got.Cmd, "echo world")
	}
	if got.Tags != "updated" {
		t.Errorf("Tags = %q, want %q", got.Tags, "updated")
	}
	if !got.Confirm {
		t.Error("Confirm should be true")
	}
}

func TestStore_GetAll(t *testing.T) {
	store, cleanup := setupTestStore(t)
	defer cleanup()

	store.Add(&model.Command{Name: "a", Cmd: "echo a"})
	store.Add(&model.Command{Name: "b", Cmd: "echo b"})

	all, err := store.GetAll()
	if err != nil {
		t.Fatalf("GetAll() error = %v", err)
	}
	if len(all) != 2 {
		t.Errorf("GetAll() returned %d items, want 2", len(all))
	}
}
