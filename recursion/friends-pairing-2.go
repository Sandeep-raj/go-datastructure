package recursion

import (
	"fmt"
	"log"
)

/*
Friends Pairing - 2

1. You are given an integer n, which represents n friends numbered from 1 to n.
2. Each one can remain single or can pair up with some other friend.
3. You have to print all the configurations in which friends can remain single or can be paired up.

Sample Input

3

Sample Output
1.(1) (2) (3)
2.(1) (2,3)
3.(1,2) (3)
4.(1,3) (2)
*/

func TestFPair() {
	used := make([]bool, 5)
	fPair(1, 5, used, "")
}

func fPair(idx int, n int, used []bool, res string) {
	if idx > n {
		log.Print(res)
		return
	}

	if used[idx-1] {
		fPair(idx+1, n, used, res)
	} else {
		used[idx-1] = true
		//tempUsed := deepCloneBool(used)
		fPair(idx+1, n, used, fmt.Sprintf("%s ( %d ) ", res, idx))

		for j := idx + 1; j <= n; j++ {
			if !used[j-1] {
				used[j-1] = true
				fPair(idx+1, n, used, fmt.Sprintf("%s ( %d , %d )", res, idx, j))
				used[j-1] = false
			}
		}
		used[idx-1] = false
	}
}
