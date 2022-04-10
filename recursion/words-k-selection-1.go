package recursion

import (
	"fmt"
	"log"
)

/*
Words - K Selection - 1

1. You are given a word (may have one character repeat more than once).
2. You are given an integer k.
2. You are required to generate and print all ways you can select k distinct characters out of the
     word.

Sample Input
aabbbccdde
3

Sample Output
abc
abd
abe
acd
ace
ade
bcd
bce
bde
cde
*/

func TestWordsKSelection1() {
	inpStr := "aabbbccdde"
	k := 3

	fmap := make(map[string]int)
	uniqArr := make([]string, 0)

	for i := 0; i < len(inpStr); i++ {
		if fmap[string(inpStr[i])] == 0 {
			fmap[string(inpStr[i])] = 1
			uniqArr = append(uniqArr, string(inpStr[i]))
		}
	}

	wordsKSelection1(0, k, uniqArr, "")
}

func wordsKSelection1(idx int, k int, uniqArr []string, res string) {

	if idx == len(uniqArr) {
		if k == 0 {
			log.Print(res)
		}
		return
	}

	wordsKSelection1(idx+1, k-1, uniqArr, fmt.Sprintf("%s%s", res, uniqArr[idx]))
	wordsKSelection1(idx+1, k, uniqArr, res)
}
