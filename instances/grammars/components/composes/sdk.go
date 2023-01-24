package composes

import (
	"github.com/steve-care-software/rodkina/libs/asts/domain/grammars"
	"github.com/steve-care-software/rodkina/libs/asts/domain/grammars/values"
)

// NewCompose creates a new compose instance
func NewCompose() Compose {
	composeBuilder := grammars.NewComposeBuilder()
	composeElementBuilder := grammars.NewComposeElementBuilder()
	valueBuilder := values.NewBuilder()
	return createCompose(
		composeBuilder,
		composeElementBuilder,
		valueBuilder,
	)
}

// Compose represents the compose instance
type Compose interface {
	Compose(values []byte) grammars.Compose
}
