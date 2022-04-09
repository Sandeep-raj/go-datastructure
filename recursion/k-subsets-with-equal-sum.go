package recursion

import "log"

/*
K Subsets With Equal Sum

1. You are given an array of n distinct integers.
2. You have to divide these n integers into k non-empty subsets such that sum of integers of every
     subset is same.
3. If it is not possible to divide, then print "-1".

Sample Input
6
1
2
3
4
5
6
3

Sample Output
[1, 6] [2, 5] [3, 4]
*/

func TestKSubsets() {
	// n := 6
	k := 3
	inpArr := []int{1, 2, 3, 4, 5, 6}
	totSum := 0
	sumVal := 0

	for i := 0; i < len(inpArr); i++ {
		totSum += inpArr[i]
	}

	if totSum%k != 0 {
		log.Print(-1)
		return
	} else {
		sumVal = totSum / k
	}

	totArr := make([]int, k)

	for i := 0; i < k; i++ {
		totArr[i] = sumVal
	}

	resArr := make([][]int, k)

	subsetsPartition(0, inpArr, totArr, resArr)

}

func subsetsPartition(idx int, inpArr []int, totArr []int, resArr [][]int) {

	if idx == len(inpArr) {
		log.Printf("%+v", resArr)
	}

	for i := idx; i < len(inpArr); i++ {
		for j := 0; j < len(totArr); j++ {
			if totArr[j]-inpArr[i] >= 0 {
				totArr[j] = totArr[j] - inpArr[i]
				resArr[j] = append(resArr[j], inpArr[i])
				subsetsPartition(idx+1, inpArr, totArr, resArr)
				resArr[j] = resArr[j][:len(resArr[j])-1]
				totArr[j] = totArr[j] + inpArr[i]
			}
		}
	}
}
