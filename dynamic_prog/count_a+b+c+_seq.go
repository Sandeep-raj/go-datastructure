package dynamicprog

import (
	"log"
	"strings"
)

/*
Count A+b+c+ Subsequences

1. You are given a string str.
2. You are required to calculate and print the count of subsequences of the nature a+b+c+.
For abbc -> there are 3 subsequences. abc, abc, abbc
For abcabc -> there are 7 subsequences. abc, abc, abbc, aabc, abcc, abc, abc.

Sample Input
abcabc

Sample Output
7
*/

func TestApBpCpSeq() {
	log.Print(apbpcpSeq("abcabc"))
}

func apbpcpSeq(str string) int {
	charArr := strings.Split(str, "")

	apcount := 0
	apbpcount := 0
	apbpcpcount := 0

	for i := 0; i < len(charArr); i++ {
		if charArr[i] == "a" {
			apcount = 2*apcount + 1
		} else if charArr[i] == "b" {
			apbpcount = 2*apbpcount + apcount
		} else if charArr[i] == "c" {
			apbpcpcount = 2*apbpcpcount + apbpcount
		}
	}

	return apbpcpcount
}
