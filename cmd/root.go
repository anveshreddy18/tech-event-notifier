package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// define the variables to store the global flags for this root commands.
	version string
	verbose bool
)

var rootCmd = &cobra.Command{
	Use:   "ten [commands] [flags]",
	Short: "ten stands for tech-event-notifier command",
	Long:  "ten is the root command, it has subcommands to perform query, and notify. On a higher level, these runs the golang application to fetch events from the passed in event aggregaor site and city, it then notifies in the given communication channel",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func init() {
	// add any global flags required for the rootCmd here
	rootCmd.PersistentFlags().StringVarP(&version, "version", "v", "0.0.1", "version of the application")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "V", false, "enable verbose output")
}
