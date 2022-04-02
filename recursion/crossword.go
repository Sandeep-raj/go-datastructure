package recursion

import (
	"fmt"
	"log"
)

/*
Crossword Puzzle
Sample Input

+-++++++++
+-++++++++
+-++++++++
+-----++++
+-+++-++++
+-+++-++++
+++++-++++
++------++
+++++-++++
+++++-++++
4
LONDON
DELHI
ICELAND
ANKARA

Sample Output
+L++++++++
+O++++++++
+N++++++++
+DELHI++++
+O+++C++++
+N+++E++++
+++++L++++
++ANKARA++
+++++N++++
+++++D++++
*/

func TestCrossword() {
	inpArr := make([][]rune, 10)
	inpArr[0] = []rune{'+', '-', '+', '+', '+', '+', '+', '+', '+', '+'}
	inpArr[1] = []rune{'+', '-', '+', '+', '+', '+', '+', '+', '+', '+'}
	inpArr[2] = []rune{'+', '-', '+', '+', '+', '+', '+', '+', '+', '+'}
	inpArr[3] = []rune{'+', '-', '-', '-', '-', '-', '+', '+', '+', '+'}
	inpArr[4] = []rune{'+', '-', '+', '+', '+', '-', '+', '+', '+', '+'}
	inpArr[5] = []rune{'+', '-', '+', '+', '+', '-', '+', '+', '+', '+'}
	inpArr[6] = []rune{'+', '+', '+', '+', '+', '-', '+', '+', '+', '+'}
	inpArr[7] = []rune{'+', '+', '-', '-', '-', '-', '-', '-', '+', '+'}
	inpArr[8] = []rune{'+', '+', '+', '+', '+', '-', '+', '+', '+', '+'}
	inpArr[9] = []rune{'+', '+', '+', '+', '+', '-', '+', '+', '+', '+'}

	strArr := []string{"LONDON", "DELHI", "ANKARA", "ICELAND"}
	slotArr := make([][]int, 0)

	for i := 0; i < len(inpArr); i++ {
		for j := 0; j < len(inpArr); j++ {
			if inpArr[i][j] == '-' {
				//check for horizontal
				if (j-1 >= 0 && inpArr[i][j-1] == '+') || j == 0 {
					// get length
					tempCol := j
					l := 0
					for tempCol < len(inpArr) && inpArr[i][tempCol] != '+' {
						l++
						tempCol++
					}
					if l > 1 {
						slotArr = append(slotArr, []int{i, j, 'h', l})
					}
				}
				//check for vertical
				if (i-1 >= 0 && inpArr[i-1][j] == '+') || i == 0 {
					// get length
					tempRow := i
					l := 0
					for tempRow < len(inpArr) && inpArr[tempRow][j] != '+' {
						l++
						tempRow++
					}
					if l > 1 {
						slotArr = append(slotArr, []int{i, j, 'v', l})
					}
				}
			}
		}
	}

	crossword(strArr, slotArr, inpArr)
}

func crossword(strArr []string, slotArr [][]int, inpArr [][]rune) {
	if len(strArr) == 0 {
		for i := 0; i < len(inpArr); i++ {
			for j := 0; j < len(inpArr); j++ {
				fmt.Print(string(inpArr[i][j]))
			}
			log.Println()
		}
		return
	}

	currStr := strArr[0]

	for idx, slot := range slotArr {
		if len(currStr) == slot[3] && slot[2] == int('h') {
			tempArr := deepCopyRune(inpArr)
			isValid := true
			for i := 0; i < slot[3]; i++ {
				if tempArr[slot[0]][slot[1]+i] == '-' || tempArr[slot[0]][slot[1]+i] == rune(currStr[i]) {
					tempArr[slot[0]][slot[1]+i] = rune(currStr[i])
				} else {
					isValid = false
					break
				}
			}
			if isValid {
				crossword(strArr[1:], append(slotArr[0:idx], slotArr[idx+1:]...), tempArr)
			}
		} else if len(currStr) == slot[3] && slot[2] == int('v') {
			tempArr := deepCopyRune(inpArr)
			isValid := true
			for i := 0; i < slot[3]; i++ {
				if tempArr[slot[0]+i][slot[1]] == '-' || tempArr[slot[0]+i][slot[1]] == rune(currStr[i]) {
					tempArr[slot[0]+i][slot[1]] = rune(currStr[i])
				} else {
					isValid = false
					break
				}
			}
			if isValid {
				crossword(strArr[1:], append(slotArr[0:idx], slotArr[idx+1:]...), tempArr)
			}
		}
	}
}

func deepCopyRune(src [][]rune) [][]rune {
	dst := make([][]rune, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = make([]rune, len(src))
		for j := 0; j < len(src[i]); j++ {
			dst[i][j] = src[i][j]
		}
	}

	return dst
}
