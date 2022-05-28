package dynamicprog

import (
	"log"
	"strconv"
	"strings"
)

/*
Count Encodings

1. You are given a string str of digits. (will never start with a 0)
2. You are required to encode the str as per following rules
    1 -> a
    2 -> b
    3 -> c
    ..
    25 -> y
    26 -> z
3. You are required to calculate and print the count of encodings for the string str.

     For 123 -> there are 3 encodings. abc, aw, lc
     For 993 -> there is 1 encoding. iic
     For 013 -> This is an invalid input. A string starting with 0 will not be passed.
     For 103 -> there is 1 encoding. jc
     For 303 -> there are 0 encodings. But such a string maybe passed. In this case
     print 0.

Sample Input
123

Sample Output
3
*/

func TestConutEncoding() {
	log.Print(countEncoding("12012"))
}

func countEncoding(inp string) int {
	inpArr := strings.Split(inp, "")

	lastVal, _ := strconv.Atoi(inpArr[0])

	if lastVal == 0 {
		return 0
	}

	singleLastCount := 1
	combineLastCount := 0

	for i := 1; i < len(inpArr); i++ {
		var tempSingle int
		var tempCombine int
		currVal, _ := strconv.Atoi(inpArr[i])

		combineVal := 10*lastVal + currVal

		if currVal == 0 {
			tempSingle = 0
			if combineVal > 0 && combineVal <= 26 {
				tempCombine = singleLastCount
			} else {
				tempCombine = 0
			}
		} else {
			if combineVal > 0 && combineVal <= 26 {
				tempSingle = singleLastCount + combineLastCount
				tempCombine = singleLastCount
			} else {
				tempSingle = singleLastCount + combineLastCount
				tempCombine = 0
			}
		}

		lastVal = currVal

		singleLastCount = tempSingle
		combineLastCount = tempCombine

	}

	return singleLastCount + combineLastCount
}
