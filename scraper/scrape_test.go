package scraper

import (
	"testing"

	"github.com/xoreo/wolframalpha-client/core"
)

func TestScrape(t *testing.T) {
	core.InitSelenium()

	cwd, err := core.NewChromeWebDriver(8081)
	if err != nil {
		t.Fatal(err)
	}

	err = Scrape(cwd.WebDriver)
	if err != nil {
		t.Fatal(err)
	}
}
