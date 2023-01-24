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

type component struct {
	cardinality cardinalities.Cardinality
	channel     channels.Channel
	compose     composes.Compose
	suite       suites.Suite
	element     elements.Element
	line        lines.Line
	block       blocks.Block
	everything  everythings.Everything
	token       tokens.Token
}

func createComponent(
	cardinality cardinalities.Cardinality,
	channel channels.Channel,
	compose composes.Compose,
	suite suites.Suite,
	element elements.Element,
	line lines.Line,
	block blocks.Block,
	everything everythings.Everything,
	token tokens.Token,
) Component {
	out := component{
		cardinality: cardinality,
		channel:     channel,
		compose:     compose,
		suite:       suite,
		element:     element,
		line:        line,
		block:       block,
		everything:  everything,
		token:       token,
	}

	return &out
}

// Cardinality returns the cardinality
func (app *component) Cardinality() cardinalities.Cardinality {
	return app.cardinality
}

// Channel returns the channel
func (app *component) Channel() channels.Channel {
	return app.channel
}

// Compose returns the compose
func (app *component) Compose() composes.Compose {
	return app.compose
}

// Suite returns the suite
func (app *component) Suite() suites.Suite {
	return app.suite
}

// Element returns the element
func (app *component) Element() elements.Element {
	return app.element
}

// Line returns the line
func (app *component) Line() lines.Line {
	return app.line
}

// Block returns the block
func (app *component) Block() blocks.Block {
	return app.block
}

// Everything returns the everything
func (app *component) Everything() everythings.Everything {
	return app.everything
}

// Token returns the token
func (app *component) Token() tokens.Token {
	return app.token
}
