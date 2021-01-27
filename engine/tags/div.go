package tags

import (
	"fmt"

	"htmlrenderer/engine"
)

// DIVAttrs -
type DIVAttrs struct {
	Class string
	Style map[string]string
}

// DIV -
func DIV(attrs DIVAttrs, children ...engine.Node) *engine.Element {
	divElement := engine.CreateElement("div", engine.Attrs{}, children)

	// divElement.CanSelfClose = true

	if len(attrs.Class) > 0 {
		divElement.SetAttribute("class", attrs.Class)
	}

	style := ""
	for key, value := range attrs.Style {
		style = fmt.Sprintf("%s%s:%s;", style, key, value)
	}

	if len(style) > 0 {
		divElement.SetAttribute("style", style)
	}

	return divElement
}
