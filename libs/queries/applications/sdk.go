package queries

import (
	"github.com/steve-care-software/rodkina/libs/asts/domain/grammars"
	"github.com/steve-care-software/rodkina/libs/asts/domain/trees"
	"github.com/steve-care-software/rodkina/libs/queries/domain/queries"
)

// NewApplication creates a new application instance
func NewApplication() Application {
	return createApplication()
}

// Application represents a query application
type Application interface {
	Matches(grammar grammars.Grammar, query queries.Query) (bool, error)
	Execute(query queries.Query, treeIns trees.Tree) (interface{}, bool, []byte, error)
}
