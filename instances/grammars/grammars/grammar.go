package grammars

import (
	"github.com/steve-care-software/rodkina/instances/grammars/components"
	"github.com/steve-care-software/rodkina/instances/grammars/tokens"
	"github.com/steve-care-software/rodkina/libs/asts/domain/grammars"
	"github.com/steve-care-software/rodkina/libs/asts/domain/grammars/cardinalities"
	"github.com/steve-care-software/rodkina/libs/asts/domain/grammars/values"
)

type grammar struct {
	component             components.Component
	token                 tokens.Token
	builder               grammars.Builder
	channelsBuilder       grammars.ChannelsBuilder
	channelBuilder        grammars.ChannelBuilder
	instanceBuilder       grammars.InstanceBuilder
	everythingBuilder     grammars.EverythingBuilder
	tokensBuilder         grammars.TokensBuilder
	tokenBuilder          grammars.TokenBuilder
	suitesBuilder         grammars.SuitesBuilder
	suiteBuilder          grammars.SuiteBuilder
	blockBuilder          grammars.BlockBuilder
	lineBuilder           grammars.LineBuilder
	containerBuilder      grammars.ContainerBuilder
	elementBuilder        grammars.ElementBuilder
	composeBuilder        grammars.ComposeBuilder
	composeElementBuilder grammars.ComposeElementBuilder
	valueBuilder          values.Builder
	cardinalityBuilder    cardinalities.Builder
}

func createGrammar(
	component components.Component,
	token tokens.Token,
	builder grammars.Builder,
	channelsBuilder grammars.ChannelsBuilder,
	channelBuilder grammars.ChannelBuilder,
	instanceBuilder grammars.InstanceBuilder,
	everythingBuilder grammars.EverythingBuilder,
	tokensBuilder grammars.TokensBuilder,
	tokenBuilder grammars.TokenBuilder,
	suitesBuilder grammars.SuitesBuilder,
	suiteBuilder grammars.SuiteBuilder,
	blockBuilder grammars.BlockBuilder,
	lineBuilder grammars.LineBuilder,
	containerBuilder grammars.ContainerBuilder,
	elementBuilder grammars.ElementBuilder,
	composeBuilder grammars.ComposeBuilder,
	composeElementBuilder grammars.ComposeElementBuilder,
	valueBuilder values.Builder,
	cardinalityBuilder cardinalities.Builder,
) Grammar {
	out := grammar{
		component:             component,
		token:                 token,
		builder:               builder,
		channelsBuilder:       channelsBuilder,
		channelBuilder:        channelBuilder,
		instanceBuilder:       instanceBuilder,
		everythingBuilder:     everythingBuilder,
		tokensBuilder:         tokensBuilder,
		tokenBuilder:          tokenBuilder,
		suitesBuilder:         suitesBuilder,
		suiteBuilder:          suiteBuilder,
		blockBuilder:          blockBuilder,
		lineBuilder:           lineBuilder,
		containerBuilder:      containerBuilder,
		elementBuilder:        elementBuilder,
		composeBuilder:        composeBuilder,
		composeElementBuilder: composeElementBuilder,
		valueBuilder:          valueBuilder,
		cardinalityBuilder:    cardinalityBuilder,
	}

	return &out
}

// Grammar returns the grammar
func (app *grammar) Grammar() grammars.Grammar {
	root := app.GrammarToken()
	channels := app.component.Channel().Channels()
	ins, err := app.builder.Create().
		WithRoot(root).
		WithChannels(channels).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

// GrammarToken returns the grammar token
func (app *grammar) GrammarToken() grammars.Token {
	return app.component.Token().FromBlock(
		"grammar",
		app.component.Block().FromLines([]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(app.rootToken(), app.component.Cardinality().Once()),
				app.component.Element().FromToken(app.channelToken(), app.component.Cardinality().Cardinality(0, nil)),
				app.component.Element().FromToken(app.instructionToken(), app.component.Cardinality().Cardinality(1, nil)),
			}),
		}),
		app.component.Suite().Suites(map[string]bool{
			`
				@myValue;
				myValue: 45;
			`: true,
			`
				@myValue;
				-myChannel;
				-mySecondChannel [prev];

				myValue: myCompose
					---
					valid: myValidCompose;
					invalid : myInvalidCompose
							& mySecondInvalidCompose
							;
				;
			`: true,
			`
				@myValue;
				-myChannel;
				-mySecondChannel [:next];

				myEverything: #myToken
					---
					valid: myValidCompose;
					invalid : myInvalidCompose
							& mySecondInvalidCompose
							;
				;
			`: true,
			`
				@myValue;
				-myChannel;
				-mySecondChannel [prev:next];

				myVariable: myToken*
				---
					valid	: firstValidCompose
							& secondValidCompose
							;
				;
			`: true,
		}),
	)
}

func (app *grammar) instructionToken() grammars.Token {
	return app.component.Token().FromBlock(
		"instruction",
		app.component.Block().FromLines([]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(app.valueAssignmentToken(), app.component.Cardinality().Once()),
			}),
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(app.composeAssignmentToken(), app.component.Cardinality().Once()),
			}),
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(app.everythingAssignmentToken(), app.component.Cardinality().Once()),
			}),
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(app.tokenAssignmentToken(), app.component.Cardinality().Once()),
			}),
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(app.externalTokenAssignmentToken(), app.component.Cardinality().Once()),
			}),
		}),
		app.component.Suite().Suites(map[string]bool{
			`myValue: 45;`: true,
			`
				myCompose: myCompose
					---
					valid: myValidCompose;
					invalid : myInvalidCompose
							& mySecondInvalidCompose
							;
				;
			`: true,
			`
				myEverything: #myToken
					---
					valid: myValidCompose;
					invalid : myInvalidCompose
							& mySecondInvalidCompose
							;
				;
			`: true,
			`
				myVariable: myToken*
				---
					valid	: firstValidCompose
							& secondValidCompose
							;
				;
			`: true,
			`
				myExternalToken: cf113f0af255e83f32351a3c32c05fc824e46119f93fb00bfece497421cd4e790b0d682a7bb54d3136c87fdd9222f2ed6a36c904958b0a797b98a22d9d94601c
				---
					valid	: firstValidCompose
							& secondValidCompose
							;
				;
			`: true,
		}),
	)
}

func (app *grammar) externalTokenAssignmentToken() grammars.Token {
	return app.component.Token().FromBlock(
		"externalTokenAssignment",
		app.component.Block().FromLines([]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(app.token.VariableName(), app.component.Cardinality().Once()),
				app.component.Element().FromValue([]byte(assignmentSign)[0]),
				app.component.Element().FromToken(app.token.Sha512Hex(), app.component.Cardinality().Once()),
				app.component.Element().FromToken(app.suiteToken(), app.component.Cardinality().Once()),
				app.component.Element().FromValue([]byte(blockSuffix)[0]),
			}),
		}),
		app.component.Suite().Suites(map[string]bool{
			`
				myExternalToken: df39913d8d243a34734495461d1c9261dc8efbb9be7afe52d2930bbb3de4b8dcd8153744e13991a6bbcd35894fe69d520f047b2039282cbca193ee0389f25ec5
				---
					valid	: firstValidCompose
							& secondValidCompose
							;
				;
			`: true,
		}),
	)
}

func (app *grammar) rootToken() grammars.Token {
	return app.component.Token().FromBlock(
		"root",
		app.component.Block().FromLines([]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromValue([]byte(rootPrefix)[0]),
				app.component.Element().FromToken(app.token.VariableName(), app.component.Cardinality().Once()),
				app.component.Element().FromValue([]byte(rootSuffix)[0]),
			}),
		}),
		app.component.Suite().Suites(map[string]bool{
			`@myRoot;`: true,
		}),
	)
}

func (app *grammar) channelToken() grammars.Token {
	pMax := uint(1)
	return app.component.Token().FromBlock(
		"channel",
		app.component.Block().FromLines([]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromValue([]byte(channelPrefix)[0]),
				app.component.Element().FromToken(app.token.VariableName(), app.component.Cardinality().Once()),
				app.component.Element().FromToken(app.channelPreviousNextToken(), app.component.Cardinality().Cardinality(0, &pMax)),
				app.component.Element().FromValue([]byte(channelSuffix)[0]),
			}),
		}),
		app.component.Suite().Suites(map[string]bool{
			`-myChanToken;`:             true,
			`-myChanToken [prev];`:      true,
			`-myChanToken [:next];`:     true,
			`-myChanToken [prev:next];`: true,
		}),
	)
}

func (app *grammar) channelPreviousNextToken() grammars.Token {
	return app.component.Token().FromBlock(
		"channelPreviousNext",
		app.component.Block().FromLines([]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromValue([]byte(channelPrevNextPrefix)[0]),
				app.component.Element().FromToken(app.channelPreviousNextInsideToken(), app.component.Cardinality().Once()),
				app.component.Element().FromValue([]byte(channelPrevNextSuffix)[0]),
			}),
		}),
		app.component.Suite().Suites(map[string]bool{
			`[prev]`:      true,
			`[:next]`:     true,
			`[prev:next]`: true,
		}),
	)
}

func (app *grammar) channelPreviousNextInsideToken() grammars.Token {
	return app.component.Token().FromBlock(
		"channelPreviousNextInside",
		app.component.Block().FromLines([]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(app.token.VariableName(), app.component.Cardinality().Once()),
				app.component.Element().FromValue([]byte(channelPrevNextDelimiter)[0]),
				app.component.Element().FromToken(app.token.VariableName(), app.component.Cardinality().Once()),
			}),
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromValue([]byte(channelPrevNextDelimiter)[0]),
				app.component.Element().FromToken(app.token.VariableName(), app.component.Cardinality().Once()),
			}),
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(app.token.VariableName(), app.component.Cardinality().Once()),
			}),
		}),
		app.component.Suite().Suites(map[string]bool{
			`prev`:      true,
			`:next`:     true,
			`prev:next`: true,
		}),
	)
}

func (app *grammar) tokenAssignmentToken() grammars.Token {
	return app.component.Token().FromBlock(
		"tokenAssignment",
		app.component.Block().FromLines([]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(app.token.VariableName(), app.component.Cardinality().Once()),
				app.component.Element().FromValue([]byte(assignmentSign)[0]),
				app.component.Element().FromToken(app.blockToken(), app.component.Cardinality().Once()),
				app.component.Element().FromToken(app.suiteToken(), app.component.Cardinality().Once()),
				app.component.Element().FromValue([]byte(blockSuffix)[0]),
			}),
		}),
		app.component.Suite().Suites(map[string]bool{
			`
				myVariable: myToken*
				---
					valid	: firstValidCompose
							& secondValidCompose
							;
				;
			`: true,
			`
				myVariable	: myToken*
							| mySecond+
							| myThird?
							| fourth[2] fifth[0,] sixth[1,] seventh[0,234]
				---
					valid	: firstValidCompose
							& secondValidCompose
							;

					invalid	: myComposeToken;
				;
			`: true,
		}),
	)
}

func (app *grammar) valueAssignmentToken() grammars.Token {
	return app.component.Token().FromBlock(
		"valueAssignment",
		app.component.Block().FromLines([]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(app.token.VariableName(), app.component.Cardinality().Once()),
				app.component.Element().FromValue([]byte(assignmentSign)[0]),
				app.component.Element().FromToken(app.token.AnyByte(), app.component.Cardinality().Cardinality(1, nil)),
			}),
		}),
		app.component.Suite().Suites(map[string]bool{
			`
				myValue: 45;
			`: true,
		}),
	)
}

func (app *grammar) suiteToken() grammars.Token {
	pMax := uint(1)
	return app.component.Token().FromBlock(
		"suite",
		app.component.Block().FromLines([]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(app.component.Token().AllCharacters("suitePrefixConst", suitePrefix), app.component.Cardinality().Once()),
				app.component.Element().FromToken(app.suiteValidToken(), app.component.Cardinality().Once()),
				app.component.Element().FromToken(app.suiteInvalidToken(), app.component.Cardinality().Cardinality(0, &pMax)),
			}),
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(app.component.Token().AllCharacters("suitePrefixConst", suitePrefix), app.component.Cardinality().Once()),
				app.component.Element().FromToken(app.suiteValidToken(), app.component.Cardinality().Cardinality(0, &pMax)),
				app.component.Element().FromToken(app.suiteInvalidToken(), app.component.Cardinality().Once()),
			}),
		}),
		app.component.Suite().Suites(map[string]bool{
			`
				---
			`: false,
			`
				---
				valid	: firstValidCompose
						& secondValidCompose
						;
			`: true,
			`
				---
				invalid : myComposeToken
					  	& mySecondComposeToken
						& myThirdComposeToken
					  	;
			`: true,
			`
				---
				valid	: firstValidCompose
						& secondValidCompose
						;

				invalid	: myComposeToken;
			`: true,
			`
				---
				valid	: firstValidCompose;

				invalid : myComposeToken
					  	& mySecondComposeToken
						& myThirdComposeToken
					  	;
			`: true,
		}),
	)
}

func (app *grammar) suiteInvalidToken() grammars.Token {
	return app.component.Token().FromBlock(
		"suiteInvalid",
		app.component.Block().FromLines([]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(app.component.Token().AllCharacters("invalidConst", invalidSuiteName), app.component.Cardinality().Once()),
				app.component.Element().FromToken(app.suiteBlockToken(), app.component.Cardinality().Once()),
			}),
		}),
		app.component.Suite().Suites(map[string]bool{
			`invalid: myComposeToken;`: true,
			`
				invalid : myComposeToken
					  	& mySecondComposeToken
						& myThirdComposeToken
					  	;
			`: true,
		}),
	)
}

func (app *grammar) suiteValidToken() grammars.Token {
	return app.component.Token().FromBlock(
		"suiteValid",
		app.component.Block().FromLines([]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(app.component.Token().AllCharacters("validConst", validSuiteName), app.component.Cardinality().Once()),
				app.component.Element().FromToken(app.suiteBlockToken(), app.component.Cardinality().Once()),
			}),
		}),
		app.component.Suite().Suites(map[string]bool{
			`valid: myComposeToken;`: true,
			`
				valid : myComposeToken
					  & mySecondComposeToken
					  ;
			`: true,
		}),
	)
}

func (app *grammar) suiteBlockToken() grammars.Token {
	return app.component.Token().FromBlock(
		"suiteBlock",
		app.component.Block().FromLines([]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromValue([]byte(assignmentSign)[0]),
				app.component.Element().FromToken(app.token.VariableName(), app.component.Cardinality().Once()),
				app.component.Element().FromToken(app.delimiterThenSuiteElementToken(), app.component.Cardinality().Cardinality(0, nil)),
				app.component.Element().FromValue([]byte(suiteSuffix)[0]),
			}),
		}),
		app.component.Suite().Suites(map[string]bool{
			`: myComposeToken;`:                        true,
			`: myComposeToken & mySecondComposeToken;`: true,
		}),
	)
}

func (app *grammar) delimiterThenSuiteElementToken() grammars.Token {
	return app.component.Token().FromBlock(
		"delimiterThenSuiteElement",
		app.component.Block().FromLines([]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromValue([]byte(suiteDelimiter)[0]),
				app.component.Element().FromToken(app.token.VariableName(), app.component.Cardinality().Once()),
			}),
		}),
		app.component.Suite().Suites(map[string]bool{
			`& myComposeToken`: true,
		}),
	)
}

func (app *grammar) blockToken() grammars.Token {
	return app.component.Token().FromBlock(
		"block",
		app.component.Block().FromLines([]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(app.lineToken(), app.component.Cardinality().Once()),
				app.component.Element().FromToken(app.delimiterThenLineToken(), app.component.Cardinality().Cardinality(0, nil)),
			}),
		}),
		app.component.Suite().Suites(map[string]bool{
			`myToken*`: true,
			`
				myToken*
				| mySecond+
				| myThird?
				| fourth[2] fifth[0,] sixth[1,] seventh[0,234]
			`: true,
		}),
	)
}

func (app *grammar) delimiterThenLineToken() grammars.Token {
	return app.component.Token().FromBlock(
		"delimiterThenLine",
		app.component.Block().FromLines([]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromValue([]byte(lineDelimiter)[0]),
				app.component.Element().FromToken(app.lineToken(), app.component.Cardinality().Once()),
			}),
		}),
		app.component.Suite().Suites(map[string]bool{
			`|myToken* mySecond+ myThird? fourth[2] fifth[0,] sixth[1,] seventh[0,234]`: true,
		}),
	)
}

func (app *grammar) lineToken() grammars.Token {
	return app.component.Token().FromBlock(
		"line",
		app.component.Block().FromLines([]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(app.elementToken(), app.component.Cardinality().Cardinality(1, nil)),
			}),
		}),
		app.component.Suite().Suites(map[string]bool{
			`myToken* mySecond+ myThird? fourth[2] fifth[0,] sixth[1,] seventh[0,234]`: true,
		}),
	)
}

func (app *grammar) elementToken() grammars.Token {
	return app.component.Token().FromBlock(
		"element",
		app.component.Block().FromLines([]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(app.token.VariableName(), app.component.Cardinality().Once()),
				app.component.Element().FromToken(app.cardinalityToken(), app.component.Cardinality().Cardinality(0, nil)),
			}),
		}),
		app.component.Suite().Suites(map[string]bool{
			`myToken*`:       true,
			`myToken+`:       true,
			`myToken?`:       true,
			`myToken[2]`:     true,
			`myToken[234]`:   true,
			`myToken[0,]`:    true,
			`myToken[1,]`:    true,
			`myToken[0,234]`: true,
		}),
	)
}

func (app *grammar) cardinalityToken() grammars.Token {
	return app.component.Token().FromBlock(
		"cardinality",
		app.component.Block().FromLines([]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromValue([]byte(cardinalitySingleOptional)[0]),
			}),
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromValue([]byte(cardinalityMultipleMandatory)[0]),
			}),
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromValue([]byte(cardinalityMultipleOptional)[0]),
			}),
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromValue([]byte(cardinalityPrefix)[0]),
				app.component.Element().FromToken(app.token.AnyNumber(), app.component.Cardinality().Cardinality(1, nil)),
				app.component.Element().FromValue([]byte(cardinalitySuffix)[0]),
			}),
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromValue([]byte(cardinalityPrefix)[0]),
				app.component.Element().FromToken(app.token.AnyNumber(), app.component.Cardinality().Cardinality(1, nil)),
				app.component.Element().FromValue([]byte(cardinalitySeparator)[0]),
				app.component.Element().FromToken(app.token.AnyNumber(), app.component.Cardinality().Cardinality(0, nil)),
				app.component.Element().FromValue([]byte(cardinalitySuffix)[0]),
			}),
		}),
		app.component.Suite().Suites(map[string]bool{
			`*`:       true,
			`+`:       true,
			`?`:       true,
			`[2]`:     true,
			`[234]`:   true,
			`[0,]`:    true,
			`[1,]`:    true,
			`[0,234]`: true,
		}),
	)
}

func (app *grammar) composeAssignmentToken() grammars.Token {
	return app.component.Token().FromBlock(
		"composeAssignment",
		app.component.Block().FromLines([]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(app.token.VariableName(), app.component.Cardinality().Once()),
				app.component.Element().FromValue([]byte(assignmentSign)[0]),
				app.component.Element().FromToken(app.composeToken(), app.component.Cardinality().Once()),
				app.component.Element().FromToken(app.suiteToken(), app.component.Cardinality().Once()),
				app.component.Element().FromValue([]byte(blockSuffix)[0]),
			}),
		}),
		app.component.Suite().Suites(map[string]bool{
			`
				myCompose: myCompose
					---
					valid: myValidCompose;
					invalid : myInvalidCompose
							& mySecondInvalidCompose
							;
				;
			`: true,
			`
				myCompose: myCompose|45
					---
					valid: myValidCompose;
				;
			`: true,
		}),
	)
}

func (app *grammar) composeToken() grammars.Token {
	return app.component.Token().FromBlock(
		"compose",
		app.component.Block().FromLines([]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(app.token.VariableName(), app.component.Cardinality().Once()),
				app.component.Element().FromToken(app.separatorAmountOfComposeToken(), app.component.Cardinality().Cardinality(0, nil)),
			}),
		}),
		app.component.Suite().Suites(map[string]bool{
			`myCompose`:     true,
			`myCompose|4`:   true,
			`myCompose|234`: true,
		}),
	)
}

func (app *grammar) separatorAmountOfComposeToken() grammars.Token {
	return app.component.Token().FromBlock(
		"composeWithAmount",
		app.component.Block().FromLines([]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromValue([]byte(amountSeparator)[0]),
				app.component.Element().FromToken(app.token.AnyNumber(), app.component.Cardinality().Cardinality(1, nil)),
			}),
		}),
		app.component.Suite().Suites(map[string]bool{
			`|4`:   true,
			`|234`: true,
		}),
	)
}

func (app *grammar) everythingAssignmentToken() grammars.Token {
	return app.component.Token().FromBlock(
		"everythingAssignment",
		app.component.Block().FromLines([]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(app.token.VariableName(), app.component.Cardinality().Once()),
				app.component.Element().FromValue([]byte(assignmentSign)[0]),
				app.component.Element().FromToken(app.everythingToken(), app.component.Cardinality().Once()),
				app.component.Element().FromToken(app.suiteToken(), app.component.Cardinality().Once()),
				app.component.Element().FromValue([]byte(blockSuffix)[0]),
			}),
		}),
		app.component.Suite().Suites(map[string]bool{
			`
				myEverything: #myToken
					---
					valid: myValidCompose;
					invalid : myInvalidCompose
							& mySecondInvalidCompose
							;
				;
			`: true,
		}),
	)
}

func (app *grammar) everythingToken() grammars.Token {
	return app.component.Token().FromBlock(
		"everything",
		app.component.Block().FromLines([]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(app.everythingWithEscapeToken(), app.component.Cardinality().Once()),
			}),
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(app.everythingWithoutEscapeToken(), app.component.Cardinality().Once()),
			}),
		}),
		app.component.Suite().Suites(map[string]bool{
			"#myToken":          true,
			"#myToken!myEscape": true,
		}),
	)
}

func (app *grammar) everythingWithEscapeToken() grammars.Token {
	return app.component.Token().FromBlock(
		"everythingWithEscape",
		app.component.Block().FromLines([]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromValue([]byte(everythingPrefixSign)[0]),
				app.component.Element().FromToken(app.token.VariableName(), app.component.Cardinality().Once()),
				app.component.Element().FromValue([]byte(everythingEscapePrefixSign)[0]),
				app.component.Element().FromToken(app.token.VariableName(), app.component.Cardinality().Once()),
			}),
		}),
		app.component.Suite().Suites(map[string]bool{
			"#myToken!myEscape": true,
		}),
	)
}

func (app *grammar) everythingWithoutEscapeToken() grammars.Token {
	return app.component.Token().FromBlock(
		"everythingWithoutEscape",
		app.component.Block().FromLines([]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromValue([]byte(everythingPrefixSign)[0]),
				app.component.Element().FromToken(app.token.VariableName(), app.component.Cardinality().Once()),
			}),
		}),
		app.component.Suite().Suites(map[string]bool{
			"#myToken": true,
		}),
	)
}
