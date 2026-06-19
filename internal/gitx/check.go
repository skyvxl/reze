package gitx

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	ErrGitNotFound      = errors.New("git executable not found")
	ErrNotGitRepository = errors.New("not a git repository")
	ErrInvalidWorkDir   = errors.New("invalid working directory")
)

func CheckGit(ctx context.Context, workDir string) error {
	cleanDir := filepath.Clean(workDir)
	info, err := os.Stat(cleanDir)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrInvalidWorkDir, cleanDir)
	}
	if !info.IsDir() {
		return fmt.Errorf("%w: %s is not a directory", ErrInvalidWorkDir, cleanDir)
	}
	if _, err := exec.LookPath("git"); err != nil {
		if errors.Is(err, exec.ErrNotFound) {
			return ErrGitNotFound
		}
		return fmt.Errorf("look up git executable: %w", err)
	}
	out, err := Run(ctx, cleanDir, "rev-parse", "--is-inside-work-tree")
	if err != nil {
		return fmt.Errorf("%w: %s", ErrNotGitRepository, out)
	}
	if out != "true" {
		return ErrNotGitRepository
	}
	return nil
}
