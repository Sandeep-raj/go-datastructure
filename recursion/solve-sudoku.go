package recursion

import (
	"fmt"
	"log"
)

/*
Solve Sudoku

1. You are give a partially filled 9*9 2-D array(arr) which represents an incomplete sudoku state.
2. You are required to assign the digits from 1 to 9 to the empty cells following some rules.
Rule 1 -> Digits from 1-9 must occur exactly once in each row.
Rule 2 -> Digits from 1-9 must occur exactly once in each column.
Rule 3 -> Digits from 1-9 must occur exactly once in each 3x3 sub-array of the given 2D array.

Sample Input
3 0 6 5 0 8 4 0 0
5 2 0 0 0 0 0 0 0
0 8 7 0 0 0 0 3 1
0 0 3 0 1 0 0 8 0
9 0 0 8 6 3 0 0 5
0 5 0 0 9 0 6 0 0
1 3 0 0 0 0 2 5 0
0 0 0 0 0 0 0 7 4
0 0 5 2 0 6 3 0 0

Sample Output
3 1 6 5 7 8 4 9 2
5 2 9 1 3 4 7 6 8
4 8 7 6 2 9 5 3 1
2 6 3 4 1 5 9 8 7
9 7 4 8 6 3 1 2 5
8 5 1 7 9 2 6 4 3
1 3 8 9 4 7 2 5 6
6 9 2 3 5 1 8 7 4
7 4 5 2 8 6 3 1 9
*/

func TestSolveSudoku() {
	inpArr := make([][]int, 9)
	inpArr[0] = []int{3, 0, 6, 5, 0, 8, 4, 0, 0}
	inpArr[1] = []int{5, 2, 0, 0, 0, 0, 0, 0, 0}
	inpArr[2] = []int{0, 8, 7, 0, 0, 0, 0, 3, 1}
	inpArr[3] = []int{0, 0, 3, 0, 1, 0, 0, 8, 0}
	inpArr[4] = []int{9, 0, 0, 8, 6, 3, 0, 0, 5}
	inpArr[5] = []int{0, 5, 0, 0, 9, 0, 6, 0, 0}
	inpArr[6] = []int{1, 3, 0, 0, 0, 0, 2, 5, 0}
	inpArr[7] = []int{0, 0, 0, 0, 0, 0, 0, 7, 4}
	inpArr[8] = []int{0, 0, 5, 2, 0, 6, 3, 0, 0}

	setMap := make(map[string]string)
	fillBlocks := make([][]int, 0)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if inpArr[i][j] != 0 {
				setMap[fmt.Sprintf("%d-row-%d", inpArr[i][j], i)] = fmt.Sprintf("%d-row-%d", inpArr[i][j], i)
				setMap[fmt.Sprintf("%d-col-%d", inpArr[i][j], j)] = fmt.Sprintf("%d-col-%d", inpArr[i][j], j)
				iset := int(i / 3)
				jset := int(j / 3)
				setMap[fmt.Sprintf("%d-set-%d-%d", inpArr[i][j], iset, jset)] = fmt.Sprintf("%d-set-%d-%d", inpArr[i][j], iset, jset)
			} else {
				fillBlocks = append(fillBlocks, []int{i, j})
			}
		}
	}

	solveSudoku(inpArr, fillBlocks, setMap)

}

func solveSudoku(inpArr [][]int, fblock [][]int, setMap map[string]string) bool {
	if len(fblock) == 0 {
		log.Printf("%+v", inpArr)
		log.Print(isValidSudoku(inpArr))
		return true
	}

	for _, block := range fblock {
		i := block[0]
		j := block[1]
		iset := int(i / 3)
		jset := int(j / 3)
		isValid := false
		for k := 1; k <= 9; k++ {
			if setMap[fmt.Sprintf("%d-row-%d", k, i)] == "" && setMap[fmt.Sprintf("%d-col-%d", k, j)] == "" && setMap[fmt.Sprintf("%d-set-%d-%d", k, iset, jset)] == "" {
				tempArr := DeepCopy(inpArr)
				tempSet := deepCopy(setMap)
				tempArr[i][j] = k
				tempSet[fmt.Sprintf("%d-row-%d", tempArr[i][j], i)] = fmt.Sprintf("%d-row-%d", tempArr[i][j], i)
				tempSet[fmt.Sprintf("%d-col-%d", tempArr[i][j], j)] = fmt.Sprintf("%d-col-%d", tempArr[i][j], j)
				tempSet[fmt.Sprintf("%d-set-%d-%d", tempArr[i][j], iset, jset)] = fmt.Sprintf("%d-set-%d-%d", tempArr[i][j], iset, jset)

				isValid = solveSudoku(tempArr, fblock[1:], tempSet)
			}
		}
		if !isValid {
			return isValid
		}
	}
	return true
}

func deepCopy(src map[string]string) map[string]string {
	dst := make(map[string]string)

	for key, val := range src {
		dst[key] = val
	}

	return dst
}

func isValidSudoku(src [][]int) bool {
	setMap := make(map[string]string)
	for i := 0; i < len(src); i++ {
		for j := 0; j < len(src[0]); j++ {
			iset := int(i / 3)
			jset := int(j / 3)
			if setMap[fmt.Sprintf("%d-row-%d", src[i][j], i)] == "" && setMap[fmt.Sprintf("%d-col-%d", src[i][j], j)] == "" && setMap[fmt.Sprintf("%d-set-%d-%d", src[i][j], iset, jset)] == "" {
				setMap[fmt.Sprintf("%d-row-%d", src[i][j], i)] = fmt.Sprintf("%d-row-%d", src[i][j], i)
				setMap[fmt.Sprintf("%d-col-%d", src[i][j], j)] = fmt.Sprintf("%d-col-%d", src[i][j], j)
				setMap[fmt.Sprintf("%d-set-%d-%d", src[i][j], iset, jset)] = fmt.Sprintf("%d-set-%d-%d", src[i][j], iset, jset)
			} else {
				return false
			}
		}
	}

	return true
}
