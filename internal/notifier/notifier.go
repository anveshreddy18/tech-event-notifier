package notifier

func NewNotifier(notifierType NotificationChannel, config map[NotifierEnvVars]string) notifier {
	switch notifierType {
	case WhatsAppChannel:
		return NewWhatsAppNotifier(config)
	// case TelegramChannel:
	// 	return NewTelegramNotifier(config[NotifierEnvVars("botToken")], config[NotifierEnvVars("chatID")])
	// case SMSChannel:
	// 	return NewSMSNotifier(config[NotifierEnvVars("phoneNumber")], config[NotifierEnvVars("apiKey")])
	// case EmailChannel:
	// 	return NewEmailNotifier(config)
	// case SlackChannel:
	// 	return NewSlackNotifier(config)
	// case DiscordChannel:
	// 	return NewDiscordNotifier(config)
	default:
		return nil
	}
}
