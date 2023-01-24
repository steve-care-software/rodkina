package everythings

import "github.com/steve-care-software/rodkina/libs/asts/domain/grammars"

type everything struct {
	everythingBuilder grammars.EverythingBuilder
}

func createEverything(
	everythingBuilder grammars.EverythingBuilder,
) Everything {
	out := everything{
		everythingBuilder: everythingBuilder,
	}

	return &out
}

// WithoutEscape returns an everything without escape
func (app *everything) WithoutEscape(name string, exception grammars.Token) grammars.Everything {
	return app.Everything(name, exception, nil)
}

// Everything returns an everything instance
func (app *everything) Everything(name string, exception grammars.Token, escape grammars.Token) grammars.Everything {
	builder := app.everythingBuilder.Create().
		WithName(name).
		WithException(exception)

	if escape != nil {
		builder.WithEscape(escape)
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}
