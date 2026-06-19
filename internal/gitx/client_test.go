package gitx

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestClient_Discover_ReturnsRepositoryRoot(t *testing.T) {
	tempDir := t.TempDir()
	client, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	cleanDir := filepath.Clean(tempDir)
	_, err = client.Run(context.Background(), cleanDir, "init")
	if err != nil {
		t.Fatal(err)
	}
	nestedDir := filepath.Join(cleanDir, "repo", "internal", "app")
	err = os.MkdirAll(nestedDir, 0o755)
	if err != nil {
		t.Fatal(err)
	}
	repo, err := client.Discover(context.Background(), nestedDir)
	if err != nil {
		t.Fatal(err)
	}
	gotRoot := filepath.Clean(filepath.FromSlash(repo.Root))
	wantRoot := filepath.Clean(filepath.FromSlash(tempDir))
	if gotRoot != wantRoot {
		t.Errorf("discover.Root = %s; want %s", gotRoot, wantRoot)
	}
}

func TestClient_Discover_ReturnsNotGitRepository(t *testing.T) {
	tempDir := t.TempDir()
	client, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	cleanDir := filepath.Clean(tempDir)
	_, err = client.Discover(context.Background(), cleanDir)
	if err == nil {
		t.Fatal("discover should return an error")
	}
	if !errors.Is(err, ErrNotGitRepository) {
		t.Fatal("discover should return a not git repository")
	}
}
