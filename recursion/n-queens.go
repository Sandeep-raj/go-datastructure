package recursion

import (
	"fmt"
	"log"
)

func TestNQueen() {
	n := 4
	// res := make([][][]int, 0)
	// inpArr := make([][]int, n)
	// for i := 0; i < n; i++ {
	// 	inpArr[i] = make([]int, n)
	// 	for j := 0; j < n; j++ {
	// 		inpArr[i][j] = 0
	// 	}
	// }

	// nQueen(n, inpArr, &res)

	// log.Print(len(res))
	// for _, resArr := range res {
	// 	log.Printf("%+v", resArr)
	// }

	colArr := make([]bool, n)
	pDiagArr := make([]bool, 2*n-1)
	nDiagArr := make([]bool, 2*n-1)
	res := make([]string, 0)

	nQueenBnB(n, colArr, pDiagArr, nDiagArr, "", &res)
	for _, r := range res {
		log.Print(r)
	}

}

func nQueen(n int, inpArr [][]int, res *[][][]int) {
	if n == 0 {
		*res = append(*res, inpArr)
		return
	}

	for i := 0; i < len(inpArr); i++ {
		if inpArr[n-1][i] == 0 {
			tempArr := DeepCopy(inpArr)
			captureSqaure(tempArr, n-1, i)
			tempArr[n-1][i] = n
			nQueen(n-1, tempArr, res)
		}
	}
}

func captureSqaure(inpArr [][]int, r int, c int) {
	currR := r
	for currR >= 0 {
		inpArr[currR][c] = -1
		currR = currR - 1
	}

	currC := c
	for currC < len(inpArr) {
		inpArr[r][currC] = -1
		currC = currC + 1
	}

	currR = r
	currC = c
	for currR-1 >= 0 && currC+1 < len(inpArr) {
		inpArr[currR-1][currC+1] = -1
		currR--
		currC++
	}

	currR = r
	currC = c
	for currR-1 >= 0 && currC-1 >= 0 {
		inpArr[currR-1][currC-1] = -1
		currR--
		currC--
	}
}

// N Queens - Branch And Bound
/*
1. You are given a number n, the size of a chess board.
2. You are required to place n number of queens in the n * n cells of board such that no queen can
     kill another.
Note - Queens kill at distance in all 8 directions
3. Complete the body of printNQueens function - without changing signature - to calculate and
     print all safe configurations of n-queens
*/

func nQueenBnB(n int, colArr []bool, pDiagArr []bool, nDiagArr []bool, currPlace string, res *[]string) {
	if n == 0 {
		*res = append(*res, currPlace)
		return
	}

	for i := 0; i < len(colArr); i++ {
		if !colArr[i] && !pDiagArr[n-1+i] && !nDiagArr[len(colArr)-1+n-1-i] {
			//currPlace += fmt.Sprintf("<r - %d c - %d>\n", n-1, i)

			tempColArr := deepCloneBool(colArr)
			tempColArr[i] = true
			tempPDiagArr := deepCloneBool(pDiagArr)
			tempPDiagArr[n-1+i] = true
			tempNDiagArr := deepCloneBool(nDiagArr)
			tempNDiagArr[len(colArr)-1+n-1-i] = true

			nQueenBnB(n-1, tempColArr, tempPDiagArr, tempNDiagArr, currPlace+fmt.Sprintf("<r - %d c - %d>\n", n-1, i), res)
		}
	}
}

func deepCloneBool(src []bool) []bool {
	dst := make([]bool, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = src[i]
	}

	return dst
}
