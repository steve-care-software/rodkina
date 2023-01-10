package grammars

import (
	"fmt"
	"strconv"

	"github.com/steve-care-software/rodkina/libs/asts/domain/grammars"
	"github.com/steve-care-software/rodkina/libs/asts/domain/grammars/cardinalities"
	"github.com/steve-care-software/rodkina/libs/asts/domain/grammars/values"
)

type grammar struct {
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
) *grammar {
	out := grammar{
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

// Execute executes a grammar
func (app *grammar) Execute() (grammars.Grammar, error) {
	root := app.grammarToken()
	channels := app.channels()
	return app.builder.Create().
		WithRoot(root).
		WithChannels(channels).
		Now()
}

func (app *grammar) grammarToken() grammars.Token {
	return app.tokenFromBlock(
		"grammar",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.rootToken(), app.cardinalityOnce()),
				app.elementFromToken(app.channelToken(), app.cardinality(0, nil)),
				app.elementFromToken(app.instructionToken(), app.cardinality(1, nil)),
			}),
		}),
		app.suites(map[string]bool{
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
	return app.tokenFromBlock(
		"instruction",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.valueAssignmentToken(), app.cardinalityOnce()),
			}),
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.composeAssignmentToken(), app.cardinalityOnce()),
			}),
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.everythingAssignmentToken(), app.cardinalityOnce()),
			}),
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.tokenAssignmentToken(), app.cardinalityOnce()),
			}),
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.externalTokenAssignmentToken(), app.cardinalityOnce()),
			}),
		}),
		app.suites(map[string]bool{
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
	return app.tokenFromBlock(
		"externalTokenAssignment",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.variableNameToken(), app.cardinalityOnce()),
				app.elementFromValue([]byte(assignmentSign)[0]),
				app.elementFromToken(app.sha512HexToken(), app.cardinalityOnce()),
				app.elementFromToken(app.suiteToken(), app.cardinalityOnce()),
				app.elementFromValue([]byte(blockSuffix)[0]),
			}),
		}),
		app.suites(map[string]bool{
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
	return app.tokenFromBlock(
		"root",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromValue([]byte(rootPrefix)[0]),
				app.elementFromToken(app.variableNameToken(), app.cardinalityOnce()),
				app.elementFromValue([]byte(rootSuffix)[0]),
			}),
		}),
		app.suites(map[string]bool{
			`@myRoot;`: true,
		}),
	)
}

func (app *grammar) channelToken() grammars.Token {
	pMax := uint(1)
	return app.tokenFromBlock(
		"channel",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromValue([]byte(channelPrefix)[0]),
				app.elementFromToken(app.variableNameToken(), app.cardinalityOnce()),
				app.elementFromToken(app.channelPreviousNextToken(), app.cardinality(0, &pMax)),
				app.elementFromValue([]byte(channelSuffix)[0]),
			}),
		}),
		app.suites(map[string]bool{
			`-myChanToken;`:             true,
			`-myChanToken [prev];`:      true,
			`-myChanToken [:next];`:     true,
			`-myChanToken [prev:next];`: true,
		}),
	)
}

func (app *grammar) channelPreviousNextToken() grammars.Token {
	return app.tokenFromBlock(
		"channelPreviousNext",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromValue([]byte(channelPrevNextPrefix)[0]),
				app.elementFromToken(app.channelPreviousNextInsideToken(), app.cardinalityOnce()),
				app.elementFromValue([]byte(channelPrevNextSuffix)[0]),
			}),
		}),
		app.suites(map[string]bool{
			`[prev]`:      true,
			`[:next]`:     true,
			`[prev:next]`: true,
		}),
	)
}

func (app *grammar) channelPreviousNextInsideToken() grammars.Token {
	return app.tokenFromBlock(
		"channelPreviousNextInside",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.variableNameToken(), app.cardinalityOnce()),
				app.elementFromValue([]byte(channelPrevNextDelimiter)[0]),
				app.elementFromToken(app.variableNameToken(), app.cardinalityOnce()),
			}),
			app.lineFromElements([]grammars.Element{
				app.elementFromValue([]byte(channelPrevNextDelimiter)[0]),
				app.elementFromToken(app.variableNameToken(), app.cardinalityOnce()),
			}),
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.variableNameToken(), app.cardinalityOnce()),
			}),
		}),
		app.suites(map[string]bool{
			`prev`:      true,
			`:next`:     true,
			`prev:next`: true,
		}),
	)
}

func (app *grammar) tokenAssignmentToken() grammars.Token {
	return app.tokenFromBlock(
		"tokenAssignment",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.variableNameToken(), app.cardinalityOnce()),
				app.elementFromValue([]byte(assignmentSign)[0]),
				app.elementFromToken(app.blockToken(), app.cardinalityOnce()),
				app.elementFromToken(app.suiteToken(), app.cardinalityOnce()),
				app.elementFromValue([]byte(blockSuffix)[0]),
			}),
		}),
		app.suites(map[string]bool{
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
	return app.tokenFromBlock(
		"valueAssignment",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.variableNameToken(), app.cardinalityOnce()),
				app.elementFromValue([]byte(assignmentSign)[0]),
				app.elementFromToken(app.anyByteToken(), app.cardinality(1, nil)),
			}),
		}),
		app.suites(map[string]bool{
			`
				myValue: 45;
			`: true,
		}),
	)
}

func (app *grammar) suiteToken() grammars.Token {
	pMax := uint(1)
	return app.tokenFromBlock(
		"suite",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.allCharacterToken("suitePrefixConst", suitePrefix), app.cardinalityOnce()),
				app.elementFromToken(app.suiteValidToken(), app.cardinalityOnce()),
				app.elementFromToken(app.suiteInvalidToken(), app.cardinality(0, &pMax)),
			}),
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.allCharacterToken("suitePrefixConst", suitePrefix), app.cardinalityOnce()),
				app.elementFromToken(app.suiteValidToken(), app.cardinality(0, &pMax)),
				app.elementFromToken(app.suiteInvalidToken(), app.cardinalityOnce()),
			}),
		}),
		app.suites(map[string]bool{
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
	return app.tokenFromBlock(
		"suiteInvalid",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.allCharacterToken("invalidConst", invalidSuiteName), app.cardinalityOnce()),
				app.elementFromToken(app.suiteBlockToken(), app.cardinalityOnce()),
			}),
		}),
		app.suites(map[string]bool{
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
	return app.tokenFromBlock(
		"suiteValid",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.allCharacterToken("validConst", validSuiteName), app.cardinalityOnce()),
				app.elementFromToken(app.suiteBlockToken(), app.cardinalityOnce()),
			}),
		}),
		app.suites(map[string]bool{
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
	return app.tokenFromBlock(
		"suiteBlock",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromValue([]byte(assignmentSign)[0]),
				app.elementFromToken(app.variableNameToken(), app.cardinalityOnce()),
				app.elementFromToken(app.delimiterThenSuiteElementToken(), app.cardinality(0, nil)),
				app.elementFromValue([]byte(suiteSuffix)[0]),
			}),
		}),
		app.suites(map[string]bool{
			`: myComposeToken;`:                        true,
			`: myComposeToken & mySecondComposeToken;`: true,
		}),
	)
}

func (app *grammar) delimiterThenSuiteElementToken() grammars.Token {
	return app.tokenFromBlock(
		"delimiterThenSuiteElement",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromValue([]byte(suiteDelimiter)[0]),
				app.elementFromToken(app.variableNameToken(), app.cardinalityOnce()),
			}),
		}),
		app.suites(map[string]bool{
			`& myComposeToken`: true,
		}),
	)
}

func (app *grammar) blockToken() grammars.Token {
	return app.tokenFromBlock(
		"block",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.lineToken(), app.cardinalityOnce()),
				app.elementFromToken(app.delimiterThenLineToken(), app.cardinality(0, nil)),
			}),
		}),
		app.suites(map[string]bool{
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
	return app.tokenFromBlock(
		"delimiterThenLine",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromValue([]byte(lineDelimiter)[0]),
				app.elementFromToken(app.lineToken(), app.cardinalityOnce()),
			}),
		}),
		app.suites(map[string]bool{
			`|myToken* mySecond+ myThird? fourth[2] fifth[0,] sixth[1,] seventh[0,234]`: true,
		}),
	)
}

func (app *grammar) lineToken() grammars.Token {
	return app.tokenFromBlock(
		"line",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.elementToken(), app.cardinality(1, nil)),
			}),
		}),
		app.suites(map[string]bool{
			`myToken* mySecond+ myThird? fourth[2] fifth[0,] sixth[1,] seventh[0,234]`: true,
		}),
	)
}

func (app *grammar) elementToken() grammars.Token {
	return app.tokenFromBlock(
		"element",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.variableNameToken(), app.cardinalityOnce()),
				app.elementFromToken(app.cardinalityToken(), app.cardinality(0, nil)),
			}),
		}),
		app.suites(map[string]bool{
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
	return app.tokenFromBlock(
		"cardinality",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromValue([]byte(cardinalitySingleOptional)[0]),
			}),
			app.lineFromElements([]grammars.Element{
				app.elementFromValue([]byte(cardinalityMultipleMandatory)[0]),
			}),
			app.lineFromElements([]grammars.Element{
				app.elementFromValue([]byte(cardinalityMultipleOptional)[0]),
			}),
			app.lineFromElements([]grammars.Element{
				app.elementFromValue([]byte(cardinalityPrefix)[0]),
				app.elementFromToken(app.anyNumberToken(), app.cardinality(1, nil)),
				app.elementFromValue([]byte(cardinalitySuffix)[0]),
			}),
			app.lineFromElements([]grammars.Element{
				app.elementFromValue([]byte(cardinalityPrefix)[0]),
				app.elementFromToken(app.anyNumberToken(), app.cardinality(1, nil)),
				app.elementFromValue([]byte(cardinalitySeparator)[0]),
				app.elementFromToken(app.anyNumberToken(), app.cardinality(0, nil)),
				app.elementFromValue([]byte(cardinalitySuffix)[0]),
			}),
		}),
		app.suites(map[string]bool{
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
	return app.tokenFromBlock(
		"composeAssignment",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.variableNameToken(), app.cardinalityOnce()),
				app.elementFromValue([]byte(assignmentSign)[0]),
				app.elementFromToken(app.composeToken(), app.cardinalityOnce()),
				app.elementFromToken(app.suiteToken(), app.cardinalityOnce()),
				app.elementFromValue([]byte(blockSuffix)[0]),
			}),
		}),
		app.suites(map[string]bool{
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
	return app.tokenFromBlock(
		"compose",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.variableNameToken(), app.cardinalityOnce()),
				app.elementFromToken(app.separatorAmountOfComposeToken(), app.cardinality(0, nil)),
			}),
		}),
		app.suites(map[string]bool{
			`myCompose`:     true,
			`myCompose|4`:   true,
			`myCompose|234`: true,
		}),
	)
}

func (app *grammar) separatorAmountOfComposeToken() grammars.Token {
	return app.tokenFromBlock(
		"composeWithAmount",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromValue([]byte(amountSeparator)[0]),
				app.elementFromToken(app.anyNumberToken(), app.cardinality(1, nil)),
			}),
		}),
		app.suites(map[string]bool{
			`|4`:   true,
			`|234`: true,
		}),
	)
}

func (app *grammar) everythingAssignmentToken() grammars.Token {
	return app.tokenFromBlock(
		"everythingAssignment",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.variableNameToken(), app.cardinalityOnce()),
				app.elementFromValue([]byte(assignmentSign)[0]),
				app.elementFromToken(app.everythingToken(), app.cardinalityOnce()),
				app.elementFromToken(app.suiteToken(), app.cardinalityOnce()),
				app.elementFromValue([]byte(blockSuffix)[0]),
			}),
		}),
		app.suites(map[string]bool{
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
	return app.tokenFromBlock(
		"everything",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.everythingWithEscapeToken(), app.cardinalityOnce()),
			}),
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.everythingWithoutEscapeToken(), app.cardinalityOnce()),
			}),
		}),
		app.suites(map[string]bool{
			"#myToken":          true,
			"#myToken!myEscape": true,
		}),
	)
}

func (app *grammar) everythingWithEscapeToken() grammars.Token {
	return app.tokenFromBlock(
		"everythingWithEscape",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromValue([]byte(everythingPrefixSign)[0]),
				app.elementFromToken(app.variableNameToken(), app.cardinalityOnce()),
				app.elementFromValue([]byte(everythingEscapePrefixSign)[0]),
				app.elementFromToken(app.variableNameToken(), app.cardinalityOnce()),
			}),
		}),
		app.suites(map[string]bool{
			"#myToken!myEscape": true,
		}),
	)
}

func (app *grammar) everythingWithoutEscapeToken() grammars.Token {
	return app.tokenFromBlock(
		"everythingWithoutEscape",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromValue([]byte(everythingPrefixSign)[0]),
				app.elementFromToken(app.variableNameToken(), app.cardinalityOnce()),
			}),
		}),
		app.suites(map[string]bool{
			"#myToken": true,
		}),
	)
}

func (app *grammar) variableNameToken() grammars.Token {
	return app.tokenFromBlock(
		"variableName",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.lowerCaseLetters(), app.cardinalityOnce()),
				app.elementFromToken(app.anyLetterToken(), app.cardinality(0, nil)),
			}),
		}),
		app.suites(map[string]bool{
			"m":          true,
			"myVariable": true,
			"MyVariable": false,
			"0Variable":  false,
		}),
	)
}

func (app *grammar) sha512HexToken() grammars.Token {
	amount := uint(128)
	return app.tokenFromBlock(
		"sha512Hex",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.anyHexCharToken(), app.cardinality(amount, &amount)),
			}),
		}),
		app.suites(map[string]bool{
			"cf113f0af255e83f32351a3c32c05fc824e46119f93fb00bfece497421cd4e790b0d682a7bb54d3136c87fdd9222f2ed6a36c904958b0a797b98a22d9d94601c": true,
			"cf113f0af255e83f32351a3c32c05fc824e46119f93fb00bfece497421cd4e790b0d682a7bb54d3136c87fdd9222f2ed6a36c904958b0a797b98a22d9d9460":   false,
			"gf113f0af255e83f32351a3c32c05fc824e46119f93fb00bfece497421cd4e790b0d682a7bb54d3136c87fdd9222f2ed6a36c904958b0a797b98a22d9d94601c": false,
		}),
	)
}

func (app *grammar) anyHexCharToken() grammars.Token {
	return app.tokenFromBlock(
		"anyHexChar",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.aToFLowerCaseLetterToken(), app.cardinalityOnce()),
			}),
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.aToFUpperCaseLetterToken(), app.cardinalityOnce()),
			}),
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.anyNumberToken(), app.cardinalityOnce()),
			}),
		}),
		app.suites(map[string]bool{
			"a": true,
			"b": true,
			"c": true,
			"d": true,
			"e": true,
			"f": true,
			"g": false,
			"h": false,
			"i": false,
			"j": false,
			"k": false,
			"l": false,
			"m": false,
			"n": false,
			"o": false,
			"p": false,
			"q": false,
			"r": false,
			"s": false,
			"t": false,
			"u": false,
			"v": false,
			"w": false,
			"x": false,
			"y": false,
			"z": false,
			"A": true,
			"B": true,
			"C": true,
			"D": true,
			"E": true,
			"F": true,
			"G": false,
			"H": false,
			"I": false,
			"J": false,
			"K": false,
			"L": false,
			"M": false,
			"N": false,
			"O": false,
			"P": false,
			"Q": false,
			"R": false,
			"S": false,
			"T": false,
			"U": false,
			"V": false,
			"W": false,
			"X": false,
			"Y": false,
			"Z": false,
			"0": true,
			"1": true,
			"2": true,
			"3": true,
			"4": true,
			"5": true,
			"6": true,
			"7": true,
			"8": true,
			"9": true,
		}),
	)
}

func (app *grammar) anyLetterToken() grammars.Token {
	return app.tokenFromBlock(
		"anyLetter",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.upperCaseLetterToken(), app.cardinalityOnce()),
			}),
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.lowerCaseLetters(), app.cardinalityOnce()),
			}),
		}),
		app.suites(map[string]bool{
			"a": true,
			"b": true,
			"c": true,
			"d": true,
			"e": true,
			"f": true,
			"g": true,
			"h": true,
			"i": true,
			"j": true,
			"k": true,
			"l": true,
			"m": true,
			"n": true,
			"o": true,
			"p": true,
			"q": true,
			"r": true,
			"s": true,
			"t": true,
			"u": true,
			"v": true,
			"w": true,
			"x": true,
			"y": true,
			"z": true,
			"A": true,
			"B": true,
			"C": true,
			"D": true,
			"E": true,
			"F": true,
			"G": true,
			"H": true,
			"I": true,
			"J": true,
			"K": true,
			"L": true,
			"M": true,
			"N": true,
			"O": true,
			"P": true,
			"Q": true,
			"R": true,
			"S": true,
			"T": true,
			"U": true,
			"V": true,
			"W": true,
			"X": true,
			"Y": true,
			"Z": true,
			"0": false,
		}),
	)
}

func (app *grammar) anyByteToken() grammars.Token {
	validData := map[string]bool{}
	tokenName := "anyByte"
	elementsList := []grammars.Element{}
	for i := 0; i < byteLength; i++ {
		element := app.elementFromValue(byte(i))
		elementsList = append(elementsList, element)
		validData[fmt.Sprintf("%d", i)] = true
	}

	return app.anyElementToken(tokenName, elementsList, app.suites(validData))
}

func (app *grammar) anyNumberToken() grammars.Token {
	characters := "0123456789"
	return app.anyCharacterToken("anyNumber", characters)
}

func (app *grammar) aToFUpperCaseLetterToken() grammars.Token {
	characters := "ABCDEF"
	return app.anyCharacterToken("aToFUpperCaseLetter", characters)
}

func (app *grammar) aToFLowerCaseLetterToken() grammars.Token {
	characters := "abcdef"
	return app.anyCharacterToken("aToFLowerCaseLetter", characters)
}

func (app *grammar) upperCaseLetterToken() grammars.Token {
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return app.anyCharacterToken("uppercaseLetter", characters)
}

func (app *grammar) lowerCaseLetters() grammars.Token {
	characters := "abcdefghijklmnopqrstuvwxyz"
	return app.anyCharacterToken("lowerCaseLetter", characters)
}

func (app *grammar) allCharacterToken(tokenName string, values string) grammars.Token {
	elementsList := []grammars.Element{}
	for _, oneValue := range values {
		element := app.elementFromValue(byte(oneValue))
		elementsList = append(elementsList, element)
	}

	return app.allElementsToken(tokenName, elementsList, app.suites(map[string]bool{
		values: true,
	}))
}

func (app *grammar) allElementsToken(tokenName string, elementsList []grammars.Element, suites grammars.Suites) grammars.Token {
	return app.tokenFromBlock(
		tokenName,
		app.blockFromlines([]grammars.Line{
			app.lineFromElements(elementsList),
		}),
		suites,
	)
}

func (app *grammar) anyCharacterToken(tokenName string, values string) grammars.Token {
	suitesData := map[string]bool{}
	elementsList := []grammars.Element{}
	for _, oneValue := range values {
		element := app.elementFromValue(byte(oneValue))
		elementsList = append(elementsList, element)
		suitesData[string(oneValue)] = true
	}

	return app.anyElementToken(
		tokenName,
		elementsList,
		app.suites(suitesData),
	)
}

func (app *grammar) anyElementToken(tokenName string, elementsList []grammars.Element, suites grammars.Suites) grammars.Token {
	linesList := []grammars.Line{}
	for _, oneElement := range elementsList {
		linesList = append(
			linesList,
			app.lineFromElements([]grammars.Element{
				oneElement,
			}),
		)
	}

	return app.tokenFromBlock(
		tokenName,
		app.blockFromlines(linesList),
		suites,
	)
}

func (app *grammar) tokenFromBlock(name string, block grammars.Block, suites grammars.Suites) grammars.Token {
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

func (app *grammar) blockFromlines(lines []grammars.Line) grammars.Block {
	block, err := app.blockBuilder.Create().
		WithLines(lines).
		Now()

	if err != nil {
		panic(err)
	}

	return block
}

func (app *grammar) lineFromElements(elements []grammars.Element) grammars.Line {
	containers := []grammars.Container{}
	for _, oneElement := range elements {
		container, err := app.containerBuilder.Create().WithElement(oneElement).Now()
		if err != nil {
			panic(err)
		}

		containers = append(containers, container)
	}

	return app.lineFromContainers(containers)
}

func (app *grammar) lineFromContainers(containers []grammars.Container) grammars.Line {
	line, err := app.lineBuilder.Create().
		WithContainers(containers).
		Now()

	if err != nil {
		panic(err)
	}

	return line
}

func (app *grammar) elementFromEverything(everything grammars.Everything) grammars.Element {
	ins, err := app.instanceBuilder.Create().
		WithEverything(everything).
		Now()

	if err != nil {
		panic(err)
	}

	cardinality := app.cardinalityOnce()
	element, err := app.elementBuilder.Create().
		WithInstance(ins).
		WithCardinality(cardinality).
		Now()

	if err != nil {
		panic(err)
	}

	return element
}

func (app *grammar) everythingWithoutEscape(name string, exception grammars.Token) grammars.Everything {
	return app.everything(name, exception, nil)
}

func (app *grammar) everything(name string, exception grammars.Token, escape grammars.Token) grammars.Everything {
	builder := app.everythingBuilder.Create().
		WithName(name).
		WithException(exception)

	if escape != nil {
		builder.WithEscape(escape)
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}

func (app *grammar) elementFromToken(token grammars.Token, cardinality cardinalities.Cardinality) grammars.Element {
	ins, err := app.instanceBuilder.Create().
		WithToken(token).
		Now()

	if err != nil {
		panic(err)
	}

	element, err := app.elementBuilder.Create().
		WithInstance(ins).
		WithCardinality(cardinality).
		Now()

	if err != nil {
		panic(err)
	}

	return element
}

func (app *grammar) elementFromRecursiveToken(tokenName string, cardinality cardinalities.Cardinality) grammars.Element {
	element, err := app.elementBuilder.Create().
		WithRecursive(tokenName).
		WithCardinality(cardinality).
		Now()

	if err != nil {
		panic(err)
	}

	return element
}

func (app *grammar) tokenFromValue(name string, value byte) grammars.Token {
	return app.tokenFromBlock(
		name,
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromValue(value),
			}),
		}),
		app.suites(map[string]bool{}),
	)
}

func (app *grammar) elementFromValue(value byte) grammars.Element {
	valueIns, err := app.valueBuilder.Create().
		WithName(strconv.Itoa(int(value))).
		WithNumber(value).
		Now()

	if err != nil {
		panic(err)
	}

	ins, err := app.elementBuilder.Create().
		WithValue(valueIns).
		WithCardinality(app.cardinalityOnce()).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *grammar) cardinalityOnce() cardinalities.Cardinality {
	max := uint(1)
	return app.cardinality(1, &max)
}

func (app *grammar) cardinality(min uint, pMax *uint) cardinalities.Cardinality {
	builder := app.cardinalityBuilder.Create().WithMin(min)
	if pMax != nil {
		builder.WithMax(*pMax)
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}

func (app *grammar) suites(values map[string]bool) grammars.Suites {
	list := []grammars.Suite{}
	for str, isValid := range values {
		suite := app.suite([]byte(str), isValid)
		list = append(list, suite)
	}

	if len(list) <= 0 {
		return nil
	}

	ins, err := app.suitesBuilder.Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

func (app *grammar) suite(values []byte, isValid bool) grammars.Suite {
	compose := app.compose(values)
	builder := app.suiteBuilder.Create()
	if isValid {
		builder.WithValid(compose)
	}

	if !isValid {
		builder.WithInvalid(compose)
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}

func (app *grammar) compose(values []byte) grammars.Compose {
	elements := []grammars.ComposeElement{}
	for _, oneValue := range values {
		valueIns, err := app.valueBuilder.Create().
			WithName(strconv.Itoa(int(oneValue))).
			WithNumber(oneValue).
			Now()

		if err != nil {
			panic(err)
		}

		element, err := app.composeElementBuilder.Create().
			WithValue(valueIns).
			WithOccurences(1).
			Now()

		if err != nil {
			panic(err)
		}

		elements = append(elements, element)
	}

	ins, err := app.composeBuilder.Create().WithName(string(values)).WithList(elements).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

func (app *grammar) channelFromValue(name string, value byte) grammars.Channel {
	return app.channelFromToken(
		app.tokenFromBlock(
			name,
			app.blockFromlines([]grammars.Line{
				app.lineFromElements([]grammars.Element{
					app.elementFromValue(value),
				}),
			}),
			app.suites(map[string]bool{}),
		),
	)
}

func (app grammar) channelFromToken(token grammars.Token) grammars.Channel {
	ins, err := app.channelBuilder.Create().
		WithToken(token).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app grammar) channels() grammars.Channels {
	list := []grammars.Channel{
		app.channelFromValue("space", []byte(" ")[0]),
		app.channelFromValue("tab", []byte("\t")[0]),
		app.channelFromValue("newLine", []byte("\n")[0]),
		app.channelFromValue("retChar", []byte("\r")[0]),
		app.channelFromToken(
			app.tokenFromBlock("singleLineComment",
				app.blockFromlines([]grammars.Line{
					app.lineFromElements([]grammars.Element{
						app.elementFromToken(
							app.allCharacterToken("doubleSlash", "//"),
							app.cardinalityOnce(),
						),
						app.elementFromEverything(
							app.everythingWithoutEscape(
								"everythingExceptEndOfLine",
								app.anyElementToken(
									"endOfLineSpaces",
									[]grammars.Element{
										app.elementFromValue([]byte("\n")[0]),
										app.elementFromValue([]byte("\r")[0]),
									},
									app.suites(map[string]bool{
										"\n": true,
										"\r": true,
									}),
								),
							),
						),
					}),
				}),
				app.suites(map[string]bool{
					`// this is a comment
					`: true,
				}),
			),
		),
	}

	ins, err := app.channelsBuilder.Create().
		WithList(list).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
