package recursion

import (
	"log"
	"math"
)

func TestMaxScore() {
	//n := 4
	words := []string{"dog", "cat", "dad", "good"}
	chars := []rune{'a', 'b', 'c', 'd', 'd', 'd', 'g', 'o', 'o'}
	charscore := []int{1, 0, 9, 5, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	charMap := make([]int, len(charscore))
	for i := 0; i < len(chars); i++ {
		charMap[int(chars[i])-97]++
	}

	res := make([][]string, 0)

	maxscore := maxscore(words, charMap, charscore, make([]string, 0), &res)
	log.Print(maxscore)

	for _, w := range res {
		log.Printf("%+v", w)
	}
}

func maxscore(words []string, charMap []int, charScore []int, currSet []string, res *[][]string) int {
	if len(words) == 0 {
		if len(currSet) > 0 {
			*res = append(*res, currSet)
		}
		return 0
	}

	score0 := maxscore(words[1:], deepCopyInt(charMap), charScore, deepCopyString(currSet), res)

	isvalid := true
	tempCharMap := deepClone(charMap)
	var score1 int
	for _, ch := range words[0] {
		if tempCharMap[int(ch)-97] > 0 {
			tempCharMap[int(ch)-97]--
			score1 += charScore[int(ch)-97]
		} else {
			isvalid = false
		}
	}

	if isvalid {
		tempCurrSet := deepCopyString(currSet)
		tempCurrSet = append(tempCurrSet, words[0])
		score1 = score1 + maxscore(words[1:], tempCharMap, charScore, tempCurrSet, res)
	} else {
		score1 = 0
	}

	return int(math.Max(float64(score0), float64(score1)))
}

func deepCopyInt(src []int) []int {
	dst := make([]int, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = src[i]
	}

	return dst
}

func deepCopyString(src []string) []string {
	dst := make([]string, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = src[i]
	}

	return dst
}
