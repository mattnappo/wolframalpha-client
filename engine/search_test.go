package engine

import (
	"testing"

	"github.com/xoreo/wolframalpha-client/client"
	"github.com/xoreo/wolframalpha-client/core"
)

func TestSearch(t *testing.T) {
	core.InitSelenium()

	cwd, err := core.NewChromeWebDriver(8081)
	if err != nil {
		t.Fatal(err)
	}

	search, err := client.NewSearch("integral from 0 to 2 of x^4 dx")
	if err != nil {
		t.Fatal(err)
	}

	err = Search(search, cwd)
	if err != nil {
		t.Fatal(err)
	}

	err = cwd.Stop()
	if err != nil {
		t.Fatal(err)
	}
}
