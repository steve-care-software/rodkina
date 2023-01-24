package tokens

import (
	components_blocks "github.com/steve-care-software/rodkina/instances/grammars/components/blocks"
	components_elements "github.com/steve-care-software/rodkina/instances/grammars/components/elements"
	components_lines "github.com/steve-care-software/rodkina/instances/grammars/components/lines"
	components_suites "github.com/steve-care-software/rodkina/instances/grammars/components/suites"
	"github.com/steve-care-software/rodkina/libs/asts/domain/grammars"
)

type token struct {
	suite        components_suites.Suite
	element      components_elements.Element
	block        components_blocks.Block
	line         components_lines.Line
	tokenBuilder grammars.TokenBuilder
}

func createToken(
	suite components_suites.Suite,
	element components_elements.Element,
	block components_blocks.Block,
	line components_lines.Line,
	tokenBuilder grammars.TokenBuilder,
) Token {
	out := token{
		suite:        suite,
		element:      element,
		block:        block,
		line:         line,
		tokenBuilder: tokenBuilder,
	}

	return &out
}

// AllCharacters returns the all character token
func (app *token) AllCharacters(tokenName string, values string) grammars.Token {
	elementsList := []grammars.Element{}
	for _, oneValue := range values {
		element := app.element.FromValue(byte(oneValue))
		elementsList = append(elementsList, element)
	}

	return app.allElementsToken(tokenName, elementsList, app.suite.Suites(map[string]bool{
		values: true,
	}))
}

// AnyCharacter returns the any character token
func (app *token) AnyCharacter(tokenName string, values string) grammars.Token {
	suitesData := map[string]bool{}
	elementsList := []grammars.Element{}
	for _, oneValue := range values {
		element := app.element.FromValue(byte(oneValue))
		elementsList = append(elementsList, element)
		suitesData[string(oneValue)] = true
	}

	return app.AnyElement(
		tokenName,
		elementsList,
		app.suite.Suites(suitesData),
	)
}

// AnyElement returns an any element token
func (app *token) AnyElement(tokenName string, elementsList []grammars.Element, suites grammars.Suites) grammars.Token {
	linesList := []grammars.Line{}
	for _, oneElement := range elementsList {
		linesList = append(
			linesList,
			app.line.FromElements([]grammars.Element{
				oneElement,
			}),
		)
	}

	return app.FromBlock(
		tokenName,
		app.block.FromLines(linesList),
		suites,
	)
}

// FromBlock returns a token from block
func (app *token) FromBlock(name string, block grammars.Block, suites grammars.Suites) grammars.Token {
	builder := app.tokenBuilder.Create().
		WithName(name).
		WithBlock(block)

	if suites != nil {
		builder.WithSuites(suites)
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}

func (app *token) allCharacterToken(tokenName string, values string) grammars.Token {
	elementsList := []grammars.Element{}
	for _, oneValue := range values {
		element := app.element.FromValue(byte(oneValue))
		elementsList = append(elementsList, element)
	}

	return app.allElementsToken(tokenName, elementsList, app.suite.Suites(map[string]bool{
		values: true,
	}))
}

func (app *token) allElementsToken(tokenName string, elementsList []grammars.Element, suites grammars.Suites) grammars.Token {
	return app.FromBlock(
		tokenName,
		app.block.FromLines([]grammars.Line{
			app.line.FromElements(elementsList),
		}),
		suites,
	)
}
