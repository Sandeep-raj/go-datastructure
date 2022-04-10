package recursion

import (
	"fmt"
	"log"
)

/*
Coin Change - Combinations - 1

1. You are given a number n, representing the count of coins.
2. You are given n numbers, representing the denominations of n coins.
3. You are given a number "amt".
4. You are required to calculate and print the combinations of the n coins (non-duplicate) using
     which the amount "amt" can be paid.

Sample Input
5
2
3
5
6
7
12

Sample Output
2-3-7-.
5-7-.
*/

func TestCoinChangeCombinations1() {
	inpCoin := []int{2, 3, 5, 6, 7}
	amt := 12

	coinChangeCombinations1(0, inpCoin, amt, "")
}

func coinChangeCombinations1(idx int, inpCoin []int, amt int, res string) {
	if idx == len(inpCoin) {
		if amt == 0 {
			log.Print(res)
		}
		return
	}

	if amt-inpCoin[idx] >= 0 {
		coinChangeCombinations1(idx+1, inpCoin, amt-inpCoin[idx], fmt.Sprintf("%s%d-", res, inpCoin[idx]))
	}
	coinChangeCombinations1(idx+1, inpCoin, amt, res)
}
