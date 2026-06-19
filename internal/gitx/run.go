package gitx

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
)

func (c *Client) Run(ctx context.Context, dir string, args ...string) (string, error) {
	cmd := exec.CommandContext(ctx, c.gitPath, args...) // #nosec G204 -- git path comes from exec.LookPath and args are internal git subcommands
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		return strings.TrimSpace(string(out)), fmt.Errorf("git %s: %w", args[0], err)
	}
	return strings.TrimSpace(string(out)), nil
}
