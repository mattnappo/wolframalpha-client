package engine

import (
	"testing"

	"github.com/xoreo/wolframalpha-client/core"
	"github.com/xoreo/wolframalpha-client/client"
)

func TestMSearch(t *testing.T) {
	core.InitSelenium()

	cwd, err := core.NewChromeWebDriver(8081)
	if err != nil {
		t.Fatal(err)
	}

	search, err := client.NewSearch("ten added to 10")
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
