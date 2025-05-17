package notifier

import (
	"bytes"
	"fmt"
	"net/http"
)

// NewWhatsAppNotifier creates a new WhatsApp notifier with the provided configuration.
// It expects the phone number, API key, and phone number ID to be provided.
func NewWhatsAppNotifier(config map[NotifierEnvVars]string) *WhatsAppConfig {
	return &WhatsAppConfig{
		PhoneNumber:   config[NotifierEnvVars(WhatsAppPhoneNumber)],
		APIKey:        config[NotifierEnvVars(WhatsAppAPIKey)],
		PhoneNumberID: config[NotifierEnvVars(WhatsAppPhoneNumberID)],
	}
}

// Notify sends a notification via WhatsApp.
func (c *WhatsAppConfig) Notify(event string) error {
	fmt.Println("Sending WhatsApp notification...")
	fmt.Printf("Phone Number: %s, API Key: %s, Phone Number ID: %s\n",
		c.PhoneNumber, c.APIKey, c.PhoneNumberID)
	fmt.Printf("Event: %s\n", event)

	url := fmt.Sprintf("https://graph.facebook.com/v22.0/%s/messages", c.PhoneNumberID)
	payload := []byte(`{
		"messaging_product": "whatsapp",
		"to": "` + c.PhoneNumber + `",
		"type": "text",
		"text": {
			"body": "` + event + `"
		}
	}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-200 response: %s", resp.Status)
	}
	fmt.Println("WhatsApp notification sent successfully.")
	return nil
}
