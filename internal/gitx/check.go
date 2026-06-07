package gitx

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

var (
	ErrGitNotFound      = errors.New("git executable not found")
	ErrNotGitRepository = errors.New("not a git repository")
)

func CheckGit(ctx context.Context, workDir string) error {
	if _, err := exec.LookPath("git"); err != nil {
		if errors.Is(err, exec.ErrNotFound) {
			return ErrGitNotFound
		}
		return fmt.Errorf("look up git executable: %w", err)
	}
	cmd := exec.CommandContext(ctx, "git", "rev-parse", "--is-inside-work-tree")
	cmd.Dir = workDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%w: %s", ErrNotGitRepository, strings.TrimSpace(string(output)))
	}
	if strings.TrimSpace(string(output)) != "true" {
		return ErrNotGitRepository
	}
	return nil
}
