package tags

import (
	"fmt"

	"github.com/lampalink/html-renderer/engine"
)

// TRAttrs -
type TRAttrs struct {
	Class string
	Style map[string]string
}

// TR -
func TR(attrs TRAttrs, children ...engine.Node) *engine.Element {
	trElement := engine.CreateElement("tr", engine.Attrs{}, children)

	if len(attrs.Class) > 0 {
		trElement.SetAttribute("class", attrs.Class)
	}

	style := ""
	for key, value := range attrs.Style {
		style = fmt.Sprintf("%s%s:%s;", style, key, value)
	}

	if len(style) > 0 {
		trElement.SetAttribute("style", style)
	}

	return trElement
}
