package fetcher

// EventAggregator represents the type of event aggregator to use.
type EventAggregator string

const (
	MeetupAggregator     EventAggregator = "meetup"
	EventbriteAggregator EventAggregator = "eventbrite"
)

type fetcher interface {
	Fetch() (string, error)
}

func NewFetcher(aggregator EventAggregator, city string, apiKey string) fetcher {
	switch aggregator {
	case MeetupAggregator:
		return NewMeetup(city, apiKey)
	case EventbriteAggregator:
		return NewEventbrite(city, apiKey)
	default:
		return nil
	}
}
