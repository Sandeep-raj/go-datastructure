package recursion

import (
	"fmt"
	"log"
)

/*
Nqueens Combinations - 2d As 1d - Queen Chooses

1. You are given a number n, representing the size of a n * n chess board.
2. You are required to calculate and print the combinations in which n queens can be placed on the
     n * n chess-board.
3. No queen shall be able to kill another.

Sample Input
4

Sample Output
-	q	-	-
-	-	-	q
q	-	-	-
-	-	q	-

-	-	q	-
q	-	-	-
-	-	-	q
-	q	-	-
*/

func TestNQueenCombination1() {
	n := 4
	chessBoard := make([][]string, n)

	for i := 0; i < n; i++ {
		chessBoard[i] = make([]string, n)
		for j := 0; j < n; j++ {
			chessBoard[i][j] = "-"
		}
	}

	cellMap := make(map[string]int)

	nQueenCombination1(0, n, n, chessBoard, cellMap)
}

func nQueenCombination1(idx int, k int, n int, chessBoard [][]string, cellMap map[string]int) {
	if idx == n*n {
		if k == 0 {
			log.Printf("%+v", chessBoard)
		}
		return
	}

	row := int(idx / n)
	col := int(idx % n)

	if chessBoard[row][col] == "-" && k > 0 &&
		cellMap[fmt.Sprintf("r-%d", row)] == 0 &&
		cellMap[fmt.Sprintf("c-%d", col)] == 0 &&
		cellMap[fmt.Sprintf("pd-%d", row+col)] == 0 &&
		cellMap[fmt.Sprintf("nd-%d", n+(row-col))] == 0 {
		chessBoard[row][col] = "q"
		cellMap[fmt.Sprintf("r-%d", row)] = 1
		cellMap[fmt.Sprintf("c-%d", col)] = 1
		cellMap[fmt.Sprintf("pd-%d", row+col)] = 1
		cellMap[fmt.Sprintf("nd-%d", n+(row-col))] = 1
		nQueenCombination1(idx+1, k-1, n, chessBoard, cellMap)
		chessBoard[row][col] = "-"
		cellMap[fmt.Sprintf("r-%d", row)] = 0
		cellMap[fmt.Sprintf("c-%d", col)] = 0
		cellMap[fmt.Sprintf("pd-%d", row+col)] = 0
		cellMap[fmt.Sprintf("nd-%d", n+(row-col))] = 0
	}

	nQueenCombination1(idx+1, k, n, chessBoard, cellMap)
}
