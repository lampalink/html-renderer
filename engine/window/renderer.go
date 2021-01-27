package window

import (
	"fmt"

	"htmlrenderer/engine"
)

// Render -
func Render(node engine.Node) (string, error) {
	var err error

	output := ""
	switch child := node.(type) {
	case nil:
		break
	case int, int8, int16, int32, int64:
		output = fmt.Sprintf("%s%d", output, child)
		break
	case float32, float64:
		output = fmt.Sprintf("%s%f", output, child)
		break
	case string:
		output = fmt.Sprintf("%s%s", output, child)
		break
	case bool:
		// output = fmt.Sprintf("%s%d", output, child)
		break
	case *HTMLDocument:
		if output, err = Render(child.Render()); err != nil {
			return output, err
		}
		break
	// case *engine.Children:
	// 	var content string
	// 	if content, err = Render(child.Children); err != nil {
	// 		return output, err
	// 	}
	// 	output = fmt.Sprintf("%s<%s%s>%s</%s>", output, child.Name, child.GetAttributes(), content, child.Name)
	// 	break
	case *engine.Element:
		if child.CanSelfClose && child.Children.GetCount() < 1 {
			output = fmt.Sprintf("%s<%s%s />", output, child.Name, child.GetAttributes())
			break
		}

		var content string
		if content, err = Render(child.Children.Nodes); err != nil {
			return output, err
		}

		output = fmt.Sprintf("%s<%s%s>%s</%s>", output, child.Name, child.GetAttributes(), content, child.Name)
		break
	case []*engine.Element:
		midOutput := ""
		for _, item := range child {
			var midResult string
			if midResult, err = Render(item); err != nil {
				return output, err
			}

			midOutput = fmt.Sprintf("%s%s", midOutput, midResult)
		}

		output = fmt.Sprintf("%s%s", output, midOutput)
		break
	case []engine.Node:
		midOutput := ""
		for _, item := range child {
			var midResult string
			if midResult, err = Render(item); err != nil {
				return output, err
			}

			midOutput = fmt.Sprintf("%s%s", midOutput, midResult)
		}

		output = fmt.Sprintf("%s%s", output, midOutput)
		break
	default:
		fmt.Printf("unknown node %+v of type\n", node)
		break
	}

	return output, nil
}
