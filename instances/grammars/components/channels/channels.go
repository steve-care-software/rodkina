package channels

import (
	components_blocks "github.com/steve-care-software/rodkina/instances/grammars/components/blocks"
	components_cardinalities "github.com/steve-care-software/rodkina/instances/grammars/components/cardinalities"
	components_elements "github.com/steve-care-software/rodkina/instances/grammars/components/elements"
	components_everythings "github.com/steve-care-software/rodkina/instances/grammars/components/everythings"
	components_lines "github.com/steve-care-software/rodkina/instances/grammars/components/lines"
	components_suites "github.com/steve-care-software/rodkina/instances/grammars/components/suites"
	components_tokens "github.com/steve-care-software/rodkina/instances/grammars/components/tokens"
	"github.com/steve-care-software/rodkina/libs/asts/domain/grammars"
)

type channel struct {
	token           components_tokens.Token
	element         components_elements.Element
	block           components_blocks.Block
	line            components_lines.Line
	suite           components_suites.Suite
	everything      components_everythings.Everything
	cardinality     components_cardinalities.Cardinality
	channelsBuilder grammars.ChannelsBuilder
	channelBuilder  grammars.ChannelBuilder
}

func createChannel(
	token components_tokens.Token,
	element components_elements.Element,
	block components_blocks.Block,
	line components_lines.Line,
	suite components_suites.Suite,
	everything components_everythings.Everything,
	cardinality components_cardinalities.Cardinality,
	channelsBuilder grammars.ChannelsBuilder,
	channelBuilder grammars.ChannelBuilder,
) Channel {
	out := channel{
		token:           token,
		element:         element,
		block:           block,
		line:            line,
		suite:           suite,
		everything:      everything,
		cardinality:     cardinality,
		channelsBuilder: channelsBuilder,
		channelBuilder:  channelBuilder,
	}

	return &out
}

// Channels returns the channels
func (app *channel) Channels() grammars.Channels {
	list := []grammars.Channel{
		app.channelFromValue("space", []byte(" ")[0]),
		app.channelFromValue("tab", []byte("\t")[0]),
		app.channelFromValue("newLine", []byte("\n")[0]),
		app.channelFromValue("retChar", []byte("\r")[0]),
		app.channelFromToken(
			app.token.FromBlock("singleLineComment",
				app.block.FromLines([]grammars.Line{
					app.line.FromElements([]grammars.Element{
						app.element.FromToken(
							app.token.AllCharacters("doubleSlash", "//"),
							app.cardinality.Once(),
						),
						app.element.FromEverything(
							app.everything.WithoutEscape(
								"everythingExceptEndOfLine",
								app.token.AnyElement(
									"endOfLineSpaces",
									[]grammars.Element{
										app.element.FromValue([]byte("\n")[0]),
										app.element.FromValue([]byte("\r")[0]),
									},
									app.suite.Suites(map[string]bool{
										"\n": true,
										"\r": true,
									}),
								),
							),
						),
					}),
				}),
				app.suite.Suites(map[string]bool{
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

func (app *channel) channelFromValue(name string, value byte) grammars.Channel {
	return app.channelFromToken(
		app.token.FromBlock(
			name,
			app.block.FromLines([]grammars.Line{
				app.line.FromElements([]grammars.Element{
					app.element.FromValue(value),
				}),
			}),
			app.suite.Suites(map[string]bool{}),
		),
	)
}

func (app channel) channelFromToken(token grammars.Token) grammars.Channel {
	ins, err := app.channelBuilder.Create().
		WithToken(token).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
