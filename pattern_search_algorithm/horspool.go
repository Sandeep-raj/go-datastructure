package patternsearchalgorithm

import (
	"log"
	"math"
)

func TestHorsPool() {
	str := "THIS IS A TEST TEXT"
	pat := "TEXT"

	idx := horspool(str, pat)
	log.Print(idx)
}

func horspool(str string, pat string) int {
	sTable := shiftTable(pat)

	n := len(str)
	m := len(pat)

	idx := 0
	for idx <= n-m {
		stridx := m - 1
		for stridx >= 0 && str[idx+stridx] == pat[stridx] {
			stridx--
		}

		if stridx < 0 {
			return idx
		} else {
			currShift := (*sTable)[str[idx+m-1]]
			if currShift == 0 && str[idx+m-1] != pat[m-1] {
				currShift = m
			}

			idx += int(math.Max(1, float64(currShift)))
		}
	}
	return -1
}

func shiftTable(pat string) *map[byte]int {
	m := len(pat)

	sTable := make(map[byte]int)

	for i := 0; i < m; i++ {
		sTable[pat[i]] = m - i - 1
	}

	return &sTable
}
