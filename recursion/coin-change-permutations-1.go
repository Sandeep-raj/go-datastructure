package recursion

import (
	"fmt"
	"log"
)

/*
Coin Change - Permutations - 1

1. You are given a number n, representing the count of coins.
2. You are given n numbers, representing the denominations of n coins.
3. You are given a number "amt".
4. You are required to calculate and print the permutations of the n coins (non-duplicate) using which the amount "amt" can be paid.

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
2-7-3-.
3-2-7-.
3-7-2-.
5-7-.
7-2-3-.
7-3-2-.
7-5-.
*/

func TestCoinChangePermutation1() {
	coins := []int{2, 3, 5, 6, 7}
	amt := 12

	coinChangePermutation1(coins, amt, "")
}

func coinChangePermutation1(coin []int, amt int, res string) {
	if amt == 0 {
		log.Print(res)
		return
	}

	for i := 0; i < len(coin); i++ {
		if amt-coin[i] >= 0 {
			tempcoin := deepClone(coin)
			tempcoin = append(tempcoin[0:i], tempcoin[i+1:]...)
			coinChangePermutation1(tempcoin, amt-coin[i], fmt.Sprintf("%s-%d", res, coin[i]))
		}
	}
}
