package cli

import (
	"fmt"
	"time"

	"github.com/skyvxl/reze/internal/buildinfo"
	"github.com/spf13/cobra"
)

func newVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Run: func(cmd *cobra.Command, args []string) {
			out := cmd.OutOrStdout()
			if t, err := time.Parse(time.RFC3339, buildinfo.BuildDate); err == nil {
				fmt.Fprintf(out, "%s-%s-%s\n", buildinfo.Version, buildinfo.Commit, t.Format("2006-01-02"))
				return
			}
			fmt.Fprintf(out, "%s-%s-%s\n", buildinfo.Version, buildinfo.Commit, buildinfo.BuildDate)
		},
	}
}
