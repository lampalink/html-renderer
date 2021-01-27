package window

import (
	"htmlrenderer/engine"
	"htmlrenderer/engine/tags"
)

// HTMLDocument -
type HTMLDocument struct {
	Title      string
	StyleSheet *StyleSheet
	Body       engine.Node
}

// NewHTMLDocument -
func NewHTMLDocument() *HTMLDocument {
	return &HTMLDocument{
		StyleSheet: &StyleSheet{
			Path:  []string{},
			Style: []string{},
		},
	}
}

// Render -
func (htmlDocument *HTMLDocument) Render() *engine.Element {
	return tags.HTML(tags.HTMLAttrs{}, []engine.Node{
		tags.HEAD(tags.HEADAttrs{}, htmlDocument.StyleSheet.Render()),
		tags.BODY(tags.BODYAttrs{}, htmlDocument.Body),
	})
}
