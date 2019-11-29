package engine

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/tebeka/selenium"
	"github.com/xoreo/wolframalpha-client/common"
	"github.com/xoreo/wolframalpha-client/core"
)

const (
	// calculationDivTag is the class of the div containing a singular
	// calculation from the result.
	calculationDivTag = "._3k-JE4Gq"

	// labelTag is the class of the div containing the label
	// from a result.
	labelTag = ".-ux9E2hV"

	// urlTag is the class of the div containing the url
	// from a result.
	urlTag = ".ZbCdqua6"
)

// SearchObject represents a search on WolframAlpha.
type SearchObject struct {
	SearchText string    `json:"search_text"` // The search text
	Request    string    `json:"request"`     // The WolframAlpha request (URL)
	Time       time.Time `json:"time"`        // The time of search

	Result []LatexObject `json:"result"` // The search result
}

// NewSearchObject makes a new Search struct.
func NewSearchObject(searchText string) (SearchObject, error) {
	if searchText == "" {
		return SearchObject{}, errors.New("search text must not be empty")
	}

	// Parse the search
	parsedSearch, err := parseSearch(searchText)
	if err != nil {
		return SearchObject{}, err
	}

	return SearchObject{
		SearchText: searchText,   // The search text
		Request:    parsedSearch, // The parsed search url
		Time:       time.Now(),   // The current time

		Result: nil, // Nil for now
	}, nil
}

// Search will execute the search.
func (search *SearchObject) Search(cwd *core.ChromeWebDriver) error {
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

	// Update the search result
	search.Result = latex

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
	labelDiv, err := calculation.FindElement(selenium.ByCSSSelector, labelTag)
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
	urlDiv, err := calculation.FindElement(selenium.ByCSSSelector, urlTag)
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

// parseSearchForRequest parses a search (string) and returns
// a WolframAlpha request.
func parseSearch(search string) (string, error) {
	baseRequest := "https://www.wolframalpha.com/input/?i=" // The base request

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

	// Make sure the search isn't too long
	if len(escapedSearch) >= 1900 {
		return "", errors.New("that search is too long")
	}

	// Piece the []string back together (into a string)
	escapedSearchString := strings.Join(escapedSearch, "")
	escapedSearchString = strings.Replace(escapedSearchString, " ", "+", -1)
	request := baseRequest + escapedSearchString // Make the request
	return request, nil
}
