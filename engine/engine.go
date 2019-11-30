package engine

// Engine is a nice abstraction for the scraping engine.
// Not sure if this needs to be implemented yet.
type Engine interface {
	// NewSearchObject makes a new Search struct.
	NewSearchObject(searchText string) (SearchObject, error)
}
