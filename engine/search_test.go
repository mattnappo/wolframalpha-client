package engine

import (
	"github.com/xoreo/wolframalpha-client/core"
	"testing"
)

func TestNewSearchObject(t *testing.T) {
	searchText := "what is 10 added to ten"
	searchObject, err := NewSearchObject(searchText)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(searchObject)
}

func TestSearch(t *testing.T) {
	core.InitSelenium()

	cwd, err := core.NewChromeWebDriver(8081)
	if err != nil {
		t.Fatal(err)
	}

	search, err := NewSearchObject("integral from 0 to 2 of x^4 dx")
	if err != nil {
		t.Fatal(err)
	}

	err = search.Search(cwd)
	if err != nil {
		t.Fatal(err)
	}

	err = cwd.Stop()
	if err != nil {
		t.Fatal(err)
	}
}

func TestParseSearch(t *testing.T) {
	search := "900 - 400x^3 + integral of 0 to 12 of 7x^2 * 19x^(12%)"
	parsed, err := parseSearch(search)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(parsed)
}
