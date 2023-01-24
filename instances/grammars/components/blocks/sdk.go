package blocks

import "github.com/steve-care-software/rodkina/libs/asts/domain/grammars"

// NewBlock creates a new block
func NewBlock() Block {
	blockBuilder := grammars.NewBlockBuilder()
	return createBlock(blockBuilder)
}

// Block represents a block
type Block interface {
	FromLines(lines []grammars.Line) grammars.Block
}
