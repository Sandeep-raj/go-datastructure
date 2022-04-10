package recursion

import (
	"fmt"
	"log"
)

/*
Coin Change - Combinations - 2

1. You are given a number n, representing the count of coins.
2. You are given n numbers, representing the denominations of n coins.
3. You are given a number "amt".
4. You are required to calculate and print the combinations of the n coins (same coin can be used
     again any number of times) using which the amount "amt" can be paid.

Sample Input
5
2
3
5
6
7
12

Sample Output
2-2-2-2-2-2-.
2-2-2-3-3-.
2-2-2-6-.
2-2-3-5-.
2-3-7-.
2-5-5-.
3-3-3-3-.
3-3-6-.
5-7-.
6-6-.
*/

func TestCoinChangeCombinations2() {
	inpCoin := []int{2, 3, 5, 6, 7}
	amt := 12

	coinChangeCombinations2(0, inpCoin, amt, "")
}

func coinChangeCombinations2(idx int, inpCoin []int, amt int, res string) {
	if idx == len(inpCoin) {
		if amt == 0 {
			log.Print(res)
		}
		return
	}

	for i := amt / inpCoin[idx]; i >= 1; i-- {
		tempStr := ""
		for j := 0; j < i; j++ {
			tempStr = fmt.Sprintf("%s-%d", tempStr, inpCoin[idx])
		}

		coinChangeCombinations2(idx+1, inpCoin, amt-inpCoin[idx]*i, fmt.Sprintf("%s%s", res, tempStr))
	}

	coinChangeCombinations2(idx+1, inpCoin, amt, res)
}
