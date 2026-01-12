package main

import "fmt"

func getMaxRows(maze [][]int) int {
	return len(maze)
}

func getMaxCols(maze [][]int) int {
	return len(maze[0])
}

/*
isMoveValid checksif the given coordinates represent a valid move

Parameters:

  - nr: next row coordinate to check

  - nc: next column coordiante to check

  - maze: 2D array

    Returns: boolean indicating if the move is valid
*/
func isMoveValid(nr, nc int, maze [][]int) bool {
	// rowsCount stroes the total number of rows in the array/maze
	var rowsCount int = getMaxRows(maze)

	// colsCount stroes the total number of columns in the array/maze
	var colsCount int = getMaxCols(maze)

	if nr >= 0 && nr < rowsCount && nc >= 0 && nc < colsCount && maze[nr][nc] != 0 {
		return true
	}

	return false
}

func nextMove(row, col int, moveCounter, pathCyles *int, maze [][]int, successPath *[][]int) bool {
	// If found the correct path return true
	if row == getMaxRows(maze)-1 && col == getMaxRows(maze)-1 {
		return true
	}

	// If all 4 types of movements are invalid return that is not the correct path = false
	if *moveCounter == 4 {
		// Mark the current cell as not visited
		maze[row][col] = 1
		return false
	}

	// Mark the current cell as visited
	*successPath = append(*successPath, []int{row, col})
	maze[row][col] = 0

	// rowMove defines the row direction changes for Up, Right, Left, Down movements
	// Up: +1, Right: 0, Left: 0, Down: -1
	var rowMove = [...]int{1, 0, 0, -1}

	// colMove defines the column direction changes for Up, Right, Left, Down movements
	// Up: 0, Right: +1, Left: -1, Down: 0
	var colMove = [...]int{0, 1, -1, 0}

	var nr, nc = row + rowMove[*moveCounter], col + colMove[*moveCounter]

	if isMoveValid(nr, nc, maze) {
		// If the next movement is valid - procceed
		nextMove(nr, nc, moveCounter, pathCyles, maze, successPath)
	} else {
		// Otherwise, try different movement
		*moveCounter++
		nextMove(nr, nc, moveCounter, pathCyles, maze, successPath)
	}

	return false
}

// solveMaze recursively finds a path through the maze and stores it in successPath
func solveMaze(row, col int, pathCycles *int, maze [][]int, successPath *[][]int) bool {
	// Lets make the big problme into small ones
	// First small problem is the movement without loops, we need toi think of a way using only recursion

	// First we have 4 directions - Upside, Downside, Right and Left - we need to go throught to every one

	var startRow int = row
	var startCol int = col

	var moveCounter int = 0

	if nextMove(startRow, startCol, &moveCounter, pathCycles, maze, successPath) {
		// If the success path is found return true
		return true
	} else {
		// Otherwise, return to the position that the player started from and try a diffrent path
		*pathCycles++
		solveMaze(startRow, startCol, pathCycles, maze, successPath)
	}

	return true
}

func main() {
	maze := [][]int{
		{1, 0, 1, 1},
		{1, 1, 1, 1},
		{0, 1, 1, 0},
		{0, 0, 1, 1},
	}

	// The success path as sequence of coordinates - 2D array
	var successPath [][]int = [][]int{}

	// How many times did the program returned to the start state, because encontered dead end
	var pathCycles int = 0

	if solveMaze(0, 0, &pathCycles, maze, &successPath) {
		fmt.Printf("✅ Path found: %v\n", successPath)
	} else {
		fmt.Println(successPath)
		fmt.Printf("❌ No path exists\n")
	}
}
