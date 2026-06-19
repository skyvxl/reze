package cli

import (
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
				return HumanError(err)
			}
			fmt.Fprintln(cmd.OutOrStdout(), "git: ok")
			return nil
		},
	}
}
