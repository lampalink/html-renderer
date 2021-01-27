package tags

import (
	"github.com/lampalink/html-renderer/engine"
)

// STYLEAttrs -
type STYLEAttrs struct{}

// STYLE -
func STYLE(attrs STYLEAttrs, children ...engine.Node) *engine.Element {
	styleElement := engine.CreateElement("style", engine.Attrs{}, children)

	return styleElement
}
