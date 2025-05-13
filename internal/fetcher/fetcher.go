package fetcher

import "fmt"

// So I want to write an interface for the fetcher now.

type fetcher interface {
	Fetch() error
}

func NewFetcher(aggregator string, city string) fetcher {
	switch aggregator {
	case "meetup":
		fmt.Println("Switch case meetup")
		return NewMeetup(city)
	case "eventbrite":
		fmt.Println("Switch case eventbrite")
		return NewEventbrite(city)
	default:
		fmt.Println("Switch case default")
		return nil
	}
}
