package cli

import "github.com/spf13/cobra"

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:           "reze",
		Short:         "Validate local Git repository configuration for a selected profile",
		SilenceErrors: true,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}
	rootCmd.AddCommand(newVersionCommand())
	rootCmd.AddCommand(newDoctorCommand())
	return rootCmd
}
