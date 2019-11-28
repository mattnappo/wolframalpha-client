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
	outputDiv, err := driver.FindElement(selenium.ByCSSSelector, "._2HkkNXzH")
	if err != nil {
		return err
	}

	outputString, err := outputDiv.Text()
	if err != nil {
		return err
	}
	fmt.Printf("\n\n\n\n%s\n\n\n\n", outputString)

	return nil
}
