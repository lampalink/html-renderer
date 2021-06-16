package tags

import (
	"fmt"

	"github.com/lampalink/html-renderer/engine"
)

// TABLEAttrs -
type TABLEAttrs struct {
	Class string
	Style map[string]string
}

// TABLE -
func TABLE(attrs TABLEAttrs, children ...engine.Node) *engine.Element {
	tableElement := engine.CreateElement("table", engine.Attrs{}, children)

	if len(attrs.Class) > 0 {
		tableElement.SetAttribute("class", attrs.Class)
	}

	style := ""
	for key, value := range attrs.Style {
		style = fmt.Sprintf("%s%s:%s;", style, key, value)
	}

	if len(style) > 0 {
		tableElement.SetAttribute("style", style)
	}

	return tableElement
}
