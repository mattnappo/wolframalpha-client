package client

import "testing"

func TestNewSearch(t *testing.T) {
	searchText := "what is 10 added to ten"
	search, err := NewSearch(searchText)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(search)
}

func TestParseSearch(t *testing.T) {
	search := "900 - 400x^3 + integral of 0 to 12 of 7x^2 * 19x^(12%)"
	parsed, err := parseSearch(search)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(parsed)
}
