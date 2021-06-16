package tags

import (
	"fmt"

	"github.com/lampalink/html-renderer/engine"
)

// TDAttrs -
type TDAttrs struct {
	Class   string
	Style   map[string]string
	RowSpan int
	ColSpan int
}

// TD -
func TD(attrs TDAttrs, children ...engine.Node) *engine.Element {
	tdElement := engine.CreateElement("td", engine.Attrs{}, children)

	if len(attrs.Class) > 0 {
		tdElement.SetAttribute("class", attrs.Class)
	}

	if attrs.RowSpan > 0 {
		tdElement.SetAttribute("rowspan", fmt.Sprintf("%d", attrs.RowSpan))
	}

	if attrs.ColSpan > 0 {
		tdElement.SetAttribute("colspan", fmt.Sprintf("%d", attrs.ColSpan))
	}

	style := ""
	for key, value := range attrs.Style {
		style = fmt.Sprintf("%s%s:%s;", style, key, value)
	}

	if len(style) > 0 {
		tdElement.SetAttribute("style", style)
	}

	return tdElement
}
