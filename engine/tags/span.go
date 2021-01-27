package tags

import (
	"fmt"

	"github.com/lampalink/html-renderer/engine"
)

// SPANAttrs -
type SPANAttrs struct {
	Class string
	Style map[string]string
}

// SPAN -
func SPAN(attrs SPANAttrs, children ...engine.Node) *engine.Element {
	spanElement := engine.CreateElement("span", engine.Attrs{}, children)

	if len(attrs.Class) > 0 {
		spanElement.SetAttribute("class", attrs.Class)
	}

	style := ""
	for key, value := range attrs.Style {
		style = fmt.Sprintf("%s%s:%s;", style, key, value)
	}
	if len(style) > 0 {
		spanElement.SetAttribute("style", style)
	}

	return spanElement
}
