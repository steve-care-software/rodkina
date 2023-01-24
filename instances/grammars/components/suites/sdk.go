package suites

import (
	"github.com/steve-care-software/rodkina/instances/grammars/components/composes"
	"github.com/steve-care-software/rodkina/libs/asts/domain/grammars"
)

// NewSuite creates the suite instance
func NewSuite() Suite {
	compose := composes.NewCompose()
	suitesBuilder := grammars.NewSuitesBuilder()
	suiteBuilder := grammars.NewSuiteBuilder()
	return createSuite(
		compose,
		suitesBuilder,
		suiteBuilder,
	)
}

// Suite represents the suites instance
type Suite interface {
	Suites(values map[string]bool) grammars.Suites
}
