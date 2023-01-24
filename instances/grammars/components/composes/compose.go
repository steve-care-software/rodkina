package composes

import (
	"strconv"

	"github.com/steve-care-software/rodkina/libs/asts/domain/grammars"
	"github.com/steve-care-software/rodkina/libs/asts/domain/grammars/values"
)

type compose struct {
	composeBuilder        grammars.ComposeBuilder
	composeElementBuilder grammars.ComposeElementBuilder
	valueBuilder          values.Builder
}

func createCompose(
	composeBuilder grammars.ComposeBuilder,
	composeElementBuilder grammars.ComposeElementBuilder,
	valueBuilder values.Builder,
) Compose {
	out := compose{
		composeBuilder:        composeBuilder,
		composeElementBuilder: composeElementBuilder,
		valueBuilder:          valueBuilder,
	}

	return &out
}

// Compose returns the compose instance
func (app *compose) Compose(values []byte) grammars.Compose {
	elements := []grammars.ComposeElement{}
	for _, oneValue := range values {
		valueIns, err := app.valueBuilder.Create().
			WithName(strconv.Itoa(int(oneValue))).
			WithNumber(oneValue).
			Now()

		if err != nil {
			panic(err)
		}

		element, err := app.composeElementBuilder.Create().
			WithValue(valueIns).
			WithOccurences(1).
			Now()

		if err != nil {
			panic(err)
		}

		elements = append(elements, element)
	}

	ins, err := app.composeBuilder.Create().WithName(string(values)).WithList(elements).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
