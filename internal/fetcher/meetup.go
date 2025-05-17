package fetcher

import (
	"fmt"
	"time"
)

type MeetupConfig struct {
	City         string
	Date         string
	Distance     int64
	MeetupAPIKey string
}

func NewMeetup(city string, apiKey string) *MeetupConfig {
	return &MeetupConfig{
		City:         city,
		Date:         time.Now().Format("2006-01-02"),
		Distance:     10,
		MeetupAPIKey: apiKey,
	}
}

func (c *MeetupConfig) Fetch() (string, error) {
	fmt.Println("Fetching events from Meetup...")
	fmt.Printf("City: %s, Date: %s, Distance: %d, API Key: %s\n", c.City, c.Date, c.Distance, c.MeetupAPIKey)
	fmt.Println("Fetch operation completed.")
	return "Fetch Data", nil
}
