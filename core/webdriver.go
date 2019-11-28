package core

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

// ChromeWebDriver is an abstraction class for a selenium chrome webdriver.
type ChromeWebDriver struct {
	WebDriver *selenium.WebDriver `json:"web_driver"` // The web driver itself
	Service   *selenium.Service   `json:"service"`    // The selenium service

	Options      []selenium.ServiceOption `json:"options"`      // The service configuration/options
	Capabilities selenium.Capabilities    `json:"Capabilities"` // The capabilities (further browser configuration)

	Running bool `json:"running"` // The status of the instance
	Port    int  `json:"port"`    // The port that the instance will run on
}

// NewChromeWebDriver returns a new LIVE web driver.
func NewChromeWebDriver(port int) (*ChromeWebDriver, error) {
	fmt.Printf("\n\nSeleniumPath: [%s]\nChromeDriverPath: [%s]\n\n", SeleniumPath, ChromeDriverPath)

	options := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),             // Start an X frame buffer for the browser to run in
		selenium.ChromeDriver(ChromeDriverPath), // Specify the path to the chroem driver
		selenium.Output(os.Stderr),              // Output debug information to STDERR
	}

	// Initialize the selenium service
	service, err := selenium.NewSeleniumService(SeleniumPath, port, options...)
	if err != nil {
		return nil, err
	}

	// Connect to the webdriver instance running locally.
	caps := selenium.Capabilities{"browser": "chrome"}

	// Declare the capabilities for chrome
	var chromeCaps chrome.Capabilities
	chromeCaps.Path = ChromeBinPath
	caps.AddChrome(chromeCaps)

	// Create the web driver remote itself
	webDriver, err := selenium.NewRemote(
		caps, // The capabilities
		fmt.Sprintf("http://localhost:%d/wd/hub", // The ip to listen on
			port), // The port to listen on
	)
	if err != nil {
		return nil, err
	}

	// Construct the ChromeWebDriver
	newCWD := &ChromeWebDriver{
		WebDriver: &webDriver, // The live webdriver itself
		Service:   service,    // The live selenium service

		Options:      options, // The options declared earlier
		Capabilities: caps,    // The capabilities declared earlier

		Port:    port, // The port of the instance
		Running: true, // The instance is currently running
	}
	return newCWD, nil

}

// Stop stops a ChromeWebDriver.
func (cwd *ChromeWebDriver) Stop() error {
	// Stop the service
	err := (*cwd).Service.Stop()
	if err != nil {
		return err
	}

	// Stop the webdriver
	err = *(*cwd).(*WebDriver).Quit()
	if err != nil {
		return err
	}

	cwd.Running = false
	return nil
}

// String marshals a ChromeWebDriver.
func (cwd *ChromeWebDriver) String() string {
	json, _ := json.MarshalIndent(*cwd, "", "  ")
	return string(json)
}
