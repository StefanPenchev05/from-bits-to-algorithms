package main

// Path
var mazePath [][]int = [][]int{}

func solveMaze(maze [][]int, row, col int) {
	// Add 2 seconds of sleep here
	// time.Sleep(2 * time.Second)
}

func main() {
	maze := [][]int{
		{1, 0, 1, 1},
		{1, 1, 1, 1},
		{0, 1, 1, 0},
		{0, 0, 1, 1},
	}

	solveMaze(maze, 0, 0)
}
