package fetcher

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type EventbriteConfig struct {
	City             string
	Date             string
	Distance         int64
	EventbriteAPIKey string
}

func NewEventbrite(city string, apiKey string) *EventbriteConfig {
	return &EventbriteConfig{
		City:             city,
		Date:             time.Now().Format("2006-01-02"),
		Distance:         10,
		EventbriteAPIKey: apiKey,
	}
}

func (c *EventbriteConfig) Fetch() error {
	fmt.Println("Fetching events from Eventbrite...")
	// resp, err := http.Get("https://www.eventbrite.com/api/v3/events/search/?location.address=" + c.City + "&token=" + c.EventbriteAPIKey)
	// if err != nil {
	// 	return fmt.Errorf("failed to fetch events from Eventbrite: %w", err)
	// }
	// defer resp.Body.Close()
	// if resp.StatusCode != http.StatusOK {
	// 	return fmt.Errorf("failed to fetch events from Eventbrite, status code: %d", resp.StatusCode)
	// }
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return fmt.Errorf("failed to read response body: %w", err)
	// }
	// fmt.Println(string(body))
	helper(c.EventbriteAPIKey)
	return nil
}

func helper(apiKey string) {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", "https://www.eventbriteapi.com/v3/categories/", nil)

	req.Header.Add("Authorization", "Bearer "+apiKey)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	resp_body, _ := io.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
}
