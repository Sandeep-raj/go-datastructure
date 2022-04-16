package stacksandqueues

import "log"

/*
Celebrity Problem

1. You are given a number n, representing the number of people in a party.
2. You are given n strings of n length containing 0's and 1's
3. If there is a '1' in ith row, jth spot, then person i knows about person j.
4. A celebrity is defined as somebody who knows no other person than himself but everybody else knows him.
5. If there is a celebrity print it's index otherwise print "none".

Sample Input
4
0000
1011
1101
1110

Sample Output
0
*/

func TestCelbProblem() {
	n := 4
	celebMatrx := make([][]int, n)
	celebMatrx[0] = []int{0, 1, 0, 0}
	celebMatrx[1] = []int{1, 0, 1, 1}
	celebMatrx[2] = []int{1, 1, 0, 1}
	celebMatrx[3] = []int{1, 1, 1, 0}

	log.Print(celebProblem(n, celebMatrx))
}

func celebProblem(n int, celebMtrx [][]int) int {
	stack := make([]int, n)
	for i := 0; i < n; i++ {
		stack[i] = i
	}

	for len(stack) > 1 {
		index1 := stack[len(stack)-1]
		index2 := stack[len(stack)-2]

		stack = stack[:len(stack)-2]

		if celebMtrx[index1][index2] == 0 {
			stack = append(stack, index1)
		} else {
			stack = append(stack, index2)
		}
	}

	for i := 0; i < n; i++ {
		if celebMtrx[stack[0]][i] == 1 || (i != stack[0] && celebMtrx[i][stack[0]] == 0) {
			return -1
		}
	}

	return stack[0]
}
