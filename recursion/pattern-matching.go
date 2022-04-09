package recursion

import "log"

/*
Pattern Matching

1. You are given a string and a pattern.
2. You've to check if the string is of the same structure as pattern without using any regular
     expressions.

Sample Input
graphtreesgraph
pep

Sample Output
p -> graph, e -> trees, .
*/

func TestPatternMatching() {
	// str1 := "graphtreesgraph"
	// str2 := "pep"
	str1 := "mzaddyfzaddy"
	str2 := "abcb"

	umap := make(map[string]string)
	uniq_arr := make([]string, 0)

	for i := 0; i < len(str2); i++ {
		if umap[string(str2[i])] == "" {
			umap[string(str2[i])] = "-"
			uniq_arr = append(uniq_arr, string(str2[i]))
		}
	}

	patternMatching(0, str1, str2, uniq_arr, 0, umap)

}

func patternMatching(idx int, str1 string, str2 string, uniq_arr []string, uniq_count int, resmap map[string]string) {

	if uniq_count == len(uniq_arr) {
		resStr := ""
		for i := 0; i < len(str2); i++ {
			resStr = resStr + resmap[string(str2[i])]
		}

		if resStr == str1 {
			log.Printf("%+v", resmap)
		}
		return
	}

	for i := idx; i < len(str1); i++ {
		resmap[uniq_arr[uniq_count]] = str1[idx : i+1]
		patternMatching(i+1, str1, str2, uniq_arr, uniq_count+1, resmap)
		resmap[uniq_arr[uniq_count]] = "-"
	}
}
