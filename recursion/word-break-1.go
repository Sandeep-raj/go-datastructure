package recursion

import (
	"fmt"
	"log"
)

/*
Word Break - I

1. You are given n space separated strings, which represents a dictionary of words.
2. You are given another string which represents a sentence.
3. You have to print all possible sentences from the string, such that words of the sentence are
     present in dictionary.

Sample Input
11
i like pep coding pepper eating mango man go in pepcoding
ilikepeppereatingmangoinpepcoding

Sample Output
i like pepper eating man go in pep coding
i like pepper eating man go in pepcoding
i like pepper eating mango in pep coding
i like pepper eating mango in pepcoding
*/

func TestWordBreak() {
	dictStr := []string{"i", "like", "pep", "coding", "pepper", "eating", "mango", "man", "go", "in", "pepcoding"}
	inpStr := "ilikepeppereatingmangoinpepcoding"

	dict := make(map[string]string)

	for i := 0; i < len(dictStr); i++ {
		dict[dictStr[i]] = dictStr[i]
	}

	wordBreak(0, inpStr, dict, "")
}

func wordBreak(idx int, inpStr string, dict map[string]string, resStr string) {
	if idx >= len(inpStr) {
		log.Print(resStr)
		return
	}

	for i := idx; i < len(inpStr); i++ {
		if dict[string(inpStr[idx:i+1])] != "" {
			wordBreak(i+1, inpStr, dict, fmt.Sprintf("%s %s", resStr, inpStr[idx:i+1]))
		}
	}
}
