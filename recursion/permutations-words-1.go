package recursion

import (
	"log"
	"strings"
)

/*
Permutations - Words - 1

1. You are given a word (may have one character repeat more than once).
2. You are required to generate and print all arrangements of these characters.

Sample Input
aabb

Sample Output
aabb
abab
abba
baab
baba
bbaa
*/

func TestWordPermutation1() {
	inpstr := "aabb"
	fmap := make(map[string]int)
	for i := 0; i < len(inpstr); i++ {
		fmap[string(inpstr[i])]++
	}
	res := make([]string, len(inpstr))
	wordPermutation1(0, fmap, res)
}

func wordPermutation1(idx int, fmap map[string]int, res []string) {
	if idx == len(res) {
		log.Printf("%s", strings.Join(res, ""))
		return
	}

	for k, v := range fmap {
		if v > 0 {
			res[idx] = k
			fmap[k]--
			wordPermutation1(idx+1, fmap, res)
			fmap[k]++
			res[idx] = ""
		}
	}
}
