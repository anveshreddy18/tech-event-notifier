package fetcher

// So I want to write an interface for the fetcher now.

type fetcher interface {
	Fetch() error
}

func NewFetcher(aggregator string, city string, apiKey string) fetcher {
	switch aggregator {
	case "meetup":
		return NewMeetup(city, apiKey)
	case "eventbrite":
		return NewEventbrite(city, apiKey)
	default:
		return nil
	}
}
