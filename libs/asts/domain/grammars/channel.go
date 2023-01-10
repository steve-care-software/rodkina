package grammars

type channel struct {
	token     Token
	condition ChannelCondition
}

func createChannel(
	token Token,
) Channel {
	return createChannelInternally(token, nil)
}

func createChannelWithCondition(
	token Token,
	condition ChannelCondition,
) Channel {
	return createChannelInternally(token, condition)
}

func createChannelInternally(
	token Token,
	condition ChannelCondition,
) Channel {
	out := channel{
		token:     token,
		condition: condition,
	}

	return &out
}

// Token returns the token
func (obj *channel) Token() Token {
	return obj.token
}

// HasCondition returns true if there is a condition, false otherwise
func (obj *channel) HasCondition() bool {
	return obj.condition != nil
}

// Condition returns the condition, if any
func (obj *channel) Condition() ChannelCondition {
	return obj.condition
}
