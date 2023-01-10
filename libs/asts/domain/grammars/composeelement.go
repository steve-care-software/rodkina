package grammars

import (
	"github.com/steve-care-software/rodkina/libs/asts/domain/grammars/values"
)

type composeElement struct {
	value      values.Value
	occurences uint
}

func createComposeElement(
	value values.Value,
	occurences uint,
) ComposeElement {
	out := composeElement{
		value:      value,
		occurences: occurences,
	}

	return &out
}

// Value returns the value
func (obj *composeElement) Value() values.Value {
	return obj.value
}

// Occurences returns the occurences
func (obj *composeElement) Occurences() uint {
	return obj.occurences
}
