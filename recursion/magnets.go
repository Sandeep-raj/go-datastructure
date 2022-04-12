package recursion

import (
	"fmt"
	"log"
)

/*
Magnets

1. You are given n number of domino shaped bipolar magnets.
2. You have to place these magnets in M*N following the conditions.
3. Conditions are -
   a. Each box of 1*2 or 2*1 can contain a magnet or can be empty.
   b. Empty box can be represented by X's and magnets are represented by + and
    - sign.
   c. Digits along left and top side of the board represents the number of + in
    corresponding rows and columns.
   d. Digits along right and bottom of the board represents the number of - in
    corresponding rows and columns.
    e. -1 denotes that the corresponding row and column can have any number of
    +  and - signs.
    f. No two adjacent cell can have the same sign.

Sample Input
5 6
LRLRTT
LRLRBB
TTTTLR
BBBBTT
LRLRBB
1 -1 -1 2 1 -1
2 3 -1 -1 -1
-1 -1 -1 1 -1
2 -1 -1 2 -1 3

Sample Output
+ - + - X -
- + - + X +
X X + - + -
X X - + X +
- + X X X -
*/

func TestMagnets() {
	row := 5
	col := 6

	inpCharArr := make([][]string, row)
	inpCharArr[0] = []string{"L", "R", "L", "R", "T", "T"}
	inpCharArr[1] = []string{"L", "R", "L", "R", "B", "B"}
	inpCharArr[2] = []string{"T", "T", "T", "T", "L", "R"}
	inpCharArr[3] = []string{"B", "B", "B", "B", "T", "T"}
	inpCharArr[4] = []string{"L", "R", "L", "R", "B", "B"}

	ptopEdge := []int{1, -1, -1, 2, 1, -1}
	pleftEdge := []int{2, 3, -1, -1, -1}
	nrightEdge := []int{-1, -1, -1, -1, -1}
	nbottomEdge := []int{2, -1, -1, 2, -1, 3}

	box := make([][]string, row)
	for i := 0; i < row; i++ {
		box[i] = make([]string, col)
		for j := 0; j < col; j++ {
			box[i][j] = "X"
		}
	}

	res := magnets(0, inpCharArr, box, ptopEdge, pleftEdge, nrightEdge, nbottomEdge)
	if !res {
		log.Print("No Result")
	}
}

func magnets(idx int, inpCharArr [][]string, box [][]string, top []int, left []int, right []int, bottom []int) bool {

	if idx >= len(inpCharArr)*len(inpCharArr[0]) {
		if isValid(box, top, left, right, bottom) {
			magnetPrint(box)
			return true
		}
		return false
	}

	row := int(idx / len(inpCharArr[0]))
	col := int(idx % len(inpCharArr[0]))

	if inpCharArr[row][col] == "L" {
		if isSafe(row, col, box, top, left, right, bottom, "+", "L") {
			box[row][col] = "+"
			box[row][col+1] = "-"
			res := magnets(idx+2, inpCharArr, box, top, left, right, bottom)
			if res {
				return true
			}
			box[row][col] = "X"
			box[row][col+1] = "X"
		}

		if isSafe(row, col, box, top, left, right, bottom, "-", "L") {
			box[row][col] = "-"
			box[row][col+1] = "+"
			res := magnets(idx+2, inpCharArr, box, top, left, right, bottom)
			if res {
				return true
			}
			box[row][col] = "X"
			box[row][col+1] = "X"
		}
	} else if inpCharArr[row][col] == "T" {
		if isSafe(row, col, box, top, left, right, bottom, "+", "T") {
			box[row][col] = "+"
			box[row+1][col] = "-"
			res := magnets(idx+1, inpCharArr, box, top, left, right, bottom)
			if res {
				return true
			}
			box[row][col] = "X"
			box[row+1][col] = "X"
		}

		if isSafe(row, col, box, top, left, right, bottom, "-", "T") {
			box[row][col] = "-"
			box[row+1][col] = "+"
			res := magnets(idx+1, inpCharArr, box, top, left, right, bottom)
			if res {
				return true
			}
			box[row][col] = "X"
			box[row+1][col] = "X"
		}
	}

	res := magnets(idx+1, inpCharArr, box, top, left, right, bottom)
	if res {
		return true
	} else {
		return false
	}
}

func isSafe(row int, col int, box [][]string, top []int, left []int, right []int, bottom []int, sign string, orientation string) bool {
	if orientation == "L" {
		// place +-
		if sign == "+" && !((col-1 >= 0 && box[row][col-1] == "+") ||
			(row-1 >= 0 && box[row-1][col] == "+") ||
			(col+2 < len(box[0]) && box[row][col+2] == "-") ||
			(row-1 >= 0 && box[row-1][col+1] == "-")) &&
			isContraint(row, col, box, top, left, right, bottom) &&
			isContraint(row, col+1, box, top, left, right, bottom) {
			return true
		}
		// place -+
		if sign == "-" && !((col-1 >= 0 && box[row][col-1] == "-") ||
			(row-1 >= 0 && box[row-1][col] == "-") ||
			(col+2 < len(box[0]) && box[row][col+2] == "+") ||
			(row-1 >= 0 && box[row-1][col+1] == "+")) &&
			isContraint(row, col, box, top, left, right, bottom) &&
			isContraint(row, col+1, box, top, left, right, bottom) {
			return true
		}
	} else {
		if sign == "+" && !((col-1 >= 0 && box[row][col-1] == "+") ||
			(row-1 >= 0 && box[row-1][col] == "+") ||
			(col+1 < len(box[0]) && box[row][col+1] == "+")) &&
			isContraint(row, col, box, top, left, right, bottom) &&
			isContraint(row+1, col, box, top, left, right, bottom) {
			return true
		}
		if sign == "-" && !((col-1 >= 0 && box[row][col-1] == "-") ||
			(row-1 >= 0 && box[row-1][col] == "-") ||
			(col+1 < len(box[0]) && box[row][col+1] == "-")) &&
			isContraint(row, col, box, top, left, right, bottom) &&
			isContraint(row+1, col, box, top, left, right, bottom) {
			return true
		}
	}

	return false
}

func isContraint(row int, col int, box [][]string, top []int, left []int, right []int, bottom []int) bool {
	isconstraint := true

	tCounter := 0
	bCounter := 0
	for i := 0; i < len(box); i++ {
		if box[i][col] == "+" {
			tCounter++
		}
		if box[i][col] == "-" {
			bCounter++
		}
	}

	lCounter := 0
	rCounter := 0
	for i := 0; i < len(box[0]); i++ {
		if box[row][i] == "+" {
			lCounter++
		}
		if box[row][i] == "-" {
			rCounter++
		}
	}

	if (top[col] > -1 && top[col] < tCounter) ||
		(bottom[col] > -1 && bottom[col] < bCounter) ||
		(left[row] > -1 && left[row] < lCounter) ||
		(right[row] > -1 && right[row] < rCounter) {
		isconstraint = false
	}

	return isconstraint
}

func magnetPrint(box [][]string) {
	for i := 0; i < len(box); i++ {
		for j := 0; j < len(box[0]); j++ {
			fmt.Printf("%s  ", box[i][j])
		}
		fmt.Println()
	}
	fmt.Println("------------------------------------------------------")
}

func isValid(box [][]string, top []int, left []int, right []int, bottom []int) bool {
	for i := 0; i < len(box); i++ {
		for j := 0; j < len(box[i]); j++ {
			tCounter := 0
			bCounter := 0
			lCounter := 0
			rCounter := 0
			for k := 0; k < len(box); k++ {
				if box[k][j] == "+" {
					tCounter++
				}
				if box[k][j] == "-" {
					bCounter++
				}
			}

			for k := 0; k < len(box[0]); k++ {
				if box[i][k] == "+" {
					lCounter++
				}
				if box[i][k] == "-" {
					rCounter++
				}
			}

			if (top[j] > -1 && top[j] != tCounter) ||
				(bottom[j] > -1 && bottom[j] != bCounter) ||
				(left[i] > -1 && left[i] != lCounter) ||
				(right[i] > -1 && right[i] != rCounter) {
				return false
			}
		}
	}
	return true
}
