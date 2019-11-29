package engine

import "testing"

func TestNewLatexObject(t *testing.T) {
	latexObject, err := NewLatexObject("label", "url")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(latexObject.String())
}