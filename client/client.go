package client

import (
	"errors"
	"strings"
	"time"

	"github.com/xoreo/wolframalpha-client/common"
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

	// Parse the search
	parsedSearch, err := parseSearch(search)
	if err != nil {
		return Search{}, err
	}

	return Search{
		Search:  search,       // The search text
		Request: parsedSearch, // The parsed search url
		Time:    time.Now(),   // The current time
	}, nil
}

// parseSearchForRequest parses a search (string) and returns
// a WolframAlpha request.
func parseSearch(search string) (string, error) {
	baseRequest := "https://www.wolframalpha.com/input/?i=" // The base request
	search = strings.Replace(search, " ", "+", -1)          // Replace spaces with +

	var escapedSearch []string
	for i := 0; i < len(search); i++ { // For each letter in the search
		illegal := false
		for j, char := range common.UnsafeChars { // For each illegal char
			// Make that char legal
			if string(search[i]) == char {
				escapedSearch = append(escapedSearch, common.SafeChars[j])
				illegal = true
				break
			}
		}

		if !illegal {
			escapedSearch = append(escapedSearch, string(search[i]))
		}
	}

	// Make sure the search isn't too long
	if len(escapedSearch) >= 1900 {
		return "", errors.New("that search is too long")
	}

	// Piece the []string back together (into a string)
	escapedSearchString := strings.Join(escapedSearch, "")

	request := baseRequest + escapedSearchString // Make the request
	return request, nil
}
