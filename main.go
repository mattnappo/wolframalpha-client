package main

import (
	"flag"

	"github.com/xoreo/wolframalpha-client/api"
	"github.com/xoreo/wolframalpha-client/common"
)

// StartAPIFlag determines if the API server should be started.
var StartAPIFlag = flag.Bool("start-api", false, "Start the API server")

func main() {
	flag.Parse()

	// If the flag is true
	if *StartAPIFlag {
		// Start the API server
		api, err := api.NewAPI()
		if err != nil {
			panic(err)
		}

		err = api.StartServing(common.APIServerPort)
		if err != nil {
			panic(err)
		}
	}

}
