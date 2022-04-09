package recursion

import (
	"fmt"
	"log"
)

/*
K-partitions

1. You are given two integers n and k, where n represents number of elements and k represents
     number of subsets.
2. You have to partition n elements in k subsets and print all such configurations.

Sample Input
3
2

Sample Output
1. [1, 2] [3]
2. [1, 3] [2]
3. [1] [2, 3]
*/

func TestKPartition() {
	n := 4
	k := 3
	res := make([]string, k)
	kPartition(k, n, 1, res, 0)
}

func kPartition(k int, n int, idx int, res []string, fillCount int) {
	if idx > n {
		if fillCount == k {
			// print something
			log.Printf("%+v", res)
			return
		}
		return
	}

	for i := 0; i < len(res); i++ {
		if res[i] != "" {
			tempRes := res[i]
			res[i] = fmt.Sprintf("%s%d", tempRes, idx)
			kPartition(k, n, idx+1, res, fillCount)
			res[i] = tempRes
		} else {
			tempRes := res[i]
			res[i] = fmt.Sprintf("%s%d", tempRes, idx)
			kPartition(k, n, idx+1, res, fillCount+1)
			res[i] = tempRes
			break
		}
	}
}
