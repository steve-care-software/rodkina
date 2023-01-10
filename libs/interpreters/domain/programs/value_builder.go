package programs

import (
	"errors"
)

type valueBuilder struct {
	pInput    *uint
	constant  []byte
	execution Application
	program   Program
}

func createValueBuilder() ValueBuilder {
	out := valueBuilder{
		pInput:    nil,
		constant:  nil,
		execution: nil,
		program:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *valueBuilder) Create() ValueBuilder {
	return createValueBuilder()
}

// WithInput adds an input to the builder
func (app *valueBuilder) WithInput(input uint) ValueBuilder {
	app.pInput = &input
	return app
}

// WithConstant adds a constant to the builder
func (app *valueBuilder) WithConstant(constant []byte) ValueBuilder {
	app.constant = constant
	return app
}

// WithExecution adds an execution to the builder
func (app *valueBuilder) WithExecution(execution Application) ValueBuilder {
	app.execution = execution
	return app
}

// WithProgram adds a program to the builder
func (app *valueBuilder) WithProgram(program Program) ValueBuilder {
	app.program = program
	return app
}

// Now builds a new Value instance
func (app *valueBuilder) Now() (Value, error) {
	if app.pInput != nil {
		return createValueWithInput(app.pInput), nil
	}

	if app.constant != nil {
		return createValueWithConstant(app.constant), nil
	}

	if app.execution != nil {
		return createValueWithExecution(app.execution), nil
	}

	if app.program != nil {
		return createValueWithProgram(app.program), nil
	}

	return nil, errors.New("the Value is invalid")
}
