package stacksandqueues

import "log"

/*
Battle tanks

In a battle field there are N number of tanks numbered from 0 to n-1 in a
left to right order.

you are given an array of sizes[i] representing the size of the ith tank.

A tank can only hit another tank to it's right, only if every tank
in between is shorter than both of them.
A tank can hit another jth tank if i < j and
i.e. the ith tank can see the jth tank if i < j and min(sizes[i], sizes[j]) > max(sizes[i+1], sizes[i+2], ..., sizes[j-1]).

You are required to return an integer denoting the maximum tanks a single tank can hit.

Sample Input
6
10
6
8
5
11
9

Sample Output
3

Input
sizes = [6,1,4,2,7,3]

Output
3
*/

func TestBattleTanks() {
	//inpArr := []int{6, 1, 4, 2, 7, 3}
	inpArr := []int{10, 6, 8, 5, 11, 9}
	battleTanks(inpArr)
}

func battleTanks(inpArr []int) {
	stack := make([]int, 0)
	res := make([]int, len(inpArr))

	for i := len(inpArr) - 1; i >= 0; i-- {
		attcakedTanks := 0

		if len(stack) > 0 {
			attcakedTanks = 1
		}

		for len(stack) > 0 && stack[len(stack)-1] < inpArr[i] {
			stack = stack[:len(stack)-1]
			attcakedTanks++
		}

		res[i] = attcakedTanks
		stack = append(stack, inpArr[i])
	}

	log.Printf("%+v", res)

}
