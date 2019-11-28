package core

import "testing"

func TestNewChromeWebDriver(t *testing.T) {
	InitSelenium()

	cwd, err := NewChromeWebDriver(8081)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(cwd.String())
}
