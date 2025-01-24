package main

import (
	"bufio"
	"fmt"
	"os"
	b "tetris/block"
	g "tetris/grid"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		grid := g.EmptyGrid()

		i := 0
		for i < len(input) {
			// Read the block letter (single character).
			blockLetter := rune(input[i])
			i++

			// Read the offset (digit) and increment twice to skip comma.
			offset := int(input[i] - '0')
			i += 2

			// Process the round with the block letter and offset.
			newGrid, err := simulateRound(grid, blockLetter, offset)

			if err != nil {
				fmt.Println(err.Error())
				return
			}
			grid = *newGrid
		}

		// Output the final grid height after processing all rounds.
		fmt.Println(grid.Height)
	}
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
