package gitx

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/skyvxl/reze/internal/guard"
)

var (
	ErrGitNotFound      = errors.New("git executable not found")
	ErrNotGitRepository = errors.New("not a git repository")
	ErrInvalidWorkDir   = errors.New("invalid working directory")
)

type Client struct {
	gitPath string
}

func NewClient() (*Client, error) {
	git, err := exec.LookPath("git")
	if err != nil {
		if errors.Is(err, exec.ErrNotFound) {
			return nil, ErrGitNotFound
		}
		return nil, fmt.Errorf("look up git executable: %w", err)
	}
	return &Client{gitPath: git}, nil
}

func (c *Client) Discover(ctx context.Context, startDir string) (guard.Repository, error) {
	if ctx.Err() != nil {
		return guard.Repository{}, ctx.Err()
	}
	cleanDir := filepath.Clean(startDir)
	info, err := os.Stat(cleanDir)
	if err != nil {
		return guard.Repository{}, fmt.Errorf("%w: %s", ErrInvalidWorkDir, cleanDir)
	}
	if !info.IsDir() {
		return guard.Repository{}, fmt.Errorf("%w: %s is not a directory", ErrInvalidWorkDir, cleanDir)
	}
	out, err := c.Run(ctx, cleanDir, "rev-parse", "--show-toplevel")
	if err != nil {
		if ctx.Err() != nil {
			return guard.Repository{}, ctx.Err()
		}
		return guard.Repository{}, fmt.Errorf("%w: %s", ErrNotGitRepository, out)
	}
	if out == "" {
		return guard.Repository{}, ErrNotGitRepository
	}
	return guard.Repository{Root: out}, nil
}
