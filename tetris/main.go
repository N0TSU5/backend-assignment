package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tetris/game"
	g "tetris/grid"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// Extract the blocks and offsets from STDIN into rounds.
		line := strings.TrimSpace(scanner.Text())
		rounds := strings.Split(line, ",")

		// Create an empty grid and start the game with the rounds.
		grid := g.EmptyGrid()
		newGrid, err := game.SimulateGame(grid, rounds)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		grid = *newGrid

		// STDOUT the new grid's height.
		fmt.Println(grid.Height)
	}
}
