package recursion

import (
	"fmt"
	"log"
)

/*
Words - K Selection - 3

1. You are given a word (may have one character repeat more than once).
2. You are given an integer k.
2. You are required to generate and print all ways you can select k characters out of the word.

Sample Input
aabbbccdde
3

Sample Output
aab
aac
aad
aae
abb
abc
abd
abe
acc
acd
ace
add
ade
bbb
bbc
bbd
bbe
bcc
bcd
bce
bdd
bde
ccd
cce
cdd
cde
dde
*/

func TestWordsKSelection3() {
	inpStr := "aabbbccdde"
	k := 3

	uniqChars := make([]string, 0)
	fmap := make(map[string]int)
	for i := 0; i < len(inpStr); i++ {
		if fmap[string(inpStr[i])] == 0 {
			uniqChars = append(uniqChars, string(inpStr[i]))
		}
		fmap[string(inpStr[i])]++
	}

	wordsKSelection3(0, uniqChars, fmap, k, "")

}

func wordsKSelection3(idx int, unipChars []string, fmap map[string]int, k int, res string) {

	if idx == len(unipChars) {
		if k == 0 {
			log.Print(res)
		}
		return
	}

	currChar := unipChars[idx]

	for i := fmap[currChar]; i >= 0; i-- {
		if i <= k {
			tempStr := ""
			for j := 0; j < i; j++ {
				tempStr = fmt.Sprintf("%s%s", tempStr, currChar)
			}

			wordsKSelection3(idx+1, unipChars, fmap, k-i, fmt.Sprintf("%s%s", res, tempStr))
		}
	}
}
