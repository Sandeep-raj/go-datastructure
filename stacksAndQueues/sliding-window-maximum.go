package stacksandqueues

import (
	"log"
)

/*
Sliding Window Maximum

1. You are given a number n, representing the size of array a.
2. You are given n numbers, representing the elements of array a.
3. You are given a number k, representing the size of window.
4. You are required to find and print the maximum element in every window of size k.

e.g.
for the array [2 9 3 8 1 7 12 6 14 4 32 0 7 19 8 12 6] and k = 4,
the answer is [9 9 8 12 12 14 14 32 32 32 32 19 19 19]
*/

func TestSWM() {
	inpArr := []int{2, 9, 3, 8, 1, 7, 12, 6, 14, 4, 32, 0, 7, 19, 8, 12, 6}
	k := 4

	swm(inpArr, k)
}

func swm(inpArr []int, k int) {
	// nger
	nger := make([]int, len(inpArr))
	stack := make([]int, 0)
	for i := len(inpArr) - 1; i >= 0; i-- {
		for len(stack) > 0 && inpArr[stack[len(stack)-1]] < inpArr[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			nger[i] = -1
		} else {
			nger[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	// sliding window
	swmax := make([]int, len(inpArr)-k+1)
	for i := 0; i <= len(inpArr)-k; i++ {
		tempIndex := i
		for nger[tempIndex] != -1 && nger[tempIndex] < i+k {
			tempIndex = nger[tempIndex]
		}

		swmax[i] = inpArr[tempIndex]
	}

	log.Print(swmax)

}
