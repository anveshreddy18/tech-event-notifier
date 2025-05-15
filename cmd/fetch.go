package cmd

import (
	"fmt"

	"github.com/anveshreddy18/tech-event-notifier/internal/fetcher"
	"github.com/anveshreddy18/tech-event-notifier/internal/utils"
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
	RunE: func(cmd *cobra.Command, args []string) error {
		aggregator := args[0]
		fmt.Printf("Fetching events from aggregator %s in city %s\n", aggregator, fetchOptions.City)
		apiKey := utils.GetEnv("EVENT_AGGREGATOR_API_KEY", "")
		if apiKey == "" {
			fmt.Println("EVENT_AGGREGATOR_API_KEY environment variable is not set")
			return fmt.Errorf("missing EVENT_AGGREGATOR_API_KEY")
		}
		// create a new fetcher based on the aggregator
		eventFetcher := fetcher.NewFetcher(aggregator, fetchOptions.City, apiKey)
		if eventFetcher == nil {
			fmt.Printf("Unsupported aggregator: %s\n", aggregator)
			return fmt.Errorf("unsupported aggregator: %s", aggregator)
		}
		// Fetch events using the fetcher
		if err := eventFetcher.Fetch(); err != nil {
			fmt.Printf("Error fetching events: %v\n", err)
			return err
		}
		fmt.Println("Events fetched successfully.")
		return nil
	},
}

func init() {
	// add the fetch command to the root command
	rootCmd.AddCommand(fetchCmd)

	// add any flags required for the fetch command here
	fetchCmd.Flags().StringVarP(&fetchOptions.City, "city", "c", "", "City to fetch events from")
	fetchCmd.MarkFlagRequired("city")
}
