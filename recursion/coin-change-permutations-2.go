package recursion

import (
	"fmt"
	"log"
)

/*
Coin Change - Permutations - 2

1. You are given a number n, representing the count of coins.
2. You are given n numbers, representing the denominations of n coins.
3. You are given a number "amt".
4. You are required to calculate and print the permutations of the n coins (same coin can be used again any number of times) using which the amount "amt" can be paid.

Sample Input

3
2
3
5
7

Sample Output
2-2-3-.
2-3-2-.
2-5-.
3-2-2-.
5-2-.
*/

func TestCoinChangePermutation2() {
	coins := []int{2, 3, 5}
	amt := 7

	coinChangePermutation2(coins, amt, "")
}

func coinChangePermutation2(coin []int, amt int, res string) {
	if amt == 0 {
		log.Print(res)
		return
	}

	for i := 0; i < len(coin); i++ {
		if amt-coin[i] >= 0 {
			coinChangePermutation2(coin, amt-coin[i], fmt.Sprintf("%s-%d", res, coin[i]))
		}
	}
}
