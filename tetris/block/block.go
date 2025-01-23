package block

import (
	"fmt"
)

type Block struct {
	Shape         [][]int //The 2D array representing the block itself. A 1 represents a 'sub-block' of the entire block.
	Height, Width int     //The dimensions of the block.
	Letter        rune    //The letter.
}

// Creates a new block based on the provided letter.
func CreateBlock(letter rune) (*Block, error) {
	shape, exists := BlockShapes[letter]

	if !exists {
		return nil, fmt.Errorf("block with letter '%c' does not exist in BlockShapes", letter)
	}

	return &Block{
		Shape:  shape,
		Height: len(shape),
		Width:  len(shape[0]),
		Letter: letter,
	}, nil
}

// Finds the lowest filled cell for each column in the block.
func (b *Block) LowestHeightPerColumn() [][2]int {
	var lowPoints [][2]int

	// For each column, start from the bottom and append the first 'full' cell.
	for colIdx := 0; colIdx < b.Width; colIdx++ {
		for rowIdx := b.Height - 1; rowIdx >= 0; rowIdx-- {
			if b.Shape[rowIdx][colIdx] == 1 {
				lowPoints = append(lowPoints, [2]int{rowIdx, colIdx})
				break
			}
		}
		continue
	}

	return lowPoints
}
