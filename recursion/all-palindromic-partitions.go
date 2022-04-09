package recursion

import "log"

/*
All Palindromic Partitions

1. You are given a string of length n.
2. You have to partition the given string in such a way that every partition is a palindrome.

Sample Input
pep

Sample Output
(p) (e) (p)
(pep)
*/

func TestPallindromicPartitions() {
	pallindromicPartitions("abaaba", make([]string, 0))
}

func pallindromicPartitions(str string, res []string) {
	if len(str) == 0 {
		log.Printf("%+v", res)
	}

	for i := 0; i < len(str); i++ {
		if isPallindrome(str[0 : i+1]) {
			res = append(res, str[0:i+1])
			pallindromicPartitions(str[i+1:], res)
			res = res[:len(res)-1]
		}
	}
}

func isPallindrome(str string) bool {
	if len(str) == 1 {
		return true
	}

	slen := len(str)

	for i := 0; i < slen/2; i++ {
		if str[i] != str[slen-1-i] {
			return false
		}
	}

	return true
}
