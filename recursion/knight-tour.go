package recursion

import "log"

/*
Knights Tour

1. You are given a number n, the size of a chess board.
2. You are given a row and a column, as a starting point for a knight piece.
3. You are required to generate the all moves of a knight starting in (row, col) such that knight visits
     all cells of the board exactly once.
4. Complete the body of printKnightsTour function - without changing signature - to calculate and
     print all configurations of the chess board representing the route
     of knight through the chess board. Use sample input and output to get more idea.
*/

var rowAction = []int{-2, -1, 1, 2, 2, 1, -1, -2}
var colAction = []int{1, 2, 2, 1, -1, -2, -2, -1}

func knightTour(n int, r int, c int, inpArr [][]int, res *[][][]int) {
	if n == 0 {
		*res = append(*res, inpArr)
		return
	}

	for i := 0; i < 8; i++ {
		newr := r + rowAction[i]
		newc := c + colAction[i]
		if newr >= 0 && newr < len(inpArr) && newc >= 0 && newc < len(inpArr) && inpArr[newr][newc] == 0 {
			tempArr := DeepCopy(inpArr)
			tempArr[newr][newc] = n
			knightTour(n-1, newr, newc, tempArr, res)
		}
	}
}

func TestKnightTour() {
	n := 5
	r := 2
	c := 0

	inpArr := make([][]int, n)
	for i := 0; i < n; i++ {
		inpArr[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == r && j == c {
				inpArr[i][j] = n * n
			} else {
				inpArr[i][j] = 0
			}
		}
	}

	res := make([][][]int, 0)

	knightTour(n*n-1, r, c, inpArr, &res)

	for _, r := range res {
		log.Printf("%+v", r)
	}
}
