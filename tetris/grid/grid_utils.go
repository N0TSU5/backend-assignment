package grid

import (
	"errors"
	"math"
	b "tetris/block"
)

func findDropRow(grid Grid, block b.Block, offset int) int {
	dropRow := math.MaxInt

	// In a block, the sub-blocks that will collide first will be the lowest sub-blocks in each column.
	// So, here, we find the first row which collides with one of those lowest blocks and return.
	for _, point := range block.LowestHeightPerColumn() {
		colIdx := point[1] + offset
		for rowIdx := 0; rowIdx < grid.Height; rowIdx++ {
			// Calculate the row of the first instance where the cell is full in a given column.
			if grid.Cells[rowIdx][colIdx] == 1 {
				collisionRow := rowIdx - point[0] - 1
				dropRow = int(math.Min(float64(dropRow), float64(collisionRow)))
			}
		}
	}

	if dropRow == math.MaxInt {
		return grid.Height - block.Height
	}

	return dropRow
}

func dropBlockAtRow(grid Grid, block b.Block, offset int, dropRow int) (*Grid, [2]int, error) {
	newCells := grid.Cells

	if dropRow < 0 {
		padding := make([][GRID_WIDTH]int, -dropRow)
		newCells = append(padding, newCells...)
		dropRow = 0
	}

	// The populating of the grid is such that, [0][0] on the block corresponds to [dropRow][offset] for the grid.
	// Each sub-block in the block is then traversed and the difference is used to populate the grid accordingly.
	for rowIdx := 0; rowIdx < block.Height; rowIdx++ {
		for colIdx := 0; colIdx < block.Width; colIdx++ {
			gridCellRow := dropRow + rowIdx
			gridCellColumn := offset + colIdx

			if gridCellColumn >= GRID_WIDTH {
				return nil, [2]int{}, errors.New("cannot drop block, offset goes beyond grid width")
			}

			if block.Shape[rowIdx][colIdx] == 1 {
				newCells[gridCellRow][gridCellColumn] = 1
			}
		}
	}

	return &Grid{Cells: newCells, Height: len(newCells)}, [2]int{dropRow, dropRow + block.Height - 1}, nil
}
