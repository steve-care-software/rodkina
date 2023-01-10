package grammars

import (
	"errors"
)

type lineBuilder struct {
	containers []Container
}

func createLineBuilder() LineBuilder {
	out := lineBuilder{
		containers: nil,
	}

	return &out
}

// Create initializes the builder
func (app *lineBuilder) Create() LineBuilder {
	return createLineBuilder()
}

// WithContainers add containers to the builder
func (app *lineBuilder) WithContainers(containers []Container) LineBuilder {
	app.containers = containers
	return app
}

// Now builds a new Line instance
func (app *lineBuilder) Now() (Line, error) {
	if app.containers != nil && len(app.containers) <= 0 {
		app.containers = nil
	}

	if app.containers == nil {
		return nil, errors.New("there must be at least 1 Container in order to build a Line instance")
	}

	return createLine(app.containers), nil
}
