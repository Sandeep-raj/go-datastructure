package recursion

import "log"

/*
Queens Combinations - 2d As 2d - Box Chooses

1. You are given a number n, representing the size of a n * n chess board.
2. You are required to calculate and print the combinations in which n queens can be placed on the
     n * n chess-board.

Sample Input
2

Sample Output
qq
--

q-
q-

q-
-q

-q
q-

-q
-q

--
qq
*/

func TestQueenCombination1() {
	n := 2
	chessBoard := make([][]string, n)
	for i := 0; i < n; i++ {
		chessBoard[i] = make([]string, n)
		for j := 0; j < n; j++ {
			chessBoard[i][j] = "-"
		}
	}

	queenCombination1(0, n, n, chessBoard)
}

func queenCombination1(idx int, n int, k int, chessBoard [][]string) {

	if idx == n*n {
		if k == 0 {
			log.Printf("%+v", chessBoard)
		}
		return
	}

	row := int(idx / n)
	col := int(idx % n)

	if chessBoard[row][col] == "-" {
		queenCombination1(idx+1, n, k, chessBoard)
		chessBoard[row][col] = "q"
		queenCombination1(idx+1, n, k-1, chessBoard)
		chessBoard[row][col] = "-"
	}
}
