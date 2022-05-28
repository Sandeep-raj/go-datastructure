package dynamicprog

import "log"

/*
Goldmine

1. You are given a number n, representing the number of rows.
2. You are given a number m, representing the number of columns.
3. You are given n*m numbers, representing elements of 2d array a, which represents a gold mine.
4. You are standing in front of left wall and are supposed to dig to the right wall. You can start from
     any row in the left wall.
5. You are allowed to move 1 cell right-up (d1), 1 cell right (d2) or 1 cell right-down(d3).
6. Each cell has a value that is the amount of gold available in the cell.
7. You are required to identify the maximum amount of gold that can be dug out from the mine.


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
33
*/

func TestGoldMine() {
	mine := make([][]int, 6)
	mine[0] = []int{0, 1, 4, 2, 8, 2}
	mine[1] = []int{4, 3, 6, 5, 0, 4}
	mine[2] = []int{1, 2, 4, 1, 4, 6}
	mine[3] = []int{2, 0, 7, 3, 2, 2}
	mine[4] = []int{3, 1, 5, 9, 2, 4}
	mine[5] = []int{2, 7, 0, 8, 5, 1}

	res := goldMineMax(mine)

	log.Print(res)

}

func goldMineMax(mine [][]int) int {
	row := len(mine)
	col := len(mine[0])

	maxGoldArr := make([][]int, row)

	for i := 0; i < row; i++ {
		maxGoldArr[i] = make([]int, col)

		for j := 0; j < col; j++ {
			maxGoldArr[i][j] = mine[i][j]
		}
	}

	x := []int{1, 1, 1}
	y := []int{-1, 0, 1}

	for j := col - 2; j >= 0; j-- {
		for i := 0; i < row; i++ {
			var maxVal int
			for k := 0; k < len(x); k++ {
				if i+y[k] < row && i+y[k] >= 0 {
					if maxVal < maxGoldArr[i+y[k]][j+x[k]] {
						maxVal = maxGoldArr[i+y[k]][j+x[k]]
					}
				}
			}

			maxGoldArr[i][j] += maxVal
		}
	}

	var res int
	for i := 0; i < row; i++ {
		if res < maxGoldArr[i][0] {
			res = maxGoldArr[i][0]
		}
	}

	return res
}
