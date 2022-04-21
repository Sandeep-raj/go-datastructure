package stacksandqueues

import (
	"log"
	"math"
)

/*
Asteroid Collision

1. You are given an array asteroids of integers representing asteroids in a row.
2. For each asteroid, the absolute value represents its size, and the sign represents its direction (positive meaning right, negative meaning left).
3. Each asteroid moves at the same speed.
4. You need to find out the state of the asteroids after all collisions.
5. If two asteroids meet, the smaller one will explode. If both are the same size, both will explode. Two asteroids moving in the same direction will never meet.

Sample Input
7
3
2
-1
3
-3
3
-4

Sample Output
1
-4
*/

func TestAsteroidCollision() {
	//inpArr := []int{3, 2, -1, 3, -3, 3, -4}
	inpArr := []int{10, 2, -5}
	asteroidCollision(inpArr)
}

func asteroidCollision(inpArr []int) {
	stack := make([]int, 0)

	for i := 0; i < len(inpArr); i++ {
		for len(stack) > 0 && stack[len(stack)-1]*inpArr[i] < 0 && math.Abs(float64(stack[len(stack)-1])) < math.Abs(float64(inpArr[i])) {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			stack = append(stack, inpArr[i])
		} else {
			if stack[len(stack)-1]*inpArr[i] > 0 {
				stack = append(stack, inpArr[i])
			} else {
				if math.Abs(float64(stack[len(stack)-1])) == math.Abs(float64(inpArr[i])) {
					stack = stack[:len(stack)-1]
				}
				continue
			}
		}
	}

	log.Printf("%+v", stack)
}
