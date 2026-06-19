package cli

import (
	"context"
	"errors"

	"github.com/skyvxl/reze/internal/gitx"
)

func HumanError(err error) error {
	switch {
	case err == nil:
		return nil
	case errors.Is(err, context.Canceled):
		return errors.New("operation canceled")
	case errors.Is(err, gitx.ErrGitNotFound):
		return errors.New("git not found: install Git and make sure it is available in PATH")
	case errors.Is(err, gitx.ErrNotGitRepository):
		return errors.New("current directory is not a git repository")
	default:
		return err
	}
}
