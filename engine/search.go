package engine

import (
	"fmt"
	"io/ioutil"

	"github.com/xoreo/wolframalpha-client/client"
	"github.com/xoreo/wolframalpha-client/core"

	"github.com/tebeka/selenium"
)

const (
	// outputDivTag the class of the div containing the output from
	// WolframAlpha.
	outputDivTag = "._1jPLqSjg"
	// outputDivTag = "._2ckn5Li6"

	// calculationDivTag is the class of the div containing a singular
	// calculation from the result.
	calculationDivTag = "._3k-JE4Gq"
)

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
	outputDiv, err := driver.FindElement(
		selenium.ByCSSSelector, outputDivTag)
	if err != nil {
		return err
	}

	// Get all of the calculation divs
	calculations, err := outputDiv.FindElements(
		selenium.ByCSSSelector, calculationDivTag,
	)
	fmt.Printf("\n\nSIZE: %d\n\n", len(calculations))
	if err != nil {
		return err
	}

	fmt.Printf("calculations:\n\n\n\n\n")
	fmt.Println(calculations)

	screenshot, err := driver.Screenshot()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("screenshot.png", screenshot, 0644)
	if err != nil {
		return err
	}

	// Collect all of the latex
	var latex []LatexObject
	for _, calculation := range calculations {
		fmt.Printf("\n\nI AM RUNNING\n\n")
		// Extract the label of the calculation
		labelDiv, err := calculation.FindElement(selenium.ByCSSSelector, ".-ux9E2hV")
		if err != nil {
			return err
		}
		s, err = labelDiv.TagName()
		if err != nil {
			return err
		}
		fmt.Printf("\n\n\n%s THING", s)

		// Get the text from the div
		label, err := labelDiv.Text()
		if err != nil {
			return err
		}

		// Find the div containing the url
		urlDiv, err := calculation.FindElement(selenium.ByCSSSelector, ".ZbCdqua6")
		if err != nil {
			return err
		}

		// Get the URL itself
		url, err := urlDiv.GetAttribute("src")
		if err != nil {
			return err
		}

		// Create the new LatexObject
		newLatex, err := NewLatexObject(label, url)
		if err != nil {
			return err
		}

		latex = append(latex, newLatex)
	}
	fmt.Printf("\n\n\n\n")
	fmt.Println(latex)

	// // Get the output text in that div
	// outputString, err := outputDiv.Text()
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("\n\n\n\n%s\n\n\n\n", outputString)

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
