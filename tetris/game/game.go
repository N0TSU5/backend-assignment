package game

import (
	"fmt"
	"strconv"
	b "tetris/block"
	g "tetris/grid"
)

const GRID_WIDTH = 10

// Simulates a full game given a list of blocks and the offset at which they should be dropped at.
func SimulateGame(grid g.Grid, rounds []string) (*g.Grid, error) {
	for _, round := range rounds {
		blockLetter := rune(round[0])
		offset, err := strconv.Atoi(round[1:])

		if err != nil {
			return nil, fmt.Errorf("error parsing offset:'%s'", round)
		}

		newGrid, err := simulateRound(grid, blockLetter, offset)

		if err != nil {
			return nil, err
		}
		grid = *newGrid
	}

	return &grid, nil
}

// Simulates a round for a given block at a given offset.
func simulateRound(grid g.Grid, letter rune, offset int) (*g.Grid, error) {
	// Create the block with the letter.
	block, err := b.CreateBlock(letter)
	if err != nil {
		return nil, err
	}

	// Drop the block onto the grid.
	newGrid, populatedRows, err := g.DropBlock(grid, *block, offset)

	if err != nil {
		return nil, err
	}
	grid = *newGrid

	// Filter any full rows.
	newGrid, err = g.FilterRows(grid, populatedRows[0], populatedRows[1])

	if err != nil {
		return nil, err
	}
	grid = *newGrid

	// fmt.Println("Grid height after dropping block", string(block.Letter), "at index", offset, ":", grid.Height)
	return &grid, nil
}
