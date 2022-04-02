package recursion

import (
	"fmt"
	"log"
)

/*
Abbreviation Using Backtracking

1. You are given a word.
2. You have to generate all abbrevations of that word.

Sample Input

pep

Sample Output
pep
pe1
p1p
p2
1ep
1e1
2p
3
*/

func abbrevations(inpStr string, currStr string, currVal int, res *[]string) {
	var tempStr string

	if currVal == 0 {
		tempStr = currStr
	} else {
		tempStr = fmt.Sprintf("%s%d", currStr, currVal)
	}

	if inpStr == "" {
		*res = append(*res, tempStr)
		return
	}

	abbrevations(inpStr[1:], currStr, currVal+1, res)

	abbrevations(inpStr[1:], fmt.Sprintf("%s%s", tempStr, string(inpStr[0])), 0, res)
}

func TestAbbreveation() {
	res := make([]string, 0)
	abbrevations("sandeep", "", 0, &res)

	for _, s := range res {
		log.Print(s)
	}
}
