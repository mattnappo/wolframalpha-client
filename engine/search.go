package engine

import (
	"fmt"
	"io/ioutil"

	"github.com/xoreo/wolframalpha-client/client"
	"github.com/xoreo/wolframalpha-client/core"

	"github.com/tebeka/selenium"
)

// calculationDivTag is the class of the div containing a singular
// calculation from the result.
const calculationDivTag = "._3k-JE4Gq"

// Search will make a search on WolframAlpha and return the output.
func Search(search client.Search, cwd *core.ChromeWebDriver) error {
	driver := *cwd.WebDriver

	// Make the get request to the url of the search
	err := driver.Get(search.Request)
	if err != nil {
		return err
	}

	// Wait for the calculations to complete
	err = driver.Wait(waitForOutput)
	if err != nil {
		return err
	}

	// Find all calculation divs
	calculations, err := driver.FindElements(
		selenium.ByCSSSelector, calculationDivTag,
	)
	if err != nil {
		return err
	}

	// Take a screenshot of the current window
	screenshot, err := driver.Screenshot()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("screenshot.png", screenshot, 0644)
	if err != nil {
		return err
	}

	// Collect all of the latex from the calculations
	var latex []LatexObject
	for _, calculation := range calculations {
		// Get the label and the url
		label, err := extractLabel(calculation)
		if err != nil {
			return err
		}
		url, err := extractURL(calculation)
		if err != nil {
			return err
		}

		// Create the new LatexObject with the data
		newLatex, err := NewLatexObject(label, url)
		if err != nil {
			return err
		}

		latex = append(latex, newLatex) // Add it to the list of LaTeX objects
	}
	fmt.Printf("\nlatex: %v\n", latex)

	return nil
}

// waitForOutput waits for the output/result to populate the DOM.
func waitForOutput(wd selenium.WebDriver) (bool, error) {
	// I know that this is a terrible solution
	for {
		var err error
		// Find the output code based on css class
		_, err = wd.FindElement(selenium.ByCSSSelector, calculationDivTag)
		if err == nil {
			return true, nil
		}
	}
}

// extractLabel extracts the label from a calculation (response div).
func extractLabel(calculation selenium.WebElement) (string, error) {
	// Extract the label of the calculation
	labelDiv, err := calculation.FindElement(selenium.ByCSSSelector, ".-ux9E2hV")
	if err != nil {
		return "", err
	}

	// Get the text from the div
	label, err := labelDiv.Text()
	if err != nil {
		return "", err
	}

	return label, nil
}

// extractURL extracts the URL from a calculation (response div).
func extractURL(calculation selenium.WebElement) (string, error) {
	// Find the div containing the url
	urlDiv, err := calculation.FindElement(selenium.ByCSSSelector, ".ZbCdqua6")
	if err != nil {
		return "", err
	}

	// Get the URL itself
	url, err := urlDiv.GetAttribute("src")
	if err != nil {
		return "", err
	}

	return url, nil
}
