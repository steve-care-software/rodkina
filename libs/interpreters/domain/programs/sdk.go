package programs

import (
	"github.com/steve-care-software/rodkina/libs/interpreters/domain/programs/modules"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewInstructionsBuilder creates a new instructions builder
func NewInstructionsBuilder() InstructionsBuilder {
	return createInstructionsBuilder()
}

// NewInstructionBuilder creates a new instruction builder
func NewInstructionBuilder() InstructionBuilder {
	return createInstructionBuilder()
}

// NewApplicationBuilder creates a new application builder instance
func NewApplicationBuilder() ApplicationBuilder {
	return createApplicationBuilder()
}

// NewAttachmentsBuilder creates a new attachments builder
func NewAttachmentsBuilder() AttachmentsBuilder {
	return createAttachmentsBuilder()
}

// NewAttachmentBuilder creates a new attachment builder
func NewAttachmentBuilder() AttachmentBuilder {
	return createAttachmentBuilder()
}

// NewValueBuilder creates a new value builder
func NewValueBuilder() ValueBuilder {
	return createValueBuilder()
}

// Builder represents a program builder
type Builder interface {
	Create() Builder
	WithInstructions(instructions Instructions) Builder
	WithOutputs(outputs []uint) Builder
	Now() (Program, error)
}

// Program represents a program
type Program interface {
	Instructions() Instructions
	HasOutputs() bool
	Outputs() []uint
}

// InstructionsBuilder represents instructions builder
type InstructionsBuilder interface {
	Create() InstructionsBuilder
	WithList(list []Instruction) InstructionsBuilder
	Now() (Instructions, error)
}

// Instructions represents instructions
type Instructions interface {
	List() []Instruction
}

// InstructionBuilder represents an instruction builder
type InstructionBuilder interface {
	Create() InstructionBuilder
	WithValue(value Value) InstructionBuilder
	WithExecution(execution Application) InstructionBuilder
	Now() (Instruction, error)
}

// Instruction represents an instruction
type Instruction interface {
	IsValue() bool
	Value() Value
	IsExecution() bool
	Execution() Application
}

// ApplicationBuilder represents an application builder
type ApplicationBuilder interface {
	Create() ApplicationBuilder
	WithIndex(index uint) ApplicationBuilder
	WithModule(module modules.Module) ApplicationBuilder
	WithAttachments(attachments Attachments) ApplicationBuilder
	Now() (Application, error)
}

// Application represents an application
type Application interface {
	Index() uint
	Module() modules.Module
	HasAttachments() bool
	Attachments() Attachments
}

// AttachmentsBuilder represents the attachments builder
type AttachmentsBuilder interface {
	Create() AttachmentsBuilder
	WithList(list []Attachment) AttachmentsBuilder
	Now() (Attachments, error)
}

// Attachments represents attachments
type Attachments interface {
	List() []Attachment
}

// AttachmentBuilder represents an attachment builder
type AttachmentBuilder interface {
	Create() AttachmentBuilder
	WithValue(value Value) AttachmentBuilder
	WithLocal(local uint) AttachmentBuilder
	Now() (Attachment, error)
}

// Attachment represents an attachment
type Attachment interface {
	Value() Value
	Local() uint
}

// ValueBuilder represents a value builder
type ValueBuilder interface {
	Create() ValueBuilder
	WithInput(input uint) ValueBuilder
	WithConstant(constant []byte) ValueBuilder
	WithExecution(execution Application) ValueBuilder
	WithProgram(program Program) ValueBuilder
	Now() (Value, error)
}

// Value represents a value
type Value interface {
	IsInput() bool
	Input() *uint
	IsConstant() bool
	Constant() []byte
	IsExecution() bool
	Execution() Application
	IsProgram() bool
	Program() Program
}
