package dynamicprog

import "log"

/*
Fibonacci-dp

1. You are given a number n.
2. You are required to print the nth element of fibonnaci sequence.

Note -> Notice precisely how we have defined the fibonnaci sequence
0th element -> 0
1st element -> 1
2nd element -> 1
3rd element -> 2
4th element -> 3
5th element -> 5
6th element -> 8

Example
Sample Input
10

Sample Output
55
*/

func TestFibonacci() {
	dp := make(map[int]int)

	val := fibonacci(10, &dp)
	log.Print(val)
}

func fibonacci(n int, dp *map[int]int) int {
	if n <= 1 {
		return n
	}

	if (*dp)[n] != 0 {
		return (*dp)[n]
	}

	n2 := fibonacci(n-2, dp)
	n1 := fibonacci(n-1, dp)

	(*dp)[n] = n2 + n1
	return n2 + n1
}
