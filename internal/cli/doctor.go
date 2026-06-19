package cli

import (
	"github.com/skyvxl/reze/internal/gitx"
	"github.com/skyvxl/reze/internal/guard"
	"github.com/spf13/cobra"
)

func newDoctorCommand() *cobra.Command {
	return &cobra.Command{
		Use:          "doctor",
		Short:        "Check development environment",
		SilenceUsage: true,

		RunE: func(cmd *cobra.Command, args []string) error {
			gitClient, err := gitx.NewClient()
			if err != nil {
				return HumanError(err)
			}
			doctor := guard.NewDoctor(gitClient)
			opts := guard.DoctorOptions{
				StartDir: ".",
			}
			report, err := doctor.Run(cmd.Context(), opts)
			if err != nil {
				return HumanError(err)
			}
			err = guard.PrintReportText(cmd.OutOrStdout(), report)
			if err != nil {
				return err
			}
			return nil
		},
	}
}
