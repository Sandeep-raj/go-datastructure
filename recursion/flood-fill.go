package recursion

import "log"

/*
Flood Fill

1. You are given a number n, representing the number of rows.
2. You are given a number m, representing the number of columns.
3. You are given n*m numbers, representing elements of 2d array a. The numbers can be 1 or 0 only.
4. You are standing in the top-left corner and have to reach the bottom-right corner.
Only four moves are allowed 't' (1-step up), 'l' (1-step left), 'd' (1-step down) 'r' (1-step right). You can only move to cells which have 0 value in them. You can't move out of the boundaries or in the cells which have value 1 in them (1 means obstacle)
5. Complete the body of floodfill function - without changing signature - to print all paths that can be used to move from top-left to bottom-right.

Sample Input

3 3
0 0 0
1 0 1
0 0 0

Sample Output
rddr

*/

func floodFill(r int, c int, row int, col int, floodArr [][]int, currPath string, resArr *[]string) {
	floodArr[r][c] = 1
	if r == row-1 && c == col-1 {
		*resArr = append(*resArr, currPath)
		return
	}

	if r-1 > 0 && floodArr[r-1][c] == 0 {
		floodFill(r-1, c, row, col, DeepCopy(floodArr), currPath+"t", resArr)
	}

	if c-1 > 0 && floodArr[r][c-1] == 0 {
		floodFill(r, c-1, row, col, DeepCopy(floodArr), currPath+"l", resArr)
	}

	if r+1 < row && floodArr[r+1][c] == 0 {
		floodFill(r+1, c, row, col, DeepCopy(floodArr), currPath+"d", resArr)
	}

	if c+1 < col && floodArr[r][c+1] == 0 {
		floodFill(r, c+1, row, col, DeepCopy(floodArr), currPath+"r", resArr)
	}
}

func DeepCopy(src [][]int) [][]int {
	dst := make([][]int, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = make([]int, len(src[i]))
		for j := 0; j < len(src[i]); j++ {
			dst[i][j] = src[i][j]
		}
	}
	return dst
}

func TestFloodFill() {
	row := 3
	col := 3
	floodArr := [][]int{
		{0, 0, 0},
		{0, 0, 1},
		{0, 0, 0},
	}

	res := make([]string, 0)

	floodFill(0, 0, row, col, floodArr, "", &res)

	for _, s := range res {
		log.Print(s)
	}
}
