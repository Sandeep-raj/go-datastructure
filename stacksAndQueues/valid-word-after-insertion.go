package stacksandqueues

import "log"

/*
Check If Word Is Valid After Insertion

1. You are given a string s determine if it is valid or not.
2. A valid string is a string that can be created by inserting abc at any index any number of times.

Example:
aabcbc can be created
"" -> "abc" -> "aabcbc"

while it is impossible to create abccba.

Sample Input
aabcbc

Sample Output
true
*/

func TestValidWord() {
	inpStr := "aabcbaabcbabaababccbcccc"
	seq := "abc"

	validWord(inpStr, seq)
}

func validWord(inpStr string, seq string) {
	stack := make([]string, 0)

	for i := 0; i < len(inpStr); i++ {
		currChar := string(inpStr[i])

		stack = append(stack, currChar)
		if currChar == string(seq[len(seq)-1]) {
			j := len(seq) - 1
			for len(stack) > 0 && j >= 0 && stack[len(stack)-1] == string(seq[j]) {
				stack = stack[:len(stack)-1]
				j--
			}

			if j != -1 {
				log.Print("false")
				return
			}
		}
	}

	if len(stack) == 0 {
		log.Print("true")
	} else {
		log.Print("false")
	}
}
