package recursion

import (
	"fmt"
	"log"
)

/*
Get Maze Paths
1. You are given a number n and a number m representing number of rows and columns in a maze.
2. You are standing in the top-left corner and have to reach the bottom-right corner. Only two moves are allowed 'h' (1-step horizontal) and 'v' (1-step vertical).
3. Complete the body of getMazePath function - without changing signature - to get the list of all paths that can be used to move from top-left to bottom-right.
Use sample input and output to take idea about output.

Input
3
3

Output
[hhvv, hvhv, hvvh, vhhv, vhvh, vvhh]
*/

func getMazePath(r int, c int, curPath string, res *[]string) {
	if r == 1 && c == 1 {
		*res = append(*res, curPath)
	}

	if r > 1 {
		getMazePath(r-1, c, curPath+"v", res)
	}

	if c > 1 {
		getMazePath(r, c-1, curPath+"h", res)
	}
}

func TestMazePath() {
	res := make([]string, 0)
	//getMazePath(3, 3, "", &res)
	getMazePathWithJumps(3, 5, "", &res)

	for _, str := range res {
		log.Print(str)
	}
}

/*
Get Maze Path With Jumps

1. You are given a number n and a number m representing number of rows and columns in a maze.
2. You are standing in the top-left corner and have to reach the bottom-right corner.
3. In a single move you are allowed to jump 1 or more steps horizontally (as h1, h2, .. ), or 1 or more steps vertically (as v1, v2, ..) or 1 or more steps diagonally (as d1, d2, ..).
4. Complete the body of getMazePath function - without changing signature - to get the list of all paths that can be used to move from top-left to bottom-right.
Use sample input and output to take idea about output.

Sample Input
2
2

Sample Output
[h1v1, v1h1, d1]
*/

func getMazePathWithJumps(r int, c int, currPath string, res *[]string) {
	if r == 1 && c == 1 {
		*res = append(*res, currPath)
	}

	if r > 1 {
		for i := 1; i <= r-1; i++ {
			getMazePathWithJumps(r-i, c, fmt.Sprintf("%s%s%d", currPath, "v", i), res)
		}
	}

	if c > 1 {
		for i := 1; i <= c-1; i++ {
			getMazePathWithJumps(r, c-i, fmt.Sprintf("%s%s%d", currPath, "h", i), res)
		}
	}

	if r > 1 && c > 1 {
		index := 0
		if r < c {
			index = r
		} else {
			index = c
		}

		for i := 1; i <= index; i++ {
			getMazePathWithJumps(r-i, c-i, fmt.Sprintf("%s%s%d", currPath, "d", i), res)
		}
	}
}
