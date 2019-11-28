package client

import (
	"errors"
	"time"
)

// Search represents a search on WolframAlpha.
type Search struct {
	Search  string    `json:"search"`  // The search text
	Request string    `json:"request"` // The WolframAlpha request (URL)
	Time    time.Time `json:"time"`    // The time of search
}

// NewSearch makes a new Search struct.
func NewSearch(search string) (Search, error) {
	if search == "" {
		return Search{}, errors.New("search text must not be empty")
	}

	return Search{
		Search:  search,     // The search text
		Request: "",         // Blank for now
		Time:    time.Now(), // The current time
	}, nil
}
