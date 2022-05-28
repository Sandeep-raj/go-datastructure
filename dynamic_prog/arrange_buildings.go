package dynamicprog

import "log"

/*
Arrange Buildings

1. You are given a number n, which represents the length of a road. The road has n plots on it's each side.
2. The road is to be so planned that there should not be consecutive buildings on either side of the road.
3. You are required to find and print the number of ways in which the buildings can be built on both side of roads.

Sample Input
6

Sample Output
441
*/

func TestArrangeBuildings() {
	log.Print(arrangeBuildings(6))
}

func arrangeBuildings(n int) int {
	combinationBuildingSide1 := countBinaryStringsTopDown(n)
	return combinationBuildingSide1 * combinationBuildingSide1
}
