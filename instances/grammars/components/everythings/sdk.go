package everythings

import "github.com/steve-care-software/rodkina/libs/asts/domain/grammars"

// NewEverything creates a new everything component
func NewEverything() Everything {
	everythingBuilder := grammars.NewEverythingBuilder()
	return createEverything(everythingBuilder)
}

// Everything represents an everything component
type Everything interface {
	WithoutEscape(name string, exception grammars.Token) grammars.Everything
	Everything(name string, exception grammars.Token, escape grammars.Token) grammars.Everything
}
