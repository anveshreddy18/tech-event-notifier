package main

import (
	"fmt"
	"os"

	"github.com/anveshreddy18/tech-event-notifier/cmd"
	"github.com/anveshreddy18/tech-event-notifier/internal/fetcher"
	"github.com/anveshreddy18/tech-event-notifier/internal/utils"
)

func main() {

	if len(os.Args) > 1 {
		cmd.Execute()
	} else {
		// we default to meetup aggregator
		aggregator := "meetup"
		city := "Bangalore"
		// fetch the api key from the environment variable
		apiKey := utils.GetEnv("MEETUP_API_KEY", "")
		if apiKey == "" {
			panic("MEETUP_API_KEY environment variable is not set")
		}

		fetcher := fetcher.NewFetcher(aggregator, city, apiKey)
		if err := fetcher.Fetch(); err != nil {
			panic(fmt.Sprintf("Error fetching events: %v", err))
		}
		fmt.Println("Events fetched successfully.")
	}
}
