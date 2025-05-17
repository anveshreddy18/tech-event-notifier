package main

import (
	"fmt"
	"os"

	"github.com/anveshreddy18/tech-event-notifier/cmd"
	"github.com/anveshreddy18/tech-event-notifier/internal/fetcher"
	"github.com/anveshreddy18/tech-event-notifier/internal/notifier"
	"github.com/anveshreddy18/tech-event-notifier/internal/utils"
)

const (
	DefaultAggregator          = fetcher.EventbriteAggregator
	DefaultNotificationChannel = notifier.WhatsAppChannel
	DefaultCity                = "Bangalore"
)

func main() {

	if len(os.Args) > 1 {
		cmd.Execute()
	} else {

		// Start of the fetching process

		// we default to meetup aggregator
		aggregator := DefaultAggregator
		city := DefaultCity
		// fetch the api key from the environment variable
		apiKey := utils.GetEnv("EVENTBRITE_API_KEY", "")
		if apiKey == "" {
			panic("EVENTBRITE_API_KEY environment variable is not set")
		}

		fetcher := fetcher.NewFetcher(aggregator, city, apiKey)
		msg, err := fetcher.Fetch()
		if err != nil {
			panic(fmt.Sprintf("Error fetching events: %v", err))
		}
		fmt.Println("Events fetched successfully.")

		// Start of the notification process

		configMap := utils.FillConfigMap(DefaultNotificationChannel)

		notifier := notifier.NewNotifier(DefaultNotificationChannel, configMap)
		if notifier == nil {
			panic("Notifier is not supported or could not be created")
		}
		if err := notifier.Notify(msg); err != nil {
			panic(fmt.Sprintf("Error sending notification: %v", err))
		}
		fmt.Println("Notification sent successfully.")

		// End of the notification process
		fmt.Println("All operations completed successfully.")
		fmt.Println("You can also run the command with 'ten fetch' to fetch events from the aggregator and 'ten notify' to send notifications.")
	}
}
