package blocks

import "github.com/steve-care-software/rodkina/libs/asts/domain/grammars"

type block struct {
	blockBuilder grammars.BlockBuilder
}

func createBlock(
	blockBuilder grammars.BlockBuilder,
) Block {
	out := block{
		blockBuilder: blockBuilder,
	}

	return &out
}

// FromLines returns a block from lines
func (app *block) FromLines(lines []grammars.Line) grammars.Block {
	block, err := app.blockBuilder.Create().
		WithLines(lines).
		Now()

	if err != nil {
		panic(err)
	}

	return block
}
