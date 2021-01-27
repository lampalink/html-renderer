package tags

import (
	"github.com/lampalink/html-renderer/engine"
)

// HTMLAttrs -
type HTMLAttrs struct{}

// HTML -
func HTML(attrs HTMLAttrs, children ...engine.Node) *engine.Element {
	htmlElement := engine.CreateElement("html", engine.Attrs{}, children)

	return htmlElement
}
