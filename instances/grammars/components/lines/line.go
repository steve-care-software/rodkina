package lines

import "github.com/steve-care-software/rodkina/libs/asts/domain/grammars"

type line struct {
	lineBuilder      grammars.LineBuilder
	containerBuilder grammars.ContainerBuilder
}

func createLine(
	lineBuilder grammars.LineBuilder,
	containerBuilder grammars.ContainerBuilder,
) Line {
	out := line{
		lineBuilder:      lineBuilder,
		containerBuilder: containerBuilder,
	}

	return &out
}

// FromElements returns a line from elements
func (app *line) FromElements(elements []grammars.Element) grammars.Line {
	containers := []grammars.Container{}
	for _, oneElement := range elements {
		container, err := app.containerBuilder.Create().WithElement(oneElement).Now()
		if err != nil {
			panic(err)
		}

		containers = append(containers, container)
	}

	return app.lineFromContainers(containers)
}

func (app *line) lineFromContainers(containers []grammars.Container) grammars.Line {
	line, err := app.lineBuilder.Create().
		WithContainers(containers).
		Now()

	if err != nil {
		panic(err)
	}

	return line
}
