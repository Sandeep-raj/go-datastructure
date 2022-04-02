package recursion

import "log"

/*
Target Sum Subsets
1. You are given a number n, representing the count of elements.
2. You are given n numbers.
3. You are given a number "tar".
4. Complete the body of printTargetSumSubsets function - without changing signature - to calculate and print all subsets of given elements, the contents of which sum to "tar". Use sample input and output to get more idea.

Sample Input

5
10
20
30
40
50
60

Sample Output
10, 20, 30, .
10, 50, .
20, 40, .
*/

func TestTargetSum() {
	inpArr := []int{10, 20, 30, 40, 50}
	sum := 60
	res := make([][]int, 0)

	targetSum(inpArr, sum, make([]int, 0), &res)

	for _, r := range res {
		log.Print(r)
	}
}

func targetSum(inpArr []int, sum int, currArr []int, res *[][]int) {
	if sum == 0 {
		*res = append(*res, currArr)
		return
	}

	for i := 0; i < len(inpArr); i++ {
		if sum-inpArr[i] >= 0 {
			tempArr := deepClone(inpArr)
			tempInpArr := tempArr[i+1:]
			tempCurrArr := append(deepClone(currArr), inpArr[i])
			targetSum(tempInpArr, sum-inpArr[i], tempCurrArr, res)
		}
	}
}

func deepClone(src []int) []int {
	res := make([]int, len(src))
	for i := 0; i < len(src); i++ {
		res[i] = src[i]
	}
	return res
}
