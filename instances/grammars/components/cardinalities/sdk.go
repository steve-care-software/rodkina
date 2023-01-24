package cardinalities

import "github.com/steve-care-software/rodkina/libs/asts/domain/grammars/cardinalities"

const byteLength = 256

// NewCardinality creates a new cardinality
func NewCardinality() Cardinality {
	cardinalityBuilder := cardinalities.NewBuilder()
	return createCardinality(cardinalityBuilder)
}

// Cardinality represents the cardinality grammar
type Cardinality interface {
	Once() cardinalities.Cardinality
	Cardinality(min uint, pMax *uint) cardinalities.Cardinality
}
