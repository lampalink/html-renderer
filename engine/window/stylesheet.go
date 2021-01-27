package window

import (
	"github.com/lampalink/html-renderer/engine"
	"github.com/lampalink/html-renderer/engine/tags"
)

// StyleSheet -
type StyleSheet struct {
	Path  []string
	Style []string
}

// AddPath -
func (styleSheet *StyleSheet) AddPath(cssPath string) {
	styleSheet.Path = append(styleSheet.Path, cssPath)
}

// AddStyle -
func (styleSheet *StyleSheet) AddStyle(style string) {
	styleSheet.Style = append(styleSheet.Style, style)
}

// Render -
func (styleSheet *StyleSheet) Render() []*engine.Element {
	styles := []*engine.Element{}

	if len(styleSheet.Path) > 0 {
		for _, href := range styleSheet.Path {
			styles = append(styles, tags.LINK(tags.LINKAttrs{
				Rel:  "stylesheet",
				Href: href,
			}))
		}
	}

	if len(styleSheet.Style) > 0 {
		for _, style := range styleSheet.Style {
			styles = append(styles, tags.STYLE(tags.STYLEAttrs{}, style))
		}
	}

	return styles
}
