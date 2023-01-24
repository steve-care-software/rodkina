package elements

import (
	"strconv"

	components_cardinalities "github.com/steve-care-software/rodkina/instances/grammars/components/cardinalities"
	"github.com/steve-care-software/rodkina/libs/asts/domain/grammars"
	"github.com/steve-care-software/rodkina/libs/asts/domain/grammars/cardinalities"
	"github.com/steve-care-software/rodkina/libs/asts/domain/grammars/values"
)

type element struct {
	cardinality     components_cardinalities.Cardinality
	instanceBuilder grammars.InstanceBuilder
	elementBuilder  grammars.ElementBuilder
	valueBuilder    values.Builder
}

func createElement(
	cardinality components_cardinalities.Cardinality,
	instanceBuilder grammars.InstanceBuilder,
	elementBuilder grammars.ElementBuilder,
	valueBuilder values.Builder,
) Element {
	out := element{
		cardinality:     cardinality,
		instanceBuilder: instanceBuilder,
		elementBuilder:  elementBuilder,
		valueBuilder:    valueBuilder,
	}

	return &out
}

// FromEverything returns the element from everything
func (app *element) FromEverything(everything grammars.Everything) grammars.Element {
	ins, err := app.instanceBuilder.Create().
		WithEverything(everything).
		Now()

	if err != nil {
		panic(err)
	}

	cardinality := app.cardinality.Once()
	element, err := app.elementBuilder.Create().
		WithInstance(ins).
		WithCardinality(cardinality).
		Now()

	if err != nil {
		panic(err)
	}

	return element
}

// FromToken retruns an element from token
func (app *element) FromToken(token grammars.Token, cardinality cardinalities.Cardinality) grammars.Element {
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

// FromValue creates an element from value
func (app *element) FromValue(value byte) grammars.Element {
	valueIns, err := app.valueBuilder.Create().
		WithName(strconv.Itoa(int(value))).
		WithNumber(value).
		Now()

	if err != nil {
		panic(err)
	}

	ins, err := app.elementBuilder.Create().
		WithValue(valueIns).
		WithCardinality(app.cardinality.Once()).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
