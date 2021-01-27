package engine

import (
	"fmt"
)

// Node -
type Node interface {
	// ToString() string
}

// Name -
type Name = string

// Attrs -
type Attrs = map[string]string

// Children -
type Children struct {
	Nodes []Node
}

// GetCount -
func (children *Children) GetCount() int {
	return len(children.Nodes)
}

// Element -
type Element struct {
	Name         Name
	Attrs        Attrs
	Children     *Children
	CanSelfClose bool
}

// GetAttributes -
func (element *Element) GetAttributes() string {
	var attributes string

	for name, value := range element.Attrs {
		attributes = fmt.Sprintf("%s %s=\"%s\"", attributes, name, value)
	}

	return attributes
}

// CreateElement -
func CreateElement(name Name, attrs Attrs, children []Node) *Element {
	return &Element{
		Name:         name,
		Attrs:        attrs,
		Children:     &Children{Nodes: children},
		CanSelfClose: false,
	}
}

// SetAttribute -
func (element *Element) SetAttribute(name string, value string) {
	element.Attrs[name] = value
}
