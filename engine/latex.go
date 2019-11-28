package engine

import "errors"

// LatexObject represents a LaTeX object.
type LatexObject struct {
	Label string `json:"label"` // What the object represents
	URL   string `json:"url"`   // The url to the object
}

// NewLatexObject creates a new LatexObject.
func NewLatexObject(label, url string) (LatexObject, error) {
	if label == "" || url == "" {
		return LatexObject{},
			errors.New("arguments to construct LatexObject must not be nil")
	}

	return LatexObject{
		Label: label, // The label
		URL:   url,   // The URL to the LaTeX
	}, nil
}
