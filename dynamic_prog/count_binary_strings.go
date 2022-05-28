package dynamicprog

import (
	"fmt"
	"log"
)

/*
Count Binary Strings

1. You are given a number n.
2. You are required to print the number of binary strings of length n with no consecutive 0's.

Sample Input
6
Sample Output
21
*/

func TestCountBinStr() {
	// dp := make(map[string]int)
	// log.Print(countBinaryStrings(6, 1, &dp))
	log.Print(countBinaryStringsTopDown(6))
}

func countBinaryStrings(n int, last int, dp *map[string]int) int {
	if n == 0 {
		return 1
	}

	if (*dp)[fmt.Sprintf("%d_%d", n, last)] != 0 {
		return (*dp)[fmt.Sprintf("%d_%d", n, last)]
	}

	log.Print(n, last)

	val := 0
	if last == 0 {
		val += countBinaryStrings(n-1, 1, dp)
	} else {
		val += countBinaryStrings(n-1, 0, dp)
		val += countBinaryStrings(n-1, 1, dp)
	}

	(*dp)[fmt.Sprintf("%d_%d", n, last)] = val

	return val
}

func countBinaryStringsTopDown(n int) int {
	oczeros := 1
	ocones := 1

	for i := 2; i <= n; i++ {
		ncones := oczeros + ocones
		nczeros := ocones

		oczeros = nczeros
		ocones = ncones
	}

	return ocones + oczeros
}
