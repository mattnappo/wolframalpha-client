package engine

import (
	"fmt"

	"github.com/xoreo/wolframalpha-client/client"
	"github.com/xoreo/wolframalpha-client/core"

	"github.com/tebeka/selenium"
)

// Search will make a search on WolframAlpha and return the output.
func Search(search client.Search, cwd *core.ChromeWebDriver) error {
	driver := *cwd.WebDriver

	// Make the get request to the url of the search
	err := driver.Get(search.Request)
	if err != nil {
		return err
	}

	// Find the output code based on css class
	output, err := driver.FindElement(selenium.ByCSSSelector, "._2ZqIJD6E")
	if err != nil {
		return err
	}

	fmt.Println(output.Text())

	return nil
}
