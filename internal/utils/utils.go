package utils

import (
	"fmt"
	"os"

	"github.com/anveshreddy18/tech-event-notifier/internal/notifier"
)

func GetEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func FillConfigMap(channel notifier.NotificationChannel) map[notifier.NotifierEnvVars]string {
	configMap := make(map[notifier.NotifierEnvVars]string)
	switch channel {
	case notifier.WhatsAppChannel:
		configMap[notifier.NotifierEnvVars(notifier.WhatsAppPhoneNumber)] = GetEnv("WHATSAPP_PHONE_NUMBER", "")
		configMap[notifier.NotifierEnvVars(notifier.WhatsAppAPIKey)] = GetEnv("WHATSAPP_API_KEY", "")
		configMap[notifier.NotifierEnvVars(notifier.WhatsAppPhoneNumberID)] = GetEnv("WHATSAPP_PHONE_NUMBER_ID", "")
	default:
		fmt.Printf("Unsupported notification channel: %s\n", channel)
		panic(fmt.Sprintf("Unsupported notification channel: %s", channel))
	}
	return configMap
}
