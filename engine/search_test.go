package engine

import (
	"testing"

	"github.com/xoreo/wolframalpha-client/core"
)

func TestMakeSearch(t *testing.T) {
	core.InitSelenium()

	cwd, err := core.NewChromeWebDriver(8081)
	if err != nil {
		t.Fatal(err)
	}

	err = MakeSearch(cwd)
	if err != nil {
		t.Fatal(err)
	}
}
