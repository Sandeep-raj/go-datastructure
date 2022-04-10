package recursion

import (
	"fmt"
	"log"
)

/*
Nqueens Permutations - 2d As 1d - Queen Chooses

1. You are given a number n, representing the size of a n * n chess board.
2. You are required to calculate and print the permutations in which n queens (distinct) can be
     placed on the n * n chess-board.
3. No queen shall be able to kill another.

Sample Input
4

Sample Output
-	q1	-	-
-	-	-	q2
q3	-	-	-
-	-	q4	-

-	q1	-	-
-	-	-	q2
q4	-	-	-
-	-	q3	-

-	q1	-	-
-	-	-	q3
q2	-	-	-
-	-	q4	-

-	q1	-	-
-	-	-	q4
q2	-	-	-
-	-	q3	-

-	q1	-	-
-	-	-	q3
q4	-	-	-
-	-	q2	-

-	q1	-	-
-	-	-	q4
q3	-	-	-
-	-	q2	-

-	-	q1	-
q2	-	-	-
-	-	-	q3
-	q4	-	-

-	-	q1	-
q2	-	-	-
-	-	-	q4
-	q3	-	-

-	-	q1	-
q3	-	-	-
-	-	-	q2
-	q4	-	-

-	-	q1	-
q4	-	-	-
-	-	-	q2
-	q3	-	-

-	-	q1	-
q3	-	-	-
-	-	-	q4
-	q2	-	-

-	-	q1	-
q4	-	-	-
-	-	-	q3
-	q2	-	-

-	-	q2	-
q1	-	-	-
-	-	-	q3
-	q4	-	-

-	-	q2	-
q1	-	-	-
-	-	-	q4
-	q3	-	-

-	-	q3	-
q1	-	-	-
-	-	-	q2
-	q4	-	-

-	-	q4	-
q1	-	-	-
-	-	-	q2
-	q3	-	-

-	-	q3	-
q1	-	-	-
-	-	-	q4
-	q2	-	-

-	-	q4	-
q1	-	-	-
-	-	-	q3
-	q2	-	-

-	q2	-	-
-	-	-	q1
q3	-	-	-
-	-	q4	-

-	q2	-	-
-	-	-	q1
q4	-	-	-
-	-	q3	-

-	q3	-	-
-	-	-	q1
q2	-	-	-
-	-	q4	-

-	q4	-	-
-	-	-	q1
q2	-	-	-
-	-	q3	-

-	q3	-	-
-	-	-	q1
q4	-	-	-
-	-	q2	-

-	q4	-	-
-	-	-	q1
q3	-	-	-
-	-	q2	-

-	q2	-	-
-	-	-	q3
q1	-	-	-
-	-	q4	-

-	q2	-	-
-	-	-	q4
q1	-	-	-
-	-	q3	-

-	q3	-	-
-	-	-	q2
q1	-	-	-
-	-	q4	-

-	q4	-	-
-	-	-	q2
q1	-	-	-
-	-	q3	-

-	q3	-	-
-	-	-	q4
q1	-	-	-
-	-	q2	-

-	q4	-	-
-	-	-	q3
q1	-	-	-
-	-	q2	-

-	-	q2	-
q3	-	-	-
-	-	-	q1
-	q4	-	-

-	-	q2	-
q4	-	-	-
-	-	-	q1
-	q3	-	-

-	-	q3	-
q2	-	-	-
-	-	-	q1
-	q4	-	-

-	-	q4	-
q2	-	-	-
-	-	-	q1
-	q3	-	-

-	-	q3	-
q4	-	-	-
-	-	-	q1
-	q2	-	-

-	-	q4	-
q3	-	-	-
-	-	-	q1
-	q2	-	-

-	-	q2	-
q3	-	-	-
-	-	-	q4
-	q1	-	-

-	-	q2	-
q4	-	-	-
-	-	-	q3
-	q1	-	-

-	-	q3	-
q2	-	-	-
-	-	-	q4
-	q1	-	-

-	-	q4	-
q2	-	-	-
-	-	-	q3
-	q1	-	-

-	-	q3	-
q4	-	-	-
-	-	-	q2
-	q1	-	-

-	-	q4	-
q3	-	-	-
-	-	-	q2
-	q1	-	-

-	q2	-	-
-	-	-	q3
q4	-	-	-
-	-	q1	-

-	q2	-	-
-	-	-	q4
q3	-	-	-
-	-	q1	-

-	q3	-	-
-	-	-	q2
q4	-	-	-
-	-	q1	-

-	q4	-	-
-	-	-	q2
q3	-	-	-
-	-	q1	-

-	q3	-	-
-	-	-	q4
q2	-	-	-
-	-	q1	-

-	q4	-	-
-	-	-	q3
q2	-	-	-
-	-	q1	-
*/

func TestNQueenPermutation1() {
	n := 4
	chessBoard := make([][]string, n)

	for i := 0; i < n; i++ {
		chessBoard[i] = make([]string, n)
		for j := 0; j < n; j++ {
			chessBoard[i][j] = "-"
		}
	}

	cellMap := make(map[string]int)

	nQueenPermutation1(0, n, chessBoard, cellMap)
}

func nQueenPermutation1(k int, n int, chessboard [][]string, cellMap map[string]int) {
	if k == n {
		log.Printf("%+v", chessboard)
		return
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if chessboard[i][j] == "-" && isQueenSafe(i, j, n, cellMap) {
				chessboard[i][j] = fmt.Sprintf("q%d", k+1)
				cellMap[fmt.Sprintf("r-%d", i)] = 1
				cellMap[fmt.Sprintf("c-%d", j)] = 1
				cellMap[fmt.Sprintf("pd-%d", i+j)] = 1
				cellMap[fmt.Sprintf("nd-%d", n+i-j)] = 1
				nQueenPermutation1(k+1, n, chessboard, cellMap)
				chessboard[i][j] = "-"
				cellMap[fmt.Sprintf("r-%d", i)] = 0
				cellMap[fmt.Sprintf("c-%d", j)] = 0
				cellMap[fmt.Sprintf("pd-%d", i+j)] = 0
				cellMap[fmt.Sprintf("nd-%d", n+i-j)] = 0
			}
		}
	}
}

func isQueenSafe(row int, col int, n int, cellMap map[string]int) bool {

	if cellMap[fmt.Sprintf("r-%d", row)] == 0 &&
		cellMap[fmt.Sprintf("c-%d", col)] == 0 &&
		cellMap[fmt.Sprintf("pd-%d", row+col)] == 0 &&
		cellMap[fmt.Sprintf("nd-%d", n+row-col)] == 0 {
		return true
	}

	return false
}
