# Tetris backend task

This is my implementation of the Tetris engine written in Go. The engine takes in a series of blocks and their drop positions and simulates the game on a grid, adjusting for row completion and block placement. The final output is the height of the tallest column on the grid.

## How to run

### Pre-built Executables

You can directly use the appropriate pre-built executable for your operating system:

- **Linux**: [tetris-linux](./tetris/tetris-linux)
- **macOS**: [tetris-macos](./tetris/tetris-macos)
- **Windows**: [tetris-windows.exe](./tetris/tetris-windows.exe)

Once you have the executable for your platform, follow these steps:

1. Open a terminal or command prompt.
2. Navigate to the directory where the executable is located.
3. Run the following command:

   ```bash
   ./tetris-linux   # for Linux
   ./tetris-macos   # for macOS
   tetris-windows.exe  # for Windows
Make sure your inputs are valid characters and offsets separated by a comma e.g. Q3,L5,I0
### If the executables don't work
You can download and install Go from the official website:

- [Install Go](https://go.dev/doc/install)

Then, build the project:
  ```bash
   cd tetris && go build -o tetris
  ```
You should now have an executable `tetris`
## Testing with automated script
You can test with the automated script by editing line 15 `ENTRY_POINT` of [sample_test](./tests/sample_test.py) to your respective executable.
