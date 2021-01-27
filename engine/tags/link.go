package tags

import (
	"github.com/lampalink/html-renderer/engine"
)

// LINKAttrs -
type LINKAttrs struct {
	Rel  string
	Href string
}

// LINK -
func LINK(attrs LINKAttrs, children ...engine.Node) *engine.Element {
	linkElement := engine.CreateElement("link", engine.Attrs{}, children)

	linkElement.CanSelfClose = true

	if len(attrs.Rel) > 0 {
		linkElement.SetAttribute("rel", attrs.Rel)
	}

	if len(attrs.Href) > 0 {
		linkElement.SetAttribute("href", attrs.Href)
	}

	return linkElement
}
