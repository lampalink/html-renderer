package tags

import (
	"github.com/lampalink/html-renderer/engine"
)

// HEADAttrs -
type HEADAttrs struct{}

// HEAD -
func HEAD(attrs HEADAttrs, children ...engine.Node) *engine.Element {
	headElement := engine.CreateElement("head", engine.Attrs{}, children)

	return headElement
}
