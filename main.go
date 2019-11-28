package main

import (
	"github.com/xoreo/wolframalpha-client/core"
	"github.com/xoreo/wolframalpha-client/scraper"
)

func main() {
	core.InitSelenium()

	// Construct a new core.ChromeWebDriver
	cwd, err := core.NewChromeWebDriver(8081)
	if err != nil {
		panic(err)
	}

	// Call the scrape function to scrape the webapp
	err = scraper.Scrape(cwd.WebDriver)
	if err != nil {
		panic(err)
	}
}
