package dynamicprog

import (
	"log"
	"math"
)

/*
Min Cost In Maze Traversal

1. You are given a number n, representing the number of rows.
2. You are given a number m, representing the number of columns.
3. You are given n*m numbers, representing elements of 2d array a, which represents a maze.
4. You are standing in top-left cell and are required to move to bottom-right cell.
5. You are allowed to move 1 cell right (h move) or 1 cell down (v move) in 1 motion.
6. Each cell has a value that will have to be paid to enter that cell (even for the top-left and bottom-
     right cell).
7. You are required to traverse through the matrix and print the cost of path which is least costly.

Sample Input
6
6
0 1 4 2 8 2
4 3 6 5 0 4
1 2 4 1 4 6
2 0 7 3 2 2
3 1 5 9 2 4
2 7 0 8 5 1

Sample Output
23
*/

func TestMinCostMazeTraversal() {
	maze := make([][]int, 6)
	maze[0] = []int{0, 1, 4, 2, 8, 2}
	maze[1] = []int{4, 3, 6, 5, 0, 4}
	maze[2] = []int{1, 2, 4, 1, 4, 6}
	maze[3] = []int{2, 0, 7, 3, 2, 2}
	maze[4] = []int{3, 1, 5, 9, 2, 4}
	maze[5] = []int{2, 7, 0, 8, 5, 1}

	res := minmazeTraversal(maze)
	log.Print(res)

}

func minmazeTraversal(maze [][]int) int {
	row := len(maze)
	col := len(maze[0])

	minCostArr := make([][]int, row)

	for i := 0; i < row; i++ {
		minCostArr[i] = make([]int, col)

		for j := 0; j < col; j++ {
			minCostArr[i][j] = maze[i][j]
		}
	}

	for i := row - 1; i >= 0; i-- {
		for j := col - 1; j >= 0; j-- {
			bottom := math.MaxInt
			right := math.MaxInt

			if i+1 < row {
				bottom = minCostArr[i][j] + minCostArr[i+1][j]
			}

			if j+1 < col {
				right = minCostArr[i][j] + minCostArr[i][j+1]
			}

			if bottom == right && bottom == math.MaxInt {
				continue
			} else {
				if bottom < right {
					minCostArr[i][j] = bottom
				} else {
					minCostArr[i][j] = right
				}
			}
		}
	}

	return minCostArr[0][0]
}
