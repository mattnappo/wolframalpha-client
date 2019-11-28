package engine

import (
	"fmt"

	"github.com/xoreo/wolframalpha-client/client"
	"github.com/xoreo/wolframalpha-client/core"

	"github.com/tebeka/selenium"
)

// outputDivTag the class of the div containing the output from
// WolframAlpha.
const outputDivTag = "._2yjzGRtP"

// Search will make a search on WolframAlpha and return the output.
func Search(search client.Search, cwd *core.ChromeWebDriver) error {
	driver := *cwd.WebDriver

	// Make the get request to the url of the search
	err := driver.Get(search.Request)
	if err != nil {
		return err
	}

	// Wait for the output/result to populate the DOM
	err = driver.Wait(waitForOutput)
	if err != nil {
		return err
	}

	// Find the output code based on css class
	outputDiv, err := driver.FindElement(selenium.ByCSSSelector, outputDivTag)
	if err != nil {
		return err
	}

	// Get the output text in that div
	outputString, err := outputDiv.Text()
	if err != nil {
		return err
	}
	fmt.Printf("\n\n\n\n%s\n\n\n\n", outputString)

	return nil
}

// waitForOutput waits for the output/result to populate the DOM.
func waitForOutput(wd selenium.WebDriver) (bool, error) {
	// I know that this is a terrible solution
	for {
		var err error
		// Find the output code based on css class
		_, err = wd.FindElement(selenium.ByCSSSelector, outputDivTag)
		if err == nil {
			return true, nil
		}
	}
}
