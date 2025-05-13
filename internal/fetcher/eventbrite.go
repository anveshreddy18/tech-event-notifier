package fetcher

import (
	"fmt"
	"time"
)

type EventbriteConfig struct {
	City             string
	Date             string
	Distance         int64
	EventbriteAPIKey string
}

func NewEventbrite(city string) *EventbriteConfig {

	loc, _ := time.LoadLocation("Asia/Kolkata")
	today := time.Now().In(loc).Format("2006-01-02")
	return &EventbriteConfig{
		City:             city,
		Date:             today,
		Distance:         10,
		EventbriteAPIKey: "",
	}
}

func (c *EventbriteConfig) Fetch() error {
	fmt.Println("Fetching events from Eventbrite...")
	fmt.Printf("City: %s, Date: %s, Distance: %d\n", c.City, c.Date, c.Distance)
	fmt.Println("Fetch operation completed.")
	return nil
}
