package patternsearchalgorithm

import "log"

func TestKMP() {
	inpStr := "AABAACAADAABAABA"
	pattern := "BAAB"

	index := kmp(inpStr, pattern)
	log.Print(index)

}

func kmp(str string, pat string) int {
	i := 0
	j := 0

	lmps := lmpsProcess(pat)

	for i < len(str) && j < len(pat) {
		if str[i] == pat[j] {
			i++
			j++

			if j == len(pat) {
				return i - j
			}
		} else {
			if j > 0 {
				j = lmps[j-1]
			} else {
				i++
			}
		}
	}

	return -1
}

func lmpsProcess(str string) []int {
	lastPrefix := 0
	curr := 1

	lmps := make([]int, len(str))

	for curr < len(str) {
		if str[lastPrefix] == str[curr] {
			lastPrefix++
			lmps[curr] = lastPrefix
			curr++
		} else {
			if lastPrefix == 0 {
				curr++
			} else {
				lastPrefix = lmps[lastPrefix-1]
			}
		}
	}

	return lmps
}
