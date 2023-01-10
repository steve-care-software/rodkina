package applications

import "github.com/steve-care-software/rodkina/libs/interpreters/domain/instructions/parameters"

type parameter struct {
	allParameterIndex   uint
	inputParameterIndex uint
	parameter           parameters.Parameter
}
