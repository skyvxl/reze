package cli

import (
	"errors"
	"fmt"

	"github.com/skyvxl/reze/internal/gitx"
	"github.com/spf13/cobra"
)

func newDoctorCommand() *cobra.Command {
	return &cobra.Command{
		Use:          "doctor",
		Short:        "Check development environment",
		SilenceUsage: true,

		RunE: func(cmd *cobra.Command, args []string) error {
			if err := gitx.CheckGit(cmd.Context(), "."); err != nil {
				switch {
				case errors.Is(err, gitx.ErrGitNotFound):
					return fmt.Errorf("git not found: install Git and make sure it is available in PATH")

				case errors.Is(err, gitx.ErrNotGitRepository):
					return fmt.Errorf("current directory is not a git repository")

				default:
					return err
				}
			}
			fmt.Fprintln(cmd.OutOrStdout(), "git: ok")
			return nil
		},
	}
}
