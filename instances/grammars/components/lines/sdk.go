package lines

import "github.com/steve-care-software/rodkina/libs/asts/domain/grammars"

// NewLine creates a new line
func NewLine() Line {
	lineBuilder := grammars.NewLineBuilder()
	containerBuilder := grammars.NewContainerBuilder()
	return createLine(
		lineBuilder,
		containerBuilder,
	)
}

// Line represents a line component
type Line interface {
	FromElements(elements []grammars.Element) grammars.Line
}
