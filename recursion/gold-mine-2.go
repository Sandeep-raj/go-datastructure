package recursion

import "log"

/*
Gold Mine - 2
1. You are given a number n, representing the number of rows.
2. You are given a number m, representing the number of columns.
3. You are given n*m numbers, representing elements of 2d array a, which represents a gold mine.
4. You are allowed to take one step left, right, up or down from your current position.
5. You can't visit a cell with 0 gold and the same cell more than once.
6. Each cell has a value that is the amount of gold available in the cell.
7. You are required to identify the maximum amount of gold that can be dug out from the mine if
     you start and stop collecting gold from any position in the grid that has some gold.

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
120
*/

func TestGoldMine() {

	n := 7
	m := 7
	inpArr := make([][]int, n)
	inpArr[0] = []int{10, 0, 100, 200, 0, 8, 0}
	inpArr[1] = []int{20, 0, 0, 0, 0, 6, 0}
	inpArr[2] = []int{30, 0, 0, 9, 12, 3, 4}
	inpArr[3] = []int{40, 0, 2, 5, 8, 3, 11}
	inpArr[4] = []int{0, 0, 0, 0, 0, 9, 0}
	inpArr[5] = []int{5, 6, 7, 0, 7, 4, 2}
	inpArr[6] = []int{8, 9, 10, 0, 1, 10, 8}

	maxVal := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if inpArr[i][j] != 0 {
				curMax := goldMine(inpArr, i, j)
				if curMax > maxVal {
					maxVal = curMax
				}
			}
		}
	}

	log.Print(maxVal)
}

func goldMine(inpArr [][]int, r int, c int) int {
	if r < 0 || r >= len(inpArr) || c < 0 || c >= len(inpArr[0]) || inpArr[r][c] == 0 {
		return 0
	}

	currVal := inpArr[r][c]
	inpArr[r][c] = 0
	return currVal + goldMine(inpArr, r, c-1) + goldMine(inpArr, r-1, c) + goldMine(inpArr, r, c+1) + goldMine(inpArr, r+1, c)
}
