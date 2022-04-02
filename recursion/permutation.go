package recursion

import "log"

/*
Complete the body of printPermutations function - without changing signature - to calculate and print all permutations of str.
Use sample input and output to take idea about permutations.

Sample Input

abc

Sample Output
abc
acb
bac
bca
cab
cba
*/

func getPermutation(str string, n int, currStr string, resArr *[]string) {
	if len(currStr) == n {
		*resArr = append(*resArr, currStr)
	}

	for i := 0; i < len(str); i++ {
		getPermutation(str[:i]+str[i+1:], n, currStr+string(str[i]), resArr)
	}
}

func TestGetPermutation() {
	res := make([]string, 0)
	getPermutation("abc", 3, "", &res)
	for _, s := range res {
		log.Print(s)
	}
}
