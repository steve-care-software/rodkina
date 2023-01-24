package suites

import (
	"github.com/steve-care-software/rodkina/instances/grammars/components/composes"
	"github.com/steve-care-software/rodkina/libs/asts/domain/grammars"
)

type suite struct {
	compose       composes.Compose
	suitesBuilder grammars.SuitesBuilder
	suiteBuilder  grammars.SuiteBuilder
}

func createSuite(
	compose composes.Compose,
	suitesBuilder grammars.SuitesBuilder,
	suiteBuilder grammars.SuiteBuilder,
) Suite {
	out := suite{
		compose:       compose,
		suitesBuilder: suitesBuilder,
		suiteBuilder:  suiteBuilder,
	}

	return &out
}

// Suites creates suites based on the values
func (app *suite) Suites(values map[string]bool) grammars.Suites {
	list := []grammars.Suite{}
	for str, isValid := range values {
		suite := app.suite([]byte(str), isValid)
		list = append(list, suite)
	}

	if len(list) <= 0 {
		return nil
	}

	ins, err := app.suitesBuilder.Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

func (app *suite) suite(values []byte, isValid bool) grammars.Suite {
	compose := app.compose.Compose(values)
	builder := app.suiteBuilder.Create()
	if isValid {
		builder.WithValid(compose)
	}

	if !isValid {
		builder.WithInvalid(compose)
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}
