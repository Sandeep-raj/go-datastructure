package stacksandqueues

import "log"

/*
1. You are given a string exp representing an expression.
2. Assume that the expression is balanced i.e. the opening and closing brackets match with each other.
3. But, some of the pair of brackets maybe extra/needless.
4. You are required to print true if you detect extra brackets and false otherwise. e.g.' ((a + b) + (c + d)) -> false (a + b) + ((c + d)) -> true

Sample Input
(a + b) + ((c + d))

Sample Output
true
*/

func TestDuplicateBrackets() {
	log.Print(duplicateBrackets("(a+b)+(c+d)"))
}

func duplicateBrackets(inpStr string) bool {
	stack := make([]string, 0)
	for i := 0; i < len(inpStr); i++ {
		if string(inpStr[i]) != ")" {
			stack = append(stack, string(inpStr[i]))
		} else {
			noOfElements := 0
			for len(stack)-1 >= 0 && stack[len(stack)-1] != "(" {
				stack = stack[:len(stack)-1]
				noOfElements++
			}
			if noOfElements == 0 {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return true
}
