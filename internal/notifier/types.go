package notifier

// NotifierEnvVars represents the environment variables used for notifier configuration.
type NotifierEnvVars string

// NotificationChannel represents the type of notification channel.
type NotificationChannel string

// The available notification channels.
const (
	WhatsAppChannel NotificationChannel = "whatsapp"
	TelegramChannel NotificationChannel = "telegram"
	SMSChannel      NotificationChannel = "sms"
	EmailChannel    NotificationChannel = "email"
	SlackChannel    NotificationChannel = "slack"
	DiscordChannel  NotificationChannel = "discord"
)

// WhatsAppEnvVars represents the environment variables for WhatsApp configuration.
type WhatsAppEnvVars NotifierEnvVars

// WhatsApp environment variables.
const (
	WhatsAppPhoneNumber   WhatsAppEnvVars = "WHATSAPP_PHONE_NUMBER"
	WhatsAppAPIKey        WhatsAppEnvVars = "WHATSAPP_API_KEY"
	WhatsAppPhoneNumberID WhatsAppEnvVars = "WHATSAPP_PHONE_NUMBER_ID"
)

// WhatsAppConfig holds the configuration for the WhatsApp notifier.
type WhatsAppConfig struct {
	PhoneNumber   string
	APIKey        string
	PhoneNumberID string
}

type notifier interface {
	Notify(event string) error
}
