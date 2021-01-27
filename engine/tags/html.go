package tags

import (
	"htmlrenderer/engine"
)

// HTMLAttrs -
type HTMLAttrs struct{}

// HTML -
func HTML(attrs HTMLAttrs, children ...engine.Node) *engine.Element {
	htmlElement := engine.CreateElement("html", engine.Attrs{}, children)

	return htmlElement
}
