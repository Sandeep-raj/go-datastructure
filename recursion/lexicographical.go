package recursion

import (
	"fmt"
	"log"
	"strconv"
)

/*
Lexicographical Numbers
1. You are given a number.
2. You have to print all the numbers from 1 to n in lexicographical order.

Sample Input

14

Sample Output
1
10
11
12
13
14
2
3
4
5
6
7
8
9
*/

func TestLexicographical() {
	res := make([]int, 0)
	lexicographical(14, 1, &res)
	for _, n := range res {
		log.Print(n)
	}
}

func lexicographical(n int, currNum int, res *[]int) {
	if currNum > n {
		return
	}

	*res = append(*res, currNum)

	tempCurrNum, _ := strconv.Atoi(fmt.Sprintf("%d0", currNum))
	if tempCurrNum <= n {
		lexicographical(n, tempCurrNum, res)
	}

	if (currNum+1)%10 != 0 {
		lexicographical(n, currNum+1, res)
	}
}
