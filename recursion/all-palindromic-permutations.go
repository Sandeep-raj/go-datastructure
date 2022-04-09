package recursion

import (
	"fmt"
	"log"
)

/*
All Palindromic Permutations

1. You are given a string of length n.
2. You have to print all the palindromic permutations of the given string.
3. If no palindromic permutation exists for the given string, print "-1".

Sample Input
aaabb

Sample Output
ababa
baaab
*/

func TestPallindromicPermutation() {
	inpStr := "aaabbcc"

	fmap := make(map[string]int)
	for i := 0; i < len(inpStr); i++ {
		fmap[string(inpStr[i:i+1])]++
	}

	oddChar := ""
	oddsCount := 0
	totCount := 0

	for ch, count := range fmap {
		if count%2 == 1 {
			fmap[ch] = fmap[ch] / 2
			if fmap[ch] == 0 {
				delete(fmap, ch)
			} else {
				totCount += fmap[ch]
			}
			oddChar = ch
			oddsCount++
		} else {
			fmap[ch] = fmap[ch] / 2
			totCount += fmap[ch]
		}
	}

	pallindromicPermutation(totCount, fmap, oddChar, "")

}

func pallindromicPermutation(n int, fmap map[string]int, oddchar string, res string) {
	if n == 0 {
		// print
		resStr := res
		resStr = resStr + oddchar
		resStr = resStr + StringReverse(res)
		log.Print(resStr)
		return
	}

	for k, v := range fmap {
		if v != 0 {
			fmap[k]--
			pallindromicPermutation(n-1, fmap, oddchar, fmt.Sprintf("%s%s", res, k))
			fmap[k]++
		}
	}
}

func StringReverse(str string) string {
	res := ""
	for i := len(str) - 1; i >= 0; i-- {
		res = res + string(str[i])
	}

	return res
}
