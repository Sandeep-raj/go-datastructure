package recursion

import "log"

/*
Path with Maximum Gold

In a gold mine grid of size m x n, each cell in this mine has an integer representing the amount of gold in that cell, 0 if it is empty.

Return the maximum amount of gold you can collect under the conditions:

Every time you are located in a cell you will collect all the gold in that cell.
From your position, you can walk one step to the left, right, up, or down.
You can't visit the same cell more than once.
Never visit a cell with 0 gold.
You can start and stop collecting gold from any position in the grid that has some gold.

Sample Input
0 6 0
5 8 7
0 9 0

Sample Output
24
*/

func TestMaxGold() {
	inpArr := make([][]int, 6)
	inpArr[0] = []int{0, 1, 4, 2, 8, 2}
	inpArr[1] = []int{4, 3, 6, 5, 0, 4}
	inpArr[2] = []int{1, 2, 4, 1, 4, 6}
	inpArr[3] = []int{2, 0, 7, 3, 2, 2}
	inpArr[4] = []int{3, 1, 5, 9, 2, 4}
	inpArr[5] = []int{2, 7, 0, 8, 5, 1}

	traversePath := make([][]int, len(inpArr))

	for i := 0; i < len(inpArr); i++ {
		traversePath[i] = make([]int, len(inpArr[i]))
		for j := 0; j < len(inpArr[0]); j++ {
			traversePath[i][j] = 0
		}
	}

	maxgold := 0

	for i := 0; i < len(inpArr); i++ {
		for j := 0; j < len(inpArr[0]); j++ {
			currmax := maxGold(i, j, traversePath, inpArr)
			if maxgold < currmax {
				maxgold = currmax
			}
		}
	}

	log.Print(maxgold)

}

func maxGold(row int, col int, traversePath [][]int, mine [][]int) int {
	if traversePath[row][col] == 0 && mine[row][col] > 0 {
		traversePath[row][col] = 1

		top := 0
		left := 0
		right := 0
		bottom := 0

		if row-1 >= 0 {
			top = mine[row][col] + maxGold(row-1, col, traversePath, mine)
		}

		if col-1 >= 0 {
			left = mine[row][col] + maxGold(row, col-1, traversePath, mine)
		}

		if col+1 < len(traversePath[0]) {
			right = mine[row][col] + maxGold(row, col+1, traversePath, mine)
		}

		if row+1 < len(traversePath) {
			bottom = mine[row][col] + maxGold(row+1, col, traversePath, mine)
		}

		traversePath[row][col] = 0

		return getMax([]int{top, left, right, bottom})
	}

	return 0
}

func getMax(inpArr []int) int {
	res := 0

	for i := 0; i < len(inpArr); i++ {
		if res < inpArr[i] {
			res = inpArr[i]
		}
	}
	return res
}
