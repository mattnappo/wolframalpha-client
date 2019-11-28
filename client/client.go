package client

import (
	"errors"
	"fmt"
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

	return Search{
		Search:  search,     // The search text
		Request: "",         // Blank for now
		Time:    time.Now(), // The current time
	}, nil
}

// parseSearchForRequest parses a search (string) and returns
// a WolframAlpha request.
func parseSearch(search string) (string, error) {
	baseRequest := "https://www.wolframalpha.com/input/?i=" // The base request
	search = strings.ReplaceAll(search, " ", "+")           // Replace spaces with +

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

	fmt.Println(escapedSearch)

	if len(escapedSearch) >= 1900 {
		return "", errors.New("that search is too long")
	}

	return "", nil
}
