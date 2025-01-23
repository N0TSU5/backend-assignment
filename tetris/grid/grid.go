package grid

import "tetris/block"

const GRID_WIDTH = 10

type Grid struct {
	Cells  [][GRID_WIDTH]int // The 2D array representing the grid itself. A 1 means that a block is on that cell.
	Height int               // The number of rows in the grid, can change when blocks are added.
}

// Creates an empty grid with default width and height of 1.
func EmptyGrid() Grid {
	return Grid{Cells: [][GRID_WIDTH]int{{}}, Height: 1}
}

/*
Drops a block onto the grid at a given horizontal offset.

The block is dropped vertically, potentially changing the grid height.
Returns a new grid, a list of populated rows, and an error if any.

Parameters:
- grid (Grid): The current grid.
- block (Block): The block to drop.
- offset (int): The horizontal offset for the block.

Returns:
- *Grid: The updated grid.
- []int: Rows populated by the block.
- error: An error if an issue occurs.
*/
func DropBlock(grid Grid, block block.Block, offset int) (*Grid, [2]int, error) {
	dropRow := findDropRow(grid, block, offset)
	return dropBlockAtRow(grid, block, offset, dropRow)
}

/*
Filters out full rows within the specified range.

Rows where all cells are `1` are removed, reducing the grid height. If no rows remain, an empty grid is returned.

Parameters:
- grid (Grid): The current grid to filter.
- startRow (int): The starting row index (inclusive).
- endRow (int): The ending row index (exclusive), rows in this range are checked.

Returns:
- *Grid: The filtered grid, or an empty grid if all rows are removed.
- error: Any error encountered during filtering (e.g., invalid row indices).
*/
func FilterRows(grid Grid, startRow int, endRow int) (*Grid, error) {
	newCells := grid.Cells

	for rowIdx := endRow; rowIdx >= startRow; rowIdx-- {
		// Check if all cells in the row are full.
		allOnes := true
		for _, cell := range grid.Cells[rowIdx] {
			if cell != 1 {
				allOnes = false
				break
			}
		}

		if allOnes {
			newCells = append(newCells[:rowIdx], newCells[rowIdx+1:]...)
		}
	}

	if len(newCells) == 0 {
		emptyGrid := EmptyGrid()
		return &emptyGrid, nil
	}

	return &Grid{Cells: newCells, Height: len(newCells)}, nil
}

// Returns the height of the topmost block, -1 if there are no blocks.
func (g *Grid) TopBlockHeight() int {
	for rowIdx := 0; rowIdx < g.Height; rowIdx++ {
		for _, cell := range g.Cells[rowIdx] {
			if cell == 1 {
				return g.Height - rowIdx
			}
		}
	}

	return -1
}
