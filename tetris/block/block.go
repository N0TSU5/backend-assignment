package block

// Package block provides utilities for managing and creating game blocks in a grid-based system.

import (
	"fmt"
)

type Block struct {
	Shape         [][]int //The 2D array representing the block itself. A 1 represents a 'sub-block' of the entire block.
	Height, Width int     //The dimensions of the block.
	Letter        rune    //The letter.
}

// Creates a new block based on the provided letter from BlockShapes.
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

/*
LowestHeightPerColumn finds the lowest filled cell for each column in the block and returns their positions.

This method iterates through each column of the block's shape matrix from the bottom row upwards and identifies the first filled cell in each column.
It collects the positions of these cells and returns them as a slice of [2]int, where each pair represents the row and column indices of a filled cell.

One of the returned cells will be the first to collide with existing blocks in the grid when dropped.

Returns:
  - [][2]int: A slice of integer pairs, where each pair represents the (row, column) indices of the lowest filled cell in each column.
    If a column has no filled cells, it is omitted from the result.
*/
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
