package grammars

import (
	"github.com/steve-care-software/rodkina/libs/asts/domain/grammars"
	"github.com/steve-care-software/rodkina/libs/asts/domain/grammars/cardinalities"
	"github.com/steve-care-software/rodkina/libs/asts/domain/grammars/values"
)

const grammarTokenName = "grammar"

const byteLength = 256
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
func NewGrammar() grammars.Grammar {
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
	grammarIns := createGrammar(
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

	ins, err := grammarIns.Execute()
	if err != nil {
		panic(err)
	}

	return ins
}
