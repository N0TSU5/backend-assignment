package grid

import (
	"reflect"
	"testing"
	b "tetris/block"
)

func TestDropBlock(t *testing.T) {
	tests := []struct {
		name         string
		initialGrid  Grid
		block        b.Block
		offset       int
		expectedGrid Grid
		expectedRows [2]int
	}{
		{
			name: "Drop block within grid bounds",
			initialGrid: Grid{
				Cells: [][10]int{
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				},
				Height: 1,
			},
			block: b.Block{
				Shape: [][]int{
					{1, 1},
					{1, 1},
				},
				Height: 2,
				Width:  2,
				Letter: 'Q',
			},
			offset: 1,
			expectedGrid: Grid{
				Cells: [][10]int{
					{0, 1, 1, 0, 0, 0, 0, 0, 0, 0},
					{0, 1, 1, 0, 0, 0, 0, 0, 0, 0},
				},
				Height: 2,
			},
			expectedRows: [2]int{0, 1},
		},
		{
			name: "Drop block at top of non-empty grid",
			initialGrid: Grid{
				Cells: [][10]int{
					{1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
					{1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				},
				Height: 2,
			},
			block: b.Block{
				Shape: [][]int{
					{1, 1},
					{1, 1},
				},
				Height: 2,
				Width:  2,
				Letter: 'Q',
			},
			offset: 1,
			expectedGrid: Grid{
				Cells: [][10]int{
					{0, 1, 1, 0, 0, 0, 0, 0, 0, 0},
					{0, 1, 1, 0, 0, 0, 0, 0, 0, 0},
					{1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
					{1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				},
				Height: 4,
			},
			expectedRows: [2]int{0, 1},
		},
	}

	// Run test cases
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			newGrid, populatedRows, _ := DropBlock(test.initialGrid, test.block, test.offset)

			if !reflect.DeepEqual(newGrid, &test.expectedGrid) {
				t.Errorf("Expected grid: %+v, got: %+v", test.expectedGrid, newGrid)
			}

			if populatedRows != test.expectedRows {
				t.Errorf("Expected populated rows: %v, got: %v", test.expectedRows, populatedRows)
			}
		})
	}
}

func TestFilterRows(t *testing.T) {
	tests := []struct {
		name         string
		initialGrid  Grid
		startRow     int
		endRow       int
		expectedGrid Grid
	}{
		{
			name: "No full rows to filter",
			initialGrid: Grid{
				Cells: [][10]int{
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
				},
				Height: 2,
			},
			startRow: 0,
			endRow:   1,
			expectedGrid: Grid{
				Cells: [][10]int{
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
				},
				Height: 2,
			},
		},
		{
			name: "Single full row filtered",
			initialGrid: Grid{
				Cells: [][10]int{
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, // Full row
					{0, 1, 1, 0, 0, 0, 0, 0, 0, 0},
				},
				Height: 2,
			},
			startRow: 0,
			endRow:   1,
			expectedGrid: Grid{
				Cells: [][10]int{
					{0, 1, 1, 0, 0, 0, 0, 0, 0, 0},
				},
				Height: 1,
			},
		},
		{
			name: "All rows filtered",
			initialGrid: Grid{
				Cells: [][10]int{
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, // Full row
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, // Full row
				},
				Height: 2,
			},
			startRow: 0,
			endRow:   1,
			expectedGrid: Grid{
				Cells:  [][10]int{{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
				Height: 1,
			},
		},
		{
			name: "Partial filtering with multiple full rows",
			initialGrid: Grid{
				Cells: [][10]int{
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				},
				Height: 4,
			},
			startRow: 0,
			endRow:   2,
			expectedGrid: Grid{
				Cells: [][10]int{
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				},
				Height: 3,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			filteredGrid, _ := FilterRows(test.initialGrid, test.startRow, test.endRow)
			if !reflect.DeepEqual(filteredGrid, &test.expectedGrid) {
				t.Errorf("Filtered grid mismatch: got %+v, want %+v", filteredGrid, &test.expectedGrid)
			}
		})
	}
}

func TestTopBlockHeight(t *testing.T) {
	tests := []struct {
		name     string
		grid     *Grid
		expected int
	}{
		{
			name: "Empty grid",
			grid: &Grid{
				Height: 3,
				Cells: [][10]int{
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				},
			},
			expected: -1,
		},
		{
			name: "Single block at bottom row",
			grid: &Grid{
				Height: 3,
				Cells: [][10]int{
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				},
			},
			expected: 1,
		},
		{
			name: "Single block at middle row",
			grid: &Grid{
				Height: 3,
				Cells: [][10]int{
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				},
			},
			expected: 2,
		},
		{
			name: "Single block at top row",
			grid: &Grid{
				Height: 3,
				Cells: [][10]int{
					{0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				},
			},
			expected: 3,
		},
		{
			name: "Multiple blocks in grid",
			grid: &Grid{
				Height: 3,
				Cells: [][10]int{
					{0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
					{1, 0, 1, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				},
			},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.grid.TopBlockHeight()
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}
