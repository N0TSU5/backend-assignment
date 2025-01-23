# Tetris backend task

This is my implementation of the Tetris engine written in Go. The engine takes in a series of blocks and their drop positions and simulates the game on a grid, adjusting for row completion and block placement. The final output is the height of the tallest column on the grid.

## How to run

### Pre-built Executables

You can directly use the appropriate pre-built executable for your operating system:

- **Linux**: [tetris-linux](./tetris/tetris-linux)
- **macOS**: [tetris-macos](./tetris/tetris-macos)
- **Windows**: [tetris-windows.exe](./tetris/tetris-windows.exe)


Make sure you have Go installed. You can download and install Go from the official website:

- [Install Go](https://golang.org/dl/)

### Cloning the Repository

Clone this repository to your local machine:

```bash
git clone https://github.com/your-username/tetris-simulator.git
cd tetris-simulator
