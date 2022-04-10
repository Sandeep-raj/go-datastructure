package recursion

import "log"

/*
Nknights Combinations - 2d As 1d - Knight Chooses

1. You are given a number n, representing the size of a n * n chess board.
2. You are required to calculate and print the combinations in which n knights can be placed on the
     n * n chess-board.
3. No knight shall be able to kill another.

Sample Input
2

Sample Output
k	k
-	-

k	-
k	-

k	-
-	k

-	k
k	-

-	k
-	k

-	-
k	k
*/

func TestNKnightsCombinations1() {
	n := 2

	chessboard := make([][]string, n)

	for i := 0; i < n; i++ {
		chessboard[i] = make([]string, n)
		for j := 0; j < n; j++ {
			chessboard[i][j] = "-"
		}
	}

	knightsCombinations1(0, n, n, chessboard)

}

func knightsCombinations1(idx int, n int, k int, chessboard [][]string) {

	if idx == n*n {
		if k == 0 {
			log.Printf("%+v", chessboard)
		}
		return
	}

	row := int(idx / n)
	col := int(idx % n)

	if isKnightSafe(row, col, n, chessboard) && k > 0 {
		chessboard[row][col] = "k"
		knightsCombinations1(idx+1, n, k-1, chessboard)
		chessboard[row][col] = "-"
	}

	knightsCombinations1(idx+1, n, k, chessboard)

}

func isKnightSafe(row int, col int, n int, chessboard [][]string) bool {
	if (row-2 >= 0 && col-1 >= 0 && chessboard[row-2][col-1] == "k") ||
		(row-1 >= 0 && col-2 >= 0 && chessboard[row-1][col-2] == "k") ||
		(row-2 >= 0 && col+1 < n && chessboard[row-2][col+1] == "k") ||
		(row-1 >= 0 && col+2 < n && chessboard[row-1][col+2] == "k") {
		return false
	}
	return true
}
