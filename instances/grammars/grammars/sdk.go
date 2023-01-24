package grammars

import (
	"github.com/steve-care-software/rodkina/instances/grammars/components"
	"github.com/steve-care-software/rodkina/instances/grammars/tokens"
	"github.com/steve-care-software/rodkina/libs/asts/domain/grammars"
	"github.com/steve-care-software/rodkina/libs/asts/domain/grammars/cardinalities"
	"github.com/steve-care-software/rodkina/libs/asts/domain/grammars/values"
)

const grammarTokenName = "grammar"

const assignmentSign = ":"
const everythingPrefixSign = "#"
const everythingEscapePrefixSign = "!"
const amountSeparator = "|"
const cardinalitySingleOptional = "?"
const cardinalityMultipleMandatory = "+"
const cardinalityMultipleOptional = "*"
const cardinalityPrefix = "["
const cardinalitySuffix = "]"
const cardinalitySeparator = ","
const lineDelimiter = "|"
const blockSuffix = ";"
const validSuiteName = "valid"
const invalidSuiteName = "invalid"
const suiteDelimiter = "&"
const suitePrefix = "---"
const suiteSuffix = ";"
const channelPrefix = "-"
const channelSuffix = ";"
const channelPrevNextPrefix = "["
const channelPrevNextSuffix = "]"
const channelPrevNextDelimiter = ":"
const rootPrefix = "@"
const rootSuffix = ";"
const instructionSuffix = ";"
const externalTokenPrefix = "{"
const externalTokenSuffix = "{"

// NewGrammar creates a new grammar instance
func NewGrammar() Grammar {
	component := components.NewComponent()
	token := tokens.NewToken()
	builder := grammars.NewBuilder()
	channelsBuilder := grammars.NewChannelsBuilder()
	channelBuilder := grammars.NewChannelBuilder()
	instanceBuilder := grammars.NewInstanceBuilder()
	everythingBuilder := grammars.NewEverythingBuilder()
	tokensBuilder := grammars.NewTokensBuilder()
	tokenBuilder := grammars.NewTokenBuilder()
	suitesBuilder := grammars.NewSuitesBuilder()
	suiteBuilder := grammars.NewSuiteBuilder()
	blockBuilder := grammars.NewBlockBuilder()
	lineBuilder := grammars.NewLineBuilder()
	containerBuilder := grammars.NewContainerBuilder()
	elementBuilder := grammars.NewElementBuilder()
	composeBuilder := grammars.NewComposeBuilder()
	composeElementBuilder := grammars.NewComposeElementBuilder()
	valueBuilder := values.NewBuilder()
	cardinalityBuilder := cardinalities.NewBuilder()
	return createGrammar(
		component,
		token,
		builder,
		channelsBuilder,
		channelBuilder,
		instanceBuilder,
		everythingBuilder,
		tokensBuilder,
		tokenBuilder,
		suitesBuilder,
		suiteBuilder,
		blockBuilder,
		lineBuilder,
		containerBuilder,
		elementBuilder,
		composeBuilder,
		composeElementBuilder,
		valueBuilder,
		cardinalityBuilder,
	)
}

// Grammar represents the grammar
type Grammar interface {
	Grammar() grammars.Grammar
	GrammarToken() grammars.Token
}
