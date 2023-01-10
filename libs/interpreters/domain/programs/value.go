package programs

type value struct {
	pInput    *uint
	constant  []byte
	execution Application
	program   Program
}

func createValueWithInput(
	pInput *uint,
) Value {
	return createValueInternally(pInput, nil, nil, nil)
}

func createValueWithConstant(
	constant []byte,
) Value {
	return createValueInternally(nil, constant, nil, nil)
}

func createValueWithExecution(
	execution Application,
) Value {
	return createValueInternally(nil, nil, execution, nil)
}

func createValueWithProgram(
	program Program,
) Value {
	return createValueInternally(nil, nil, nil, program)
}

func createValueInternally(
	pInput *uint,
	constant []byte,
	execution Application,
	program Program,
) Value {
	out := value{
		pInput:    pInput,
		constant:  constant,
		execution: execution,
		program:   program,
	}

	return &out
}

// IsInput returns true if input, false otherwise
func (obj *value) IsInput() bool {
	return obj.pInput != nil
}

// Input returns the input, if any
func (obj *value) Input() *uint {
	return obj.pInput
}

// IsConstant returns true if []byte, false otherwise
func (obj *value) IsConstant() bool {
	return obj.constant != nil
}

// Constant returns the []byte, if any
func (obj *value) Constant() []byte {
	return obj.constant
}

// IsExecution returns true if execution, false otherwise
func (obj *value) IsExecution() bool {
	return obj.execution != nil
}

// Execution returns the execution, if any
func (obj *value) Execution() Application {
	return obj.execution
}

// IsProgram returns true if program, false otherwise
func (obj *value) IsProgram() bool {
	return obj.program != nil
}

// Program returns the program, if any
func (obj *value) Program() Program {
	return obj.program
}
