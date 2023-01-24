package components

import (
	"github.com/steve-care-software/rodkina/instances/grammars/components/blocks"
	"github.com/steve-care-software/rodkina/instances/grammars/components/cardinalities"
	"github.com/steve-care-software/rodkina/instances/grammars/components/channels"
	"github.com/steve-care-software/rodkina/instances/grammars/components/composes"
	"github.com/steve-care-software/rodkina/instances/grammars/components/elements"
	"github.com/steve-care-software/rodkina/instances/grammars/components/everythings"
	"github.com/steve-care-software/rodkina/instances/grammars/components/lines"
	"github.com/steve-care-software/rodkina/instances/grammars/components/suites"
	"github.com/steve-care-software/rodkina/instances/grammars/components/tokens"
)

// NewComponent creates a new component
func NewComponent() Component {
	cardinality := cardinalities.NewCardinality()
	channel := channels.NewChannel()
	compose := composes.NewCompose()
	suite := suites.NewSuite()
	element := elements.NewElement()
	line := lines.NewLine()
	block := blocks.NewBlock()
	everything := everythings.NewEverything()
	token := tokens.NewToken()
	return createComponent(
		cardinality,
		channel,
		compose,
		suite,
		element,
		line,
		block,
		everything,
		token,
	)
}

// Component represents the component
type Component interface {
	Cardinality() cardinalities.Cardinality
	Channel() channels.Channel
	Compose() composes.Compose
	Suite() suites.Suite
	Element() elements.Element
	Line() lines.Line
	Block() blocks.Block
	Everything() everythings.Everything
	Token() tokens.Token
}
