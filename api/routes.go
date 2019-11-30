package api

import "path"

// setupRoutes sets up all necessary API routes.
func (api *API) setupRoutes() error {
	api.initGETs()
	api.initPOSTs()
	return nil
}

// initGETs initializes the API's GET requests.
func (api *API) initGETs() {
	root := path.Join(api.Root)

	api.Router.GET(path.Join(root, "getHistory"), api.getHistory)
}

// initPOSTs initializes the API's POST requests.
func (api *API) initPOSTs() {
	root := path.Join(api.Root)

	api.Router.POST(path.Join(root, "search"), api.search)
}
