package block

import (
	"reflect"
	"testing"
)

func TestCreateBlock(t *testing.T) {
	BlockShapes = map[rune][][]int{
		'I': {
			{1, 1, 1, 1},
		},
		'O': {
			{1, 1},
			{1, 1},
		},
	}

	t.Run("Valid Block", func(t *testing.T) {
		letter := 'I'
		block, err := CreateBlock(letter)
		if err != nil {
			t.Fatalf("Expected no error, but got: %v", err)
		}

		if block == nil {
			t.Fatal("Expected a block, but got nil")
		}

		if block.Letter != letter {
			t.Errorf("Expected block letter '%c', but got '%c'", letter, block.Letter)
		}

		if block.Height != len(BlockShapes[letter]) {
			t.Errorf("Expected block height %d, but got %d", len(BlockShapes[letter]), block.Height)
		}

		if block.Width != len(BlockShapes[letter][0]) {
			t.Errorf("Expected block width %d, but got %d", len(BlockShapes[letter][0]), block.Width)
		}
	})

	t.Run("Invalid Block", func(t *testing.T) {
		letter := 'X'
		block, err := CreateBlock(letter)

		if err == nil {
			t.Fatal("Expected an error, but got nil")
		}

		if block != nil {
			t.Errorf("Expected block to be nil for invalid letter '%c', but got: %v", letter, block)
		}
	})
}

func TestLowestHeightPerColumn(t *testing.T) {
	// Create test cases
	tests := []struct {
		name     string
		block    Block
		expected [][2]int
	}{
		{
			name: "Square Block",
			block: Block{
				Shape:  [][]int{{1, 1}, {1, 1}},
				Height: 2,
				Width:  2,
				Letter: 'Q',
			},
			expected: [][2]int{{1, 0}, {1, 1}},
		},
		{
			name: "Line Block",
			block: Block{
				Shape:  [][]int{{1, 1, 1, 1}},
				Height: 1,
				Width:  4,
				Letter: 'I',
			},
			expected: [][2]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
		},
		{
			name: "L Block",
			block: Block{
				Shape:  [][]int{{1, 0}, {1, 0}, {1, 1}},
				Height: 3,
				Width:  2,
				Letter: 'L',
			},
			expected: [][2]int{{2, 0}, {2, 1}},
		},
		{
			name: "Z Block",
			block: Block{
				Shape:  [][]int{{1, 1, 0}, {0, 1, 1}},
				Height: 2,
				Width:  3,
				Letter: 'Z',
			},
			expected: [][2]int{{0, 0}, {1, 1}, {1, 2}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.block.LowestHeightPerColumn()
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("LowestHeightPerColumn() = %v; want %v", result, test.expected)
			}
		})
	}
}
