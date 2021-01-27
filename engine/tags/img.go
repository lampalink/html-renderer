package tags

import (
	"fmt"

	"github.com/lampalink/html-renderer/engine"
)

// IMGAttrs -
type IMGAttrs struct {
	Src   string
	Class string
	Style map[string]string
}

// IMG -
func IMG(attrs IMGAttrs, children ...engine.Node) *engine.Element {
	imgElement := engine.CreateElement("img", engine.Attrs{}, children)

	imgElement.CanSelfClose = true

	if len(attrs.Src) > 0 {
		imgElement.SetAttribute("src", attrs.Src)
	}

	if len(attrs.Class) > 0 {
		imgElement.SetAttribute("class", attrs.Class)
	}

	style := ""
	for key, value := range attrs.Style {
		style = fmt.Sprintf("%s%s:%s;", style, key, value)
	}
	if len(style) > 0 {
		imgElement.SetAttribute("style", style)
	}

	return imgElement
}
