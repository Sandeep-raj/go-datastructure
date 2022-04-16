package stacksandqueues

import "log"

/*
Next Greater Element To The Left

1. You are given a number n, representing the size of array a.
2. You are given n numbers, representing elements of array a.
3. You are required to find "next greater element on the left" for all elements of array
4. Input and output is handled for you.

"Next greater element on the left" of an element x is defined as the first element to left of x having value greater than x.
Note -> If an element does not have any element on it's left side greater than it, consider -1 as it's "next greater element on left"
e.g.
for the array [2 5 9 3 1 12 6 8 7]
Next greater for 2 is -1
Next greater for 5 is -1
Next greater for 9 is -1
Next greater for 3 is 9
Next greater for 1 is 3
Next greater for 12 is -1
Next greater for 6 is 12
Next greater for 8 is 12
Next greater for 7 is 8

Sample Input
5
5
3
8
-2
7

Sample Output
-1
5
-1
8
8
*/

func TestNGEL() {
	//inpArr := []int{2, 5, 9, 3, 1, 12, 6, 8, 7}
	inpArr := []int{5, 3, 8, -2, 7}
	ngel(inpArr)
}

func ngel(inpArr []int) {
	stack := make([]int, 0)
	res := make([]int, len(inpArr))

	for i := 0; i < len(inpArr); i++ {
		for len(stack) > 0 && stack[len(stack)-1] < inpArr[i] {
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
