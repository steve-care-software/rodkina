package cardinalities

import (
	"github.com/steve-care-software/rodkina/libs/asts/domain/grammars/cardinalities"
)

type cardinality struct {
	cardinalityBuilder cardinalities.Builder
}

func createCardinality(
	cardinalityBuilder cardinalities.Builder,
) Cardinality {
	out := cardinality{
		cardinalityBuilder: cardinalityBuilder,
	}

	return &out
}

// Once returns acardinality with 1 element
func (app *cardinality) Once() cardinalities.Cardinality {
	max := uint(1)
	return app.Cardinality(1, &max)
}

// Cardinality returns a cardinality instance
func (app *cardinality) Cardinality(min uint, pMax *uint) cardinalities.Cardinality {
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
