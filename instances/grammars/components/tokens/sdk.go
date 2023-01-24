package tokens

import (
	components_blocks "github.com/steve-care-software/rodkina/instances/grammars/components/blocks"
	components_elements "github.com/steve-care-software/rodkina/instances/grammars/components/elements"
	components_lines "github.com/steve-care-software/rodkina/instances/grammars/components/lines"
	components_suites "github.com/steve-care-software/rodkina/instances/grammars/components/suites"
	"github.com/steve-care-software/rodkina/libs/asts/domain/grammars"
)

const byteLength = 256

// NewToken creates a new token instance
func NewToken() Token {
	suite := components_suites.NewSuite()
	element := components_elements.NewElement()
	block := components_blocks.NewBlock()
	line := components_lines.NewLine()
	tokenBuilder := grammars.NewTokenBuilder()
	return createToken(
		suite,
		element,
		block,
		line,
		tokenBuilder,
	)
}

// Token represents the token component
type Token interface {
	AllCharacters(tokenName string, values string) grammars.Token
	AnyCharacter(tokenName string, values string) grammars.Token
	AnyElement(tokenName string, elementsList []grammars.Element, suites grammars.Suites) grammars.Token
	FromBlock(name string, block grammars.Block, suites grammars.Suites) grammars.Token
}
