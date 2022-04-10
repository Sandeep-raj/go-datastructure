package recursion

import (
	"log"
	"strings"
)

/*
Combinations - 1

1. You are give a number of boxes (nboxes) and number of identical items (ritems).
2. You are required to place the items in those boxes and print all such configurations possible.

Sample Input
5
3

Sample Output
iii--
ii-i-
ii--i
i-ii-
i-i-i
i--ii
-iii-
-ii-i
-i-ii
--iii
*/

func TestCombinations1() {
	n := 5
	k := 3
	res := make([]string, n)
	for i := 0; i < n; i++ {
		res[i] = "-"
	}

	combinations1(0, k, res)
}

func combinations1(idx int, k int, res []string) {
	if idx == len(res) {
		if k == 0 {
			log.Printf("%s", strings.Join(res, ""))
		}
		return
	}

	if res[idx] == "-" {
		res[idx] = "i"
		combinations1(idx+1, k-1, res)
		res[idx] = "-"
	}
	combinations1(idx+1, k, res)
}
