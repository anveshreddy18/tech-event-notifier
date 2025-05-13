package cmd

import (
	"fmt"

	"github.com/anveshreddy18/tech-event-notifier/internal/fetcher"
	"github.com/spf13/cobra"
)

type FetchOptions struct {
	City string
	// Date     string // TODO: @anveshreddy18 -- for future
	// Distance string // TODO: @anveshreddy18 -- for future
}

var fetchOptions FetchOptions

var fetchCmd = &cobra.Command{
	Use:   "fetch [aggregator] [flags]",
	Short: "Fetch events from the specified aggregator",
	Long:  "Fetch events from the specified aggregator. The aggregator can be any event aggregator platform that is supported by this application.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		aggregator := args[0]
		fmt.Printf("Fetching events from aggregator %s in city %s\n", aggregator, fetchOptions.City)
		eventFetcher := fetcher.NewFetcher(aggregator, fetchOptions.City)
		if err := eventFetcher.Fetch(); err != nil {
			fmt.Printf("Error fetching events: %v\n", err)
			return
		}
		fmt.Println("Events fetched successfully.")
	},
}

func init() {
	// add the fetch command to the root command
	rootCmd.AddCommand(fetchCmd)

	// add any flags required for the fetch command here
	fetchCmd.Flags().StringVarP(&fetchOptions.City, "city", "c", "", "City to fetch events from")
	fetchCmd.MarkFlagRequired("city")
}
