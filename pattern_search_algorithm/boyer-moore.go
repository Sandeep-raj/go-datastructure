package patternsearchalgorithm

import (
	"log"
	"math"
)

func BoyerMooreTest() {
	str := "THIS IS A TEST TEXT"
	pat := "TEST"

	idx := boyermoore(str, pat)
	log.Print(idx)
}

func boyermoore(str string, pat string) int {
	n := len(str)
	m := len(pat)

	bmt := badmatchtable(pat)

	idx := 0

	for idx <= n-m {
		stridx := m - 1

		for stridx >= 0 && str[idx+stridx] == pat[stridx] {
			stridx--
		}

		if stridx < 0 {
			return idx
		} else {
			idx += int(math.Max(1, float64(stridx)-float64(bmt[str[stridx]])))
		}
	}

	return -1
}

func badmatchtable(pat string) []int {
	bmt := make([]int, 256)

	for i := 0; i < 256; i++ {
		bmt[i] = -1
	}

	for i := 0; i < len(pat); i++ {
		bmt[pat[i]] = i
	}

	return bmt
}
