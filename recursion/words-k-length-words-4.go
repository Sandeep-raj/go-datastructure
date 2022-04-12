package recursion

import (
	"fmt"
	"log"
)

/*
Words - K Length Words - 4

1. You are given a word (may have one character repeat more than once).
2. You are given an integer k.
3. You are required to generate and print all k length words by using chars of the word.

Sample Input
aaabb
3

Sample Output
aaa
aab
aba
abb
baa
bab
bba
*/

func TestKLengthWords4() {
	inpStr := "aaabb"
	k := 3

	fmap := make(map[string]int)

	for i := 0; i < len(inpStr); i++ {
		fmap[string(inpStr[i])]++
	}

	kLengthWords4(k, fmap, "")

}

func kLengthWords4(k int, fmap map[string]int, res string) {

	if k == 0 {
		log.Print(res)
		return
	}

	for key, value := range fmap {
		if value > 0 {
			fmap[key]--
			kLengthWords4(k-1, fmap, fmt.Sprintf("%s%s", res, key))
			fmap[key]++
		}
	}
}
