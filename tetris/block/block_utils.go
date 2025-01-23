package block

var BlockShapes = map[rune][][]int{
	'Q': { // Square block
		{1, 1},
		{1, 1},
	},
	'Z': { // Z block
		{1, 1, 0},
		{0, 1, 1},
	},
	'S': { // S block
		{0, 1, 1},
		{1, 1, 0},
	},
	'T': { // T block
		{1, 1, 1},
		{0, 1, 0},
	},
	'I': { // Line block
		{1, 1, 1, 1},
	},
	'L': { // L block
		{1, 0},
		{1, 0},
		{1, 1},
	},
	'J': { // J block
		{0, 1},
		{0, 1},
		{1, 1},
	},
}
