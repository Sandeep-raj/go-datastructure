package recursion

import (
	"log"
	"math"
)

/*
Tug Of War

1. You are given an array of n integers.
2. You have to divide these n integers into 2 subsets such that the difference of sum of two subsets
     is as minimum as possible.
3. If n is even, both set will contain exactly n/2 elements. If  is odd, one set will contain (n-1)/2 and
    other set will contain (n+1)/2 elements.
3. If it is not possible to divide, then print "-1".


Sample Input
6
1
2
3
4
5
6

Sample Output
[1, 3, 6] [2, 4, 5]
*/

var globalInt int

func TestTogOfWar() {
	inpArr := []int{1, 2, 3, 4, 5, 6}
	globalInt = math.MaxInt
	res := make([][]int, 2)
	finalRes := make([][]int, 0)

	tugOfWar(0, inpArr, res, &finalRes)
	log.Printf("%+v", finalRes)
}

func tugOfWar(idx int, inpArr []int, res [][]int, finalRes *[][]int) {
	if idx == len(inpArr) {
		if len(res[0]) == len(inpArr)/2 || len(res[1]) == len(inpArr)/2 {
			if math.Abs(float64(sumIntArr(res[0])-sumIntArr(res[1]))) < float64(globalInt) {
				globalInt = int(math.Abs(float64(sumIntArr(res[0]) - sumIntArr(res[1]))))
				*finalRes = append(make([][]int, 0), DeepCopy(res)...)
			}
		}
		return
	}

	currVal := inpArr[idx]
	res[0] = append(res[0], currVal)
	tugOfWar(idx+1, inpArr, res, finalRes)
	res[0] = res[0][:len(res[0])-1]

	res[1] = append(res[1], currVal)
	tugOfWar(idx+1, inpArr, res, finalRes)
	res[1] = res[1][:len(res[1])-1]
}

func sumIntArr(inpArr []int) int {
	res := 0
	for i := 0; i < len(inpArr); i++ {
		res += inpArr[i]
	}
	return res
}
