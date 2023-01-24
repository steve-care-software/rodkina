package elements

import (
	components_cardinalities "github.com/steve-care-software/rodkina/instances/grammars/components/cardinalities"
	"github.com/steve-care-software/rodkina/libs/asts/domain/grammars"
	"github.com/steve-care-software/rodkina/libs/asts/domain/grammars/cardinalities"
	"github.com/steve-care-software/rodkina/libs/asts/domain/grammars/values"
)

// NewElement creates a new element instance
func NewElement() Element {
	cardinality := components_cardinalities.NewCardinality()
	instanceBuilder := grammars.NewInstanceBuilder()
	elementBuilder := grammars.NewElementBuilder()
	valueBuilder := values.NewBuilder()
	return createElement(
		cardinality,
		instanceBuilder,
		elementBuilder,
		valueBuilder,
	)
}

// Element represents the element
type Element interface {
	FromEverything(everything grammars.Everything) grammars.Element
	FromToken(token grammars.Token, cardinality cardinalities.Cardinality) grammars.Element
	FromValue(value byte) grammars.Element
}
