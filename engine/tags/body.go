package tags

import (
	"github.com/lampalink/html-renderer/engine"
)

// BODYAttrs -
type BODYAttrs struct{}

// BODY -
func BODY(attrs BODYAttrs, children ...engine.Node) *engine.Element {
	bodyElement := engine.CreateElement("body", engine.Attrs{}, children)

	return bodyElement
}
