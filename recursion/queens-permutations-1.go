package recursion

import (
	"fmt"
	"log"
)

/*
Queens Permutations - 2d As 2d - Queen Chooses

1. You are given a number n, representing the size of a n * n chess board.
2. You are required to calculate and print the permutations in which n queens can be placed on the
     n * n chess-board.

Sample Input
2

Sample Output
q1	q2
-	-

q1	-
q2	-

q1	-
-	q2

q2	q1
-	-

-	q1
q2	-

-	q1
-	q2

q2	-
q1	-

-	q2
q1	-

-	-
q1	q2

q2	-
-	q1

-	q2
-	q1

-	-
q2	q1
*/

func TestQueenPermutation1() {
	n := 2
	chessBoard := make([][]string, n)

	for i := 0; i < n; i++ {
		chessBoard[i] = make([]string, n)
		for j := 0; j < n; j++ {
			chessBoard[i][j] = "-"
		}
	}

	queensPermutation1(n, 1, chessBoard)
}

func queensPermutation1(n int, k int, chessBoard [][]string) {
	if k == n+1 {
		log.Printf("%+v", chessBoard)
		return
	}

	for i := 0; i < n*n; i++ {
		row := int(i / n)
		col := int(i % n)

		if chessBoard[row][col] == "-" {
			chessBoard[row][col] = fmt.Sprintf("q%d", k)
			queensPermutation1(n, k+1, chessBoard)
			chessBoard[row][col] = "-"
		}
	}
}
