package recursion

import "log"

/*
Josephus Problem
1. You are given two numbers N and K.
2. N represents the total number of soldiers standing in a circle having position marked from 0 to N-1.
3. A cruel king wants to execute them but in a different way.
4. He starts executing soldiers from 0th position and proceeds around the circle in clockwise direction.
5. In each step, k-1 soldiers are skipped and the k-th soldier is executed.
6. The elimination proceeds around the circle (which is becoming smaller and smaller as the executed soldiers are removed), until only the last soldier remains, who is given freedom.
7. You have to find the position of that lucky soldier.
*/

func TestJosephus() {
	n := 100
	k := 2

	inpArr := make([]int, n)
	for i := 0; i < n; i++ {
		inpArr[i] = i + 1
	}

	log.Print(josephus(k, inpArr, 0))
}

func josephus(k int, inpArr []int, idx int) int {
	if len(inpArr) == 1 {
		return inpArr[0]
	}

	currIndex := (idx + k - 1) % len(inpArr)
	tempArr := append(inpArr[:currIndex], inpArr[currIndex+1:]...)

	return josephus(k, tempArr, currIndex)
}
