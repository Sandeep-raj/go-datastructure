package stacksandqueues

import "log"

/*
Next Smaller Element To The Right

1. You are given a number n, representing the size of array a.
2. You are given n numbers, representing elements of array a.
3. You are required to "next smaller element on the right" for all elements of array
4. Input and output is handled for you.

"Next smaller element on the right" of an element x is defined as the first element to right of x having value smaller than x.
Note -> If an element does not have any element on it's right side smaller than it, consider -1 as it's "next smaller element on right"
e.g.
for the array [2 5 9 3 1 12 6 8 7]
Next smaller for 2 is 1
Next smaller for 5 is 3
Next smaller for 9 is 3
Next smaller for 3 is 1
Next smaller for 1 is -1
Next smaller for 12 is 6
Next smaller for 6 is -1
Next smaller for 8 is 7
Next smaller for 7 is -1

Sample Input
5
5
3
8
-2
7

Sample Output
3
-2
-2
-1
-1
*/

func TestNSER() {
	inpArr := []int{5, 3, 8, -2, 7}
	nser(inpArr)
}

func nser(inpArr []int) {
	stack := make([]int, 0)
	res := make([]int, len(inpArr))

	for i := len(inpArr) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] > inpArr[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			res[i] = -1
		} else {
			res[i] = stack[len(stack)-1]
		}

		stack = append(stack, inpArr[i])
	}

	log.Printf("%+v", res)
}
