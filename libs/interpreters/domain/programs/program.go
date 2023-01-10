package programs

type program struct {
	instructions Instructions
	outputs      []uint
}

func createProgram(
	instructions Instructions,
) Program {
	return createProgramInternally(instructions, nil)
}

func createProgramWithOutputs(
	instructions Instructions,
	outputs []uint,
) Program {
	return createProgramInternally(instructions, outputs)
}

func createProgramInternally(
	instructions Instructions,
	outputs []uint,
) Program {
	out := program{
		instructions: instructions,
		outputs:      outputs,
	}

	return &out
}

// Instructions returns the instructions
func (obj *program) Instructions() Instructions {
	return obj.instructions
}

// HasOutputs returns true if there is outputs, false otherwise
func (obj *program) HasOutputs() bool {
	return obj.outputs != nil
}

// Outputs returns the outputs, if any
func (obj *program) Outputs() []uint {
	return obj.outputs
}
