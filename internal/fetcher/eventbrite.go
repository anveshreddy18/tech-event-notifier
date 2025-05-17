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

func (c *EventbriteConfig) Fetch() (string, error) {
	fmt.Println("Fetching events from Eventbrite...")
	resp, err := helper(c.EventbriteAPIKey)
	if err != nil {
		return "", err
	}
	return resp, nil
}

func helper(apiKey string) (string, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://www.eventbriteapi.com/v3/categories/", nil)
	req.Header.Add("Authorization", "Bearer "+apiKey)
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("errored when sending request to the server: %w", err)

	}

	defer resp.Body.Close()
	resp_body, _ := io.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("EventBrite Fetch : received non-200 response from: %s", resp.Status)
	}
	if len(resp_body) == 0 {
		return "", fmt.Errorf("received empty response from Eventbrite API")
	}
	// Find what to send as response, sending the json format is breaking the whatsapp API for some reason
	// For now, just returning a placeholder string
	resp_body = []byte("Test Data")
	return string(resp_body), nil
}
