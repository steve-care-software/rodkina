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

// NewChannel creates a new channel component
func NewChannel() Channel {
	token := components_tokens.NewToken()
	element := components_elements.NewElement()
	block := components_blocks.NewBlock()
	line := components_lines.NewLine()
	suite := components_suites.NewSuite()
	everything := components_everythings.NewEverything()
	cardinality := components_cardinalities.NewCardinality()
	channelsBuilder := grammars.NewChannelsBuilder()
	channelBuilder := grammars.NewChannelBuilder()
	return createChannel(
		token,
		element,
		block,
		line,
		suite,
		everything,
		cardinality,
		channelsBuilder,
		channelBuilder,
	)
}

// Channel represents the channel grammar
type Channel interface {
	Channels() grammars.Channels
}
